// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"net/upnp/log"
)

// A SSDPListener represents a listener for SSDPMulticastServer.
type SSDPMulticastListener interface {
	DeviceNotifyReceived(ssdpReq *SSDPRequest)
}

// A SSDPMulticastServer represents a packet of SSDP.
type SSDPMulticastServer struct {
	Socket   *HTTPMUSocket
	Listener SSDPMulticastListener
}

// NewSSDPMulticastServer returns a new SSDPMulticastServer.
func NewSSDPMulticastServer() *SSDPMulticastServer {
	ssdpPkt := &SSDPMulticastServer{}
	ssdpPkt.Socket = NewHTTPMUSocket()
	ssdpPkt.Listener = nil
	return ssdpPkt
}

// Start starts this server.
func (self *SSDPMulticastServer) Start() error {
	err := self.Socket.Bind()
	if err != nil {
		return err
	}
	go handleSSDPMulticastConnection(self)
	return nil
}

// Stop stops this server.
func (self *SSDPMulticastServer) Stop() error {
	err := self.Socket.Close()
	if err != nil {
		return err
	}
	return nil
}

func handleSSDPMulticastConnection(self *SSDPMulticastServer) {
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
			ssdpReq, _ := NewSSDPRequestFromPacket(ssdpPkt)
			self.Listener.DeviceNotifyReceived(ssdpReq)
		}
	}
}
