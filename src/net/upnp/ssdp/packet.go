// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
)

// A Packet represents a ssdpPkt of SSDP.
type Packet struct {
	FirstLines []string
	Headers    map[string]string
	From       net.Addr
	Bytes      []byte
}

// NewPacket returns a new Packet.
func NewPacket() *Packet {
	ssdpPkt := &Packet{}
	ssdpPkt.FirstLines = make([]string, 0)
	ssdpPkt.Headers = make(map[string]string)
	ssdpPkt.Bytes = make([]byte, 0)
	return ssdpPkt
}

// NewPacket returns a new Packet.
func NewPacketFromBytes(bytes []byte) (*Packet, error) {
	ssdpPkt := NewPacket()
	ssdpPkt.Bytes = bytes
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

	fmt.Sprintf("ssdp pkt = '%s'", string(pktBytes))
	pktFirstLineSep := []byte(CRLF)
	pktFirstLineIdx := bytes.Index(pktBytes, pktFirstLineSep)
	pktFirstLine := string(pktBytes[0:pktFirstLineIdx])
	self.FirstLines = strings.Split(pktFirstLine, SP)

	// Read Headers

	pktBodySep := []byte(CRLF + CRLF)
	pktBodyIdx := bytes.Index(pktBytes, pktBodySep)
	pktBodyIdx += len(pktBodySep)

	pktHeaderStrings := string(pktBytes[(pktFirstLineIdx + len(CRLF)):(pktBodyIdx - 1)])
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

	return pktBuf.String()
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

func (self *Packet) GetHeaderString(name string) (string, bool) {
	value, ok := self.Headers[name]
	return value, ok
}

func (self *Packet) GetHost() (string, bool) {
	return self.GetHeaderString(HOST)
}

func (self *Packet) GetDate() (string, bool) {
	return self.GetHeaderString(DATE)
}

func (self *Packet) GetLocation() (string, bool) {
	return self.GetHeaderString(LOCATION)
}

func (self *Packet) GetCacheControl() (string, bool) {
	return self.GetHeaderString(CACHE_CONTROL)
}

func (self *Packet) GetST() (string, bool) {
	return self.GetHeaderString(ST)
}

func (self *Packet) GetMX() (string, bool) {
	return self.GetHeaderString(MX)
}

func (self *Packet) GetMAN() (string, bool) {
	return self.GetHeaderString(MAN)
}

func (self *Packet) GetNT() (string, bool) {
	return self.GetHeaderString(NT)
}

func (self *Packet) GetNTS() (string, bool) {
	return self.GetHeaderString(NTS)
}

func (self *Packet) GetUSN() (string, bool) {
	return self.GetHeaderString(USN)
}

func (self *Packet) GetEXT() (string, bool) {
	return self.GetHeaderString(EXT)
}

func (self *Packet) GetSID() (string, bool) {
	return self.GetHeaderString(SID)
}

func (self *Packet) GetSEQ() (string, bool) {
	return self.GetHeaderString(SEQ)
}

func (self *Packet) GetCallback() (string, bool) {
	return self.GetHeaderString(CALLBACK)
}

func (self *Packet) GetTimeout() (string, bool) {
	return self.GetHeaderString(TIMEOUT)
}

func (self *Packet) GetServer() (string, bool) {
	return self.GetHeaderString(SERVER)
}

func (self *Packet) GetBOOTID_UPNP_ORG() (string, bool) {
	return self.GetHeaderString(BOOTID_UPNP_ORG)
}
