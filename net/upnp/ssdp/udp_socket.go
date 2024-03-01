// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"errors"
	"net"
)

const (
	errorSocketIsClosed = "socket is closed"
)

// A UDPSocket represents a socket for UDP.
type UDPSocket struct {
	Conn      *net.UDPConn
	readBuf   []byte
	Interface net.Interface
}

// NewUDPSocket returns a new UDPSocket.
func NewUDPSocket() *UDPSocket {
	uppSock := &UDPSocket{}
	uppSock.readBuf = make([]byte, MaxPacketSize)
	return uppSock
}

// Close closes the current opened socket.
func (socket *UDPSocket) Close() error {
	if socket.Conn == nil {
		return nil
	}
	err := socket.Conn.Close()
	if err != nil {
		return err
	}

	socket.Conn = nil
	socket.Interface = net.Interface{}

	return nil
}

// Read reads from the current opend socket.
func (socket *UDPSocket) Read() (*Packet, error) {
	if socket.Conn == nil {
		return nil, errors.New(errorSocketIsClosed)
	}

	n, from, err := socket.Conn.ReadFromUDP(socket.readBuf)
	if err != nil {
		return nil, err
	}

	ssdpPkt, err := NewPacketFromBytes(socket.readBuf[:n])
	if err != nil {
		return nil, err
	}

	ssdpPkt.From = *from
	ssdpPkt.Interface = socket.Interface

	return ssdpPkt, nil
}
