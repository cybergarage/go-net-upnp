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

// A SSDPPacket represents a ssdpPkt of SSDP.
type SSDPPacket struct {
	FirstLines []string
	Headers    map[string]string
	From       net.Addr
	Bytes      []byte
}

// NewSSDPPacket returns a new SSDPPacket.
func NewSSDPPacket() *SSDPPacket {
	ssdpPkt := &SSDPPacket{}
	ssdpPkt.FirstLines = make([]string, 0)
	ssdpPkt.Headers = make(map[string]string)
	ssdpPkt.Bytes = make([]byte, 0)
	return ssdpPkt
}

// NewSSDPPacket returns a new SSDPPacket.
func NewSSDPPacketFromBytes(bytes []byte) (*SSDPPacket, error) {
	ssdpPkt := NewSSDPPacket()
	ssdpPkt.Bytes = bytes
	err := ssdpPkt.parse(bytes)
	if err != nil {
		return nil, err
	}
	return ssdpPkt, nil
}

func (self *SSDPPacket) parse(pktBytes []byte) error {
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

func (self *SSDPPacket) String() string {
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

func (self *SSDPPacket) GetHeaderString(name string) (string, bool) {
	value, ok := self.Headers[name]
	return value, ok
}

func (self *SSDPPacket) GetHost() (string, bool) {
	return self.GetHeaderString(HOST)
}

func (self *SSDPPacket) GetDate() (string, bool) {
	return self.GetHeaderString(DATE)
}

func (self *SSDPPacket) GetLocation() (string, bool) {
	return self.GetHeaderString(LOCATION)
}

func (self *SSDPPacket) GetCacheControl() (string, bool) {
	return self.GetHeaderString(CACHE_CONTROL)
}

func (self *SSDPPacket) GetST() (string, bool) {
	return self.GetHeaderString(ST)
}

func (self *SSDPPacket) GetMX() (string, bool) {
	return self.GetHeaderString(MX)
}

func (self *SSDPPacket) GetMAN() (string, bool) {
	return self.GetHeaderString(MAN)
}

func (self *SSDPPacket) GetNT() (string, bool) {
	return self.GetHeaderString(NT)
}

func (self *SSDPPacket) GetNTS() (string, bool) {
	return self.GetHeaderString(NTS)
}

func (self *SSDPPacket) GetUSN() (string, bool) {
	return self.GetHeaderString(USN)
}

func (self *SSDPPacket) GetEXT() (string, bool) {
	return self.GetHeaderString(EXT)
}

func (self *SSDPPacket) GetSID() (string, bool) {
	return self.GetHeaderString(SID)
}

func (self *SSDPPacket) GetSEQ() (string, bool) {
	return self.GetHeaderString(SEQ)
}

func (self *SSDPPacket) GetCallback() (string, bool) {
	return self.GetHeaderString(CALLBACK)
}

func (self *SSDPPacket) GetTimeout() (string, bool) {
	return self.GetHeaderString(TIMEOUT)
}

func (self *SSDPPacket) GetServer() (string, bool) {
	return self.GetHeaderString(SERVER)
}

func (self *SSDPPacket) GetBOOTID_UPNP_ORG() (string, bool) {
	return self.GetHeaderString(BOOTID_UPNP_ORG)
}
