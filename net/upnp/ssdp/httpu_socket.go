// Copyright 2015 The go-net-upnp Authors. All rights reserved.
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
func (socket *HTTPUSocket) Bind(ifi net.Interface, port int) error {
	err := socket.Close()
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

	socket.Conn, err = net.ListenUDP("udp", bindAddr)
	if err != nil {
		return err
	}

	socket.Interface = ifi

	return nil
}

// Write sends the specified bytes.
func (socket *HTTPUSocket) Write(addr string, port int, b []byte) (int, error) {
	toAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return 0, err
	}

	if socket.Conn != nil {
		return socket.Conn.WriteToUDP(b, toAddr)
	}

	conn, err := net.DialUDP("udp", nil, toAddr)
	if err != nil {
		return 0, err
	}

	return conn.Write(b)
}
