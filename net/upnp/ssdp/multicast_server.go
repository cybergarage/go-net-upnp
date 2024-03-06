// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"errors"
	"net"

	"github.com/cybergarage/go-logger/log"
)

// A MulticastListener represents a listener for MulticastServer.
type MulticastListener interface {
	DeviceNotifyReceived(ssdpReq *Request)
	DeviceSearchReceived(ssdpReq *Request)
}

// A MulticastServer represents a packet of SSDP.
type MulticastServer struct {
	Socket    *HTTPMUSocket
	Listener  MulticastListener
	Interface net.Interface
}

// NewMulticastServer returns a new MulticastServer.
func NewMulticastServer() *MulticastServer {
	server := &MulticastServer{}
	server.Socket = NewHTTPMUSocket()
	server.Listener = nil
	return server
}

// Start starts this server.
func (server *MulticastServer) Start(ifi net.Interface) error {
	err := server.Socket.Bind(ifi)
	if err != nil {
		return err
	}
	server.Interface = ifi
	go handleMulticastConnection(server)
	return nil
}

// Stop stops this server.
func (server *MulticastServer) Stop() error {
	err := server.Socket.Close()
	if err != nil {
		return err
	}
	return nil
}

func handleMulticastConnection(server *MulticastServer) {
	for {
		ssdpPkt, err := server.Socket.Read()
		if err != nil {
			if !errors.Is(err, net.ErrClosed) {
				log.Error(err)
			}
			break
		}

		if server.Listener != nil {
			ssdpReq, _ := NewRequestFromPacket(ssdpPkt)
			switch {
			case ssdpReq.IsNotifyRequest():
				server.Listener.DeviceNotifyReceived(ssdpReq)
			case ssdpReq.IsSearchRequest():
				server.Listener.DeviceSearchReceived(ssdpReq)
			}
		}
	}
}
