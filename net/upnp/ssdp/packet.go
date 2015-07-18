// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/cybergarage/go-net-upnp/net/upnp/http"
)

const (
	errorPacketHeaderNotFound = "header (%s) is not found"
)

// A Packet represents a ssdpPkt of SSDP.
type Packet struct {
	FirstLines []string
	Headers    map[string]string
	From       net.UDPAddr
	Interface  net.Interface
}

// NewPacket returns a new Packet.
func NewPacket() *Packet {
	ssdpPkt := &Packet{}
	ssdpPkt.FirstLines = make([]string, 0)
	ssdpPkt.Headers = make(map[string]string)
	return ssdpPkt
}

// NewPacket returns a new Packet.
func NewPacketFromBytes(bytes []byte) (*Packet, error) {
	ssdpPkt := NewPacket()
	err := ssdpPkt.parse(bytes)
	if err != nil {
		return nil, err
	}
	return ssdpPkt, nil
}

func (self *Packet) parse(pktBytes []byte) error {
	if len(pktBytes) <= 0 {
		return errors.New(errorZeroPacket)
	}

	// Read first line

	pktFirstLineSep := []byte(CRLF)
	pktFirstLineIdx := bytes.Index(pktBytes, pktFirstLineSep)
	if pktFirstLineIdx == -1 {
		return errors.New(fmt.Sprintf(errorPacketFirstLineNotFound, string(pktBytes)))
	}
	pktFirstLine := string(pktBytes[0:pktFirstLineIdx])
	self.FirstLines = strings.Split(pktFirstLine, SP)

	// Read Headers

	pktBodySep := []byte(CRLF + CRLF)
	pktBodyIdx := bytes.Index(pktBytes, pktBodySep)
	if pktBodyIdx == -1 {
		pktBodyIdx = len(pktBytes) - 1
	} else {
		pktBodyIdx += len(pktBodySep)
	}

	pktBeginIdx := pktFirstLineIdx + len(CRLF)
	pktEndIdx := pktBodyIdx - 1

	if (pktBeginIdx < 0) || (pktEndIdx < 0) || (pktEndIdx < pktBeginIdx) || (len(pktBytes)-1) < pktBeginIdx || (len(pktBytes)-1) < pktEndIdx {
		return errors.New(fmt.Sprintf(errorPacketHeadersNotFound, pktBeginIdx, pktEndIdx, string(pktBytes)))
	}

	pktHeaderStrings := string(pktBytes[pktBeginIdx:pktEndIdx])
	for _, headerLine := range strings.Split(pktHeaderStrings, CRLF) {
		headerStrings := strings.Split(headerLine, ": ")
		if len(headerStrings) < 2 {
			continue
		}
		key := strings.ToUpper(headerStrings[0])
		value := headerStrings[1]
		self.Headers[key] = value
	}

	return nil
}

func (self *Packet) isMethod(name string) bool {
	if len(self.FirstLines) < 1 {
		return false
	}
	return (self.FirstLines[0] == name)
}

func (self *Packet) IsNotifyRequest() bool {
	return self.isMethod(NOTIFY)
}

func (self *Packet) IsSearchRequest() bool {
	return self.isMethod(M_SEARCH)
}

func (self *Packet) SetMethod(method string) error {
	self.FirstLines = make([]string, 3)
	self.FirstLines[0] = method
	self.FirstLines[1] = HTTP_PATH
	self.FirstLines[2] = fmt.Sprintf("HTTP/%s", HTTP_VERSION)
	return nil
}

func (self *Packet) SetStatusCode(code int) error {
	self.FirstLines = make([]string, 3)
	self.FirstLines[0] = fmt.Sprintf("HTTP/%s", HTTP_VERSION)
	self.FirstLines[1] = fmt.Sprintf("%d", code)
	self.FirstLines[2] = http.StatusCodeToString(code)
	return nil
}

func (self *Packet) GetStatusCode() int {
	if len(self.FirstLines) < 2 {
		return 0
	}
	code, err := strconv.Atoi(self.FirstLines[1])
	if err != nil {
		return 0
	}
	return code
}

func (self *Packet) SetHeaderString(name string, value string) error {
	self.Headers[name] = value
	return nil
}

func (self *Packet) GetHeaderString(name string) (string, error) {
	value, ok := self.Headers[name]
	if !ok {
		return "", errors.New(fmt.Sprintf(errorPacketHeaderNotFound, name))
	}

	return value, nil
}

func (self *Packet) IsHeaderString(name string, value string) bool {
	headerValue, err := self.GetHeaderString(name)
	if err != nil {
		return false
	}
	return (headerValue == value)
}

