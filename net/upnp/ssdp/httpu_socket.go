// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"fmt"
	"net"

	"github.com/cybergarage/go-net-upnp/net/upnp/util"
)

// A HTTPUSocket represents a socket of HTTPU.
type HTTPUSocket struct {
	*UDPSocket
}

// NewHTTPUSocket returns a new HTTPUSocket.
func NewHTTPUSocket() *HTTPUSocket {
	ssdpSock := &HTTPUSocket{}
	ssdpSock.UDPSocket = NewUDPSocket()
	return ssdpSock
}

// Bind binds to SSDP multicast address.
func (self *HTTPUSocket) Bind(ifi net.Interface, port int) error {
	err := self.Close()
	if err != nil {
		return err
	}

	addr, err := util.GetInterfaceAddress(ifi)
	if err != nil {
		return err
	}

	bindAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}

	self.Conn, err = net.ListenUDP("udp", bindAddr)
	if err != nil {
		return err
	}

	self.Interface = ifi

	return nil
}

// Write sends the specified bytes.
func (self *HTTPUSocket) Write(addr string, port int, b []byte) (int, error) {
	toAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return 0, err
	}

	if self.Conn != nil {
		return self.Conn.WriteToUDP(b, toAddr)
	}

	conn, err := net.DialUDP("udp", nil, toAddr)
	if err != nil {
		return 0, err
	}

	return conn.Write(b)
}
