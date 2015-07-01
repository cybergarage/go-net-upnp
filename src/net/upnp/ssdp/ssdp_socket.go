// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"fmt"
	"net"
	"net/upnp/log"
	"time"
)

// A SSDPSocket represents a Socket of SSDP.
type SSDPSocket struct {
	Socket []byte
	Conn   *net.UDPConn
}

// NewSSDPSocket returns a new SSDPSocket.
func NewSSDPSocket() *SSDPSocket {
	ssdpSock := &SSDPSocket{}
	return ssdpSock
}

// Bind binds to SSDP multicast address.
func (self *SSDPSocket) Bind() error {
	err := self.Close()
	if err != nil {
		return err
	}

	ssdpAddr, err := net.ResolveUDPAddr("udp", MULTICAST_ADDRESS)
	if err != nil {
		return err
	}

	self.Conn, err = net.ListenMulticastUDP("udp", nil, ssdpAddr)
	if err != nil {
		return err
	}

	self.Conn.SetDeadline(time.Now().Add(1e9))
	self.Conn.SetReadBuffer(MAX_PACKET_SIZE)

	return nil
}

// Bind binds to SSDP multicast address.
func (self *SSDPSocket) Close() error {
	if self.Conn == nil {
		return nil
	}
	err := self.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}

// Write sends the specified bytes.
func (self *SSDPSocket) Write(b []byte) (int, error) {
	ssdpAddr, err := net.ResolveUDPAddr("udp", MULTICAST_ADDRESS)
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
	ssdpPktBytes := make([]byte, MAX_PACKET_SIZE)

	n, from, err := self.Conn.ReadFromUDP(ssdpPktBytes)
	if err != nil {
		return nil, err
	}

	log.Trace(fmt.Sprintf("SSDPSocket::Read() = %d", n))

	ssdpPkt, err := NewSSDPPacketFromBytes(ssdpPktBytes)
	if err != nil {
		return nil, err
	}
	ssdpPkt.From = from

	return ssdpPkt, nil
}
