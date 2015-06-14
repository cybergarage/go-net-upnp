// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"net"
)

const (
	SSDP_MAX_PACKET_SIZE = 8192
)

// A SSDPPacket represents a packet of SSDP.
type SSDPPacket struct {
	Bytes []byte
	From  *net.UDPAddr
}

// NewSSDPPacket returns a new SSDPPacket.
func NewSSDPPacket() *SSDPPacket {
	ssdpPkt := &SSDPPacket{}
	ssdpPkt.Bytes = make([]byte, 0)
	return ssdpPkt
}
