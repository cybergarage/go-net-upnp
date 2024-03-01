// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"errors"
	"fmt"
	"net"
)

// A HTTPMUSocket represents a socket for HTTPMU.
type HTTPMUSocket struct {
	*UDPSocket
}

// NewHTTPMUSocket returns a new HTTPMUSocket.
func NewHTTPMUSocket() *HTTPMUSocket {
	ssdpSock := &HTTPMUSocket{}
	ssdpSock.UDPSocket = NewUDPSocket()
	return ssdpSock
}

// Bind binds to SSDP multicast address.
func (socket *HTTPMUSocket) Bind(ifi net.Interface) error {
	err := socket.Close()
	if err != nil {
		return err
	}

	mcastAddr, err := net.ResolveUDPAddr("udp", MulticastAddress)
	if err != nil {
		return err
	}

	socket.Conn, err = net.ListenMulticastUDP("udp", &ifi, mcastAddr)
	if err != nil {
		return fmt.Errorf("%w (%s)", err, ifi.Name)
	}

	socket.Interface = ifi

	return nil
}

// Write sends the specified bytes.
func (socket *HTTPMUSocket) Write(b []byte) (int, error) {
	if socket.Conn == nil {
		return 0, errors.New(errorSocketIsClosed)
	}

	ssdpAddr, err := net.ResolveUDPAddr("udp", MulticastAddress)
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
