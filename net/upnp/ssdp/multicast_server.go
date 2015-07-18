// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"net"

	"github.com/cybergarage/go-net-upnp/net/upnp/log"
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
func (self *MulticastServer) Start(ifi net.Interface) error {
	err := self.Socket.Bind(ifi)
	if err != nil {
		return err
	}
	self.Interface = ifi
	go handleMulticastConnection(self)
	return nil
}

// Stop stops this server.
func (self *MulticastServer) Stop() error {
	err := self.Socket.Close()
	if err != nil {
		return err
	}
	return nil
}

func handleMulticastConnection(self *MulticastServer) {
	for {
		ssdpPkt, err := self.Socket.Read()
		if err != nil {
			log.Error(err)
			break
		}

		if self.Listener != nil {
			ssdpReq, _ := NewRequestFromPacket(ssdpPkt)
			switch {
			case ssdpReq.IsNotifyRequest():
				self.Listener.DeviceNotifyReceived(ssdpReq)
			case ssdpReq.IsSearchRequest():
				self.Listener.DeviceSearchReceived(ssdpReq)
			}
		}
	}
}
