// Copyright 2015 Satoshi Konno. All rights reserved.
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
func (self *HTTPMUSocket) Bind(ifi net.Interface) error {
	err := self.Close()
	if err != nil {
		return err
	}

	mcastAddr, err := net.ResolveUDPAddr("udp", MULTICAST_ADDRESS)
	if err != nil {
		return err
	}

	self.Conn, err = net.ListenMulticastUDP("udp", &ifi, mcastAddr)
	if err != nil {
		return errors.New(fmt.Sprintf("%s (%s)", err.Error(), ifi.Name))
	}

	self.Interface = ifi

	return nil
}

// Write sends the specified bytes.
func (self *HTTPMUSocket) Write(b []byte) (int, error) {
	if self.Conn == nil {
		return 0, errors.New(errorSocketIsClosed)
	}

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
