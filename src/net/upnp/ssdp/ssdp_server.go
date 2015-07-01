// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"net/upnp/log"
)

// A SSDPListener represents a listener for SSDPServer.
type SSDPListener interface {
	DeviceNotifyReceived(ssdpPkt *SSDPPacket)
}

// A SSDPServer represents a packet of SSDP.
type SSDPServer struct {
	Socket   *SSDPSocket
	Listener SSDPListener
}

// NewSSDPServer returns a new SSDPServer.
func NewSSDPServer() *SSDPServer {
	ssdpPkt := &SSDPServer{}
	ssdpPkt.Socket = NewSSDPSocket()
	ssdpPkt.Listener = nil
	return ssdpPkt
}

// Start starts this server.
func (self *SSDPServer) Start() error {
	err := self.Socket.Bind()
	if err != nil {
		return err
	}
	go handleSSDPConnection(self)
	return nil
}

// Stop stops this server.
func (self *SSDPServer) Stop() error {
	err := self.Socket.Close()
	if err != nil {
		return err
	}
	return nil
}

func handleSSDPConnection(self *SSDPServer) {
	for {
		ssdpPkt, err := self.Socket.Read()
		if err != nil {
			log.Error(err)
			break
		}

		if len(ssdpPkt.Bytes) <= 0 {
			continue
		}

		if self.Listener != nil {
			self.Listener.DeviceNotifyReceived(ssdpPkt)
		}
	}
}