func (self *Packet) SetHeaderInt(name string, value int) error {
	return self.SetHeaderString(name, strconv.Itoa(value))
}

func (self *Packet) GetHeaderInt(name string) (int, error) {
	svalue, err := self.GetHeaderString(name)
	if err != nil {
		return 0, err
	}
	ivalue, err := strconv.Atoi(svalue)
	if err != nil {
		return 0, err
	}
	return ivalue, nil
}

func (self *Packet) SetHost(value string) error {
	return self.SetHeaderString(HOST, value)
}

func (self *Packet) GetHost() (string, error) {
	return self.GetHeaderString(HOST)
}

func (self *Packet) SetDate(value string) error {
	return self.SetHeaderString(DATE, value)
}

func (self *Packet) GetDate() (string, error) {
	return self.GetHeaderString(DATE)
}

func (self *Packet) SetLocation(value string) error {
	return self.SetHeaderString(LOCATION, value)
}

func (self *Packet) GetLocation() (string, error) {
	return self.GetHeaderString(LOCATION)
}

func (self *Packet) SetCacheControl(value string) error {
	return self.SetHeaderString(CACHE_CONTROL, value)
}

func (self *Packet) GetCacheControl() (string, error) {
	return self.GetHeaderString(CACHE_CONTROL)
}

func (self *Packet) SetST(value string) error {
	return self.SetHeaderString(ST, value)
}

func (self *Packet) GetST() (string, error) {
	return self.GetHeaderString(ST)
}

func (self *Packet) SetMX(value int) error {
	return self.SetHeaderInt(MX, value)
}

func (self *Packet) GetMX() (int, error) {
	return self.GetHeaderInt(MX)
}

func (self *Packet) SetMAN(value string) error {
	return self.SetHeaderString(MAN, value)
}

func (self *Packet) GetMAN() (string, error) {
	return self.GetHeaderString(MAN)
}

func (self *Packet) SetNT(value string) error {
	return self.SetHeaderString(NT, value)
}

func (self *Packet) GetNT() (string, error) {
	return self.GetHeaderString(NT)
}

func (self *Packet) SetNTS(value string) error {
	return self.SetHeaderString(NTS, value)
}

func (self *Packet) GetNTS() (string, error) {
	return self.GetHeaderString(NTS)
}

func (self *Packet) SetUSN(value string) error {
	return self.SetHeaderString(USN, value)
}

func (self *Packet) GetUSN() (string, error) {
	return self.GetHeaderString(USN)
}

func (self *Packet) SetEXT(value string) error {
	return self.SetHeaderString(EXT, value)
}

func (self *Packet) GetEXT() (string, error) {
	return self.GetHeaderString(EXT)
}

func (self *Packet) SetSID(value string) error {
	return self.SetHeaderString(SID, value)
}

func (self *Packet) GetSID() (string, error) {
	return self.GetHeaderString(SID)
}

func (self *Packet) SetSEQ(value string) error {
	return self.SetHeaderString(SEQ, value)
}

func (self *Packet) GetSEQ() (string, error) {
	return self.GetHeaderString(SEQ)
}

func (self *Packet) SetCallback(value string) error {
	return self.SetHeaderString(CALLBACK, value)
}

func (self *Packet) GetCallback() (string, error) {
	return self.GetHeaderString(CALLBACK)
}

func (self *Packet) SetTimeout(value string) error {
	return self.SetHeaderString(TIMEOUT, value)
}

func (self *Packet) GetTimeout() (string, error) {
	return self.GetHeaderString(TIMEOUT)
}

func (self *Packet) SetServer(value string) error {
	return self.SetHeaderString(SERVER, value)
}

func (self *Packet) GetServer() (string, error) {
	return self.GetHeaderString(SERVER)
}

func (self *Packet) SetBOOTID_UPNP_ORG(value string) error {
	return self.SetHeaderString(BOOTID_UPNP_ORG, value)
}

func (self *Packet) GetBOOTID_UPNP_ORG() (string, error) {
	return self.GetHeaderString(BOOTID_UPNP_ORG)
}

func (self *Packet) String() string {
	var pktBuf bytes.Buffer

	// Write First line

	firstLine := strings.Join(self.FirstLines, SP)
	pktBuf.WriteString(firstLine)
	pktBuf.WriteString(CRLF)

	// Write Headers

	for name, value := range self.Headers {
		pktBuf.WriteString(fmt.Sprintf("%s: %s%s", name, value, CRLF))
	}

	pktBuf.WriteString(CRLF)

	return pktBuf.String()
}

func (self *Packet) Bytes() []byte {
	return []byte(self.String())
}
