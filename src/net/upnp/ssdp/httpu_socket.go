// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"fmt"
	"net"
)

// A HTTPUSocket represents a socket of HTTPU.
type HTTPUSocket struct {
	*HTTPMUSocket
}

// NewHTTPUSocket returns a new HTTPUSocket.
func NewHTTPUSocket() *HTTPUSocket {
	ssdpSock := &HTTPUSocket{}
	ssdpSock.HTTPMUSocket = NewHTTPMUSocket()
	return ssdpSock
}

// Bind binds to SSDP multicast address.
func (self *HTTPUSocket) Bind(port int) error {
	err := self.Close()
	if err != nil {
		return err
	}

	localAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	self.Conn, err = net.ListenUDP("udp", localAddr)
	if err != nil {
		return err
	}

	return nil
}

// Write sends the specified bytes.
func (self *HTTPUSocket) Write(b []byte) (int, error) {
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
