// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"net"
)

const (
	SSDP_ADDR = "239.255.255.250:1900"
)

// A SSDPSocket represents a Socket of SSDP.
type SSDPSocket struct {
	Socket []byte
	Conn   *net.UDPConn
}

// NewSSDPSocket returns a new SSDPSocket.
func NewSSDPSocket() *SSDPSocket {
	dev := &SSDPSocket{}
	return dev
}

// Bind binds to SSDP multicast address.
func (self *SSDPSocket) Bind() (err error) {
	ssdpAddr, err := net.ResolveUDPAddr("udp", SSDP_ADDR)
	if err != nil {
		return err
	}

	self.Conn, err = net.ListenMulticastUDP("udp", nil, ssdpAddr)
	if err != nil {
		return err
	}

	self.Conn.SetReadBuffer(SSDP_MAX_PACKET_SIZE)

	return nil
}

// Write sends the specified bytes.
func (self *SSDPSocket) Write(b []byte) (int, error) {
	ssdpAddr, err := net.ResolveUDPAddr("udp", SSDP_ADDR)
	if err != nil {
		return 0, err
	}

	conn, err := net.DialUDP("udp", nil, ssdpAddr)
	if err != nil {
		return 0, err
	}

	return conn.Write(b)
}

// Read reads a SSDP packet.
func (self *SSDPSocket) Read() (*SSDPPacket, error) {
	ssdpPkt := NewSSDPPacket()

	_, from, err := self.Conn.ReadFromUDP(ssdpPkt.Bytes)
	if err != nil {
		return nil, err
	}

	ssdpPkt.From = from

	return ssdpPkt, nil
}
