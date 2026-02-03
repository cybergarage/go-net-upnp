// Copyright 2015 The go-net-upnp Authors. All rights reserved.
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

// NewPacketFromBytes parses a Packet from raw SSDP bytes.
func NewPacketFromBytes(bytes []byte) (*Packet, error) {
	ssdpPkt := NewPacket()
	err := ssdpPkt.parse(bytes)
	if err != nil {
		return nil, err
	}
	return ssdpPkt, nil
}

func (pkt *Packet) parse(pktBytes []byte) error {
	if len(pktBytes) == 0 {
		return errors.New(errorZeroPacket)
	}

	// Read first line

	pktFirstLineSep := []byte(CRLF)
	pktFirstLineIdx := bytes.Index(pktBytes, pktFirstLineSep)
	if pktFirstLineIdx == -1 {
		return fmt.Errorf(errorPacketFirstLineNotFound, string(pktBytes))
	}
	pktFirstLine := string(pktBytes[0:pktFirstLineIdx])
	pkt.FirstLines = strings.Split(pktFirstLine, SP)

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
		return fmt.Errorf(errorPacketHeadersNotFound, pktBeginIdx, pktEndIdx, string(pktBytes))
	}

	pktHeaderStrings := string(pktBytes[pktBeginIdx:pktEndIdx])
	for headerLine := range strings.SplitSeq(pktHeaderStrings, CRLF) {
		headerStrings := strings.Split(headerLine, ": ")
		if len(headerStrings) < 2 {
			continue
		}
		key := strings.ToUpper(headerStrings[0])
		value := headerStrings[1]
		pkt.Headers[key] = value
	}

	return nil
}

func (pkt *Packet) isMethod(name string) bool {
	if len(pkt.FirstLines) < 1 {
		return false
	}
	return (pkt.FirstLines[0] == name)
}

func (pkt *Packet) IsNotifyRequest() bool {
	return pkt.isMethod(Notify)
}

func (pkt *Packet) IsSearchRequest() bool {
	return pkt.isMethod(MSearch)
}

func (pkt *Packet) SetMethod(method string) error {
	pkt.FirstLines = make([]string, 3)
	pkt.FirstLines[0] = method
	pkt.FirstLines[1] = HTTPPath
	pkt.FirstLines[2] = fmt.Sprintf("HTTP/%s", HTTPVersion)
	return nil
}

func (pkt *Packet) SetStatusCode(code int) error {
	pkt.FirstLines = make([]string, 3)
	pkt.FirstLines[0] = fmt.Sprintf("HTTP/%s", HTTPVersion)
	pkt.FirstLines[1] = fmt.Sprintf("%d", code)
	pkt.FirstLines[2] = http.StatusCodeToString(code)
	return nil
}

func (pkt *Packet) GetStatusCode() int {
	if len(pkt.FirstLines) < 2 {
		return 0
	}
	code, err := strconv.Atoi(pkt.FirstLines[1])
	if err != nil {
		return 0
	}
	return code
}

func (pkt *Packet) SetHeaderString(name string, value string) error {
	pkt.Headers[name] = value
	return nil
}

func (pkt *Packet) GetHeaderString(name string) (string, error) {
	value, ok := pkt.Headers[name]
	if !ok {
		return "", fmt.Errorf(errorPacketHeaderNotFound, name)
	}

	return value, nil
}

func (pkt *Packet) IsHeaderString(name string, value string) bool {
	headerValue, err := pkt.GetHeaderString(name)
	if err != nil {
		return false
	}
	return (headerValue == value)
}

func (pkt *Packet) SetHeaderInt(name string, value int) error {
	return pkt.SetHeaderString(name, strconv.Itoa(value))
}

func (pkt *Packet) GetHeaderInt(name string) (int, error) {
	svalue, err := pkt.GetHeaderString(name)
	if err != nil {
		return 0, err
	}
	ivalue, err := strconv.Atoi(svalue)
	if err != nil {
		return 0, err
	}
	return ivalue, nil
}

func (pkt *Packet) SetHost(value string) error {
	return pkt.SetHeaderString(Host, value)
}

func (pkt *Packet) GetHost() (string, error) {
	return pkt.GetHeaderString(Host)
}

