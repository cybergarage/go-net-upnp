// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"fmt"
	"net"
	"net/upnp/log"
)

// A SSDPSocket represents a Socket of SSDP.
type SSDPSocket struct {
	Socket  []byte
	Conn    *net.UDPConn
	readBuf []byte
}

// NewSSDPSocket returns a new SSDPSocket.
func NewSSDPSocket() *SSDPSocket {
	ssdpSock := &SSDPSocket{}
	ssdpSock.readBuf = make([]byte, MAX_PACKET_SIZE)
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
	defer conn.Close()

	return conn.Write(b)
}

// Read reads a SSDP packet.
func (self *SSDPSocket) Read() (*SSDPPacket, error) {
	n, from, err := self.Conn.ReadFrom(self.readBuf)
	if err != nil {
		return nil, err
	}

	log.Trace(fmt.Sprintf("from %v got message %q\n", from, string(self.readBuf[:n])))

	ssdpPkt, err := NewSSDPPacketFromBytes(self.readBuf[:n])
	if err != nil {
		return nil, err
	}
	ssdpPkt.From = from

	return ssdpPkt, nil
}
