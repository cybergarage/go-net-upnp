// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"net"

	"github.com/cybergarage/go-net-upnp/net/upnp/log"
)

// A UnicastListener represents a listener for UnicastServer.
type UnicastListener interface {
	DeviceResponseReceived(ssdpRes *Response)
}

// A UnicastServer represents a packet of SSDP.
type UnicastServer struct {
	Socket    *UnicastSocket
	Listener  UnicastListener
	Interface net.Interface
}

// NewUnicastServer returns a new UnicastServer.
func NewUnicastServer() *UnicastServer {
	server := &UnicastServer{}
	server.Socket = NewUnicastSocket()
	server.Listener = nil
	return server
}

// Start starts this server.
func (self *UnicastServer) Start(ifi net.Interface, port int) error {
	err := self.Socket.Bind(ifi, port)
	if err != nil {
		return err
	}
	self.Interface = ifi
	go handleSSDPUnicastConnection(self)
	return nil
}

// Stop stops this server.
func (self *UnicastServer) Stop() error {
	err := self.Socket.Close()
	if err != nil {
		return err
	}
	return nil
}

// Search sends a M-SEARCH request of the specified ST.
func (self *UnicastServer) Search(st string, mx int) error {
	ssdpReq, err := NewSearchRequest(st, mx)
	if err != nil {
		return err
	}

	_, err = self.Socket.WriteRequest(ssdpReq)

	return err
}

func handleSSDPUnicastConnection(self *UnicastServer) {
	for {
		ssdpPkt, err := self.Socket.Read()
		if err != nil {
			log.Error(err)
			break
		}

		if self.Listener != nil {
			ssdpRes, _ := NewResponseFromPacket(ssdpPkt)
			self.Listener.DeviceResponseReceived(ssdpRes)
		}
	}
}