func (pkt *Packet) SetDate(value string) error {
	return pkt.SetHeaderString(Date, value)
}

func (pkt *Packet) GetDate() (string, error) {
	return pkt.GetHeaderString(Date)
}

func (pkt *Packet) SetLocation(value string) error {
	return pkt.SetHeaderString(Location, value)
}

func (pkt *Packet) GetLocation() (string, error) {
	return pkt.GetHeaderString(Location)
}

func (pkt *Packet) SetCacheControl(value string) error {
	return pkt.SetHeaderString(CacheControl, value)
}

func (pkt *Packet) GetCacheControl() (string, error) {
	return pkt.GetHeaderString(CacheControl)
}

func (pkt *Packet) SetST(value string) error {
	return pkt.SetHeaderString(ST, value)
}

func (pkt *Packet) GetST() (string, error) {
	return pkt.GetHeaderString(ST)
}

func (pkt *Packet) SetMX(value int) error {
	return pkt.SetHeaderInt(MX, value)
}

func (pkt *Packet) GetMX() (int, error) {
	return pkt.GetHeaderInt(MX)
}

func (pkt *Packet) SetMAN(value string) error {
	return pkt.SetHeaderString(MAN, value)
}

func (pkt *Packet) GetMAN() (string, error) {
	return pkt.GetHeaderString(MAN)
}

func (pkt *Packet) SetNT(value string) error {
	return pkt.SetHeaderString(NT, value)
}

func (pkt *Packet) GetNT() (string, error) {
	return pkt.GetHeaderString(NT)
}

func (pkt *Packet) SetNTS(value string) error {
	return pkt.SetHeaderString(NTS, value)
}

func (pkt *Packet) GetNTS() (string, error) {
	return pkt.GetHeaderString(NTS)
}

func (pkt *Packet) SetUSN(value string) error {
	return pkt.SetHeaderString(USN, value)
}

func (pkt *Packet) GetUSN() (string, error) {
	return pkt.GetHeaderString(USN)
}

func (pkt *Packet) SetEXT(value string) error {
	return pkt.SetHeaderString(EXT, value)
}

func (pkt *Packet) GetEXT() (string, error) {
	return pkt.GetHeaderString(EXT)
}

func (pkt *Packet) SetSID(value string) error {
	return pkt.SetHeaderString(SID, value)
}

func (pkt *Packet) GetSID() (string, error) {
	return pkt.GetHeaderString(SID)
}

func (pkt *Packet) SetSEQ(value string) error {
	return pkt.SetHeaderString(SEQ, value)
}

func (pkt *Packet) GetSEQ() (string, error) {
	return pkt.GetHeaderString(SEQ)
}

func (pkt *Packet) SetCallback(value string) error {
	return pkt.SetHeaderString(Callback, value)
}

func (pkt *Packet) GetCallback() (string, error) {
	return pkt.GetHeaderString(Callback)
}

func (pkt *Packet) SetTimeout(value string) error {
	return pkt.SetHeaderString(Timeout, value)
}

func (pkt *Packet) GetTimeout() (string, error) {
	return pkt.GetHeaderString(Timeout)
}

func (pkt *Packet) SetServer(value string) error {
	return pkt.SetHeaderString(Server, value)
}

func (pkt *Packet) GetServer() (string, error) {
	return pkt.GetHeaderString(Server)
}

func (pkt *Packet) SetBootIDUPnPOrg(value string) error {
	return pkt.SetHeaderString(BootIDUPnPOrg, value)
}

func (pkt *Packet) GetBootIDUPnPOrg() (string, error) {
	return pkt.GetHeaderString(BootIDUPnPOrg)
}

func (pkt *Packet) String() string {
	var pktBuf bytes.Buffer

	// Write First line

	firstLine := strings.Join(pkt.FirstLines, SP)
	pktBuf.WriteString(firstLine)
	pktBuf.WriteString(CRLF)

	// Write Headers

	for name, value := range pkt.Headers {
		pktBuf.WriteString(fmt.Sprintf("%s: %s%s", name, value, CRLF))
	}

	pktBuf.WriteString(CRLF)

	return pktBuf.String()
}

func (pkt *Packet) Bytes() []byte {
	return []byte(pkt.String())
}
