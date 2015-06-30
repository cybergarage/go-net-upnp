// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"bytes"
	"net"
	"strings"
)

const (
	SSDP_MAX_PACKET_SIZE = 8192
)

// A SSDPPacket represents a ssdpPkt of SSDP.
type SSDPPacket struct {
	FirstLines    []string
	Headers       map[string]string
	Bytes         []byte
	From          *net.UDPAddr
	Content       []byte
	ContentLength int
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
func NewSSDPPacketFromBytes(bytes []byte) *SSDPPacket {
	ssdpPkt := NewSSDPPacket()
	ssdpPkt.Bytes = bytes
	return ssdpPkt
}

func (ssdpPkt *SSDPPacket) parse(inBytes []byte) error {
	// First Line

	pktFirstLineSep := []byte(CRLF)
	pktFirstLineIdx := bytes.Index(inBytes, pktFirstLineSep)
	pktFirstLine := string(inBytes[0:pktFirstLineIdx])
	//log.Trace(fmt.Sprintf("First Line: %s", pktFirstLine))
	ssdpPkt.FirstLines = strings.Split(pktFirstLine, " ")

	// Read Response Header

	pktBodySep := []byte(CRLF + CRLF)
	pktBodyIdx := bytes.Index(inBytes, pktBodySep)
	pktBodyIdx += len(pktBodySep)

	pktHeaderStrings := string(inBytes[(pktFirstLineIdx + len(CRLF)):(pktBodyIdx - 1)])
	for _, headerLine := range strings.Split(pktHeaderStrings, CRLF) {
		headerStrings := strings.Split(headerLine, ": ")
		if len(headerStrings) < 2 {
			continue
		}
		key := headerStrings[0]
		value := headerStrings[1]
		//log.Trace(fmt.Sprintf("[%d] %s : %s", n, key, value))
		ssdpPkt.Headers[key] = value
	}

	// Decode Response Body

	ssdpPkt.Content = inBytes[pktBodyIdx:]
	ssdpPkt.ContentLength = len(ssdpPkt.Content)

	return nil
}
