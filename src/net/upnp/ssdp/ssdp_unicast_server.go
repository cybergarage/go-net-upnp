// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"net/upnp/log"
)

// A SSDPListener represents a listener for SSDPUnicastServer.
type SSDPUnicastListener interface {
	DeviceResponseReceived(ssdpRes *SSDPResponse)
}

// A SSDPUnicastServer represents a packet of SSDP.
type SSDPUnicastServer struct {
	Socket   *HTTPUSocket
	Listener SSDPUnicastListener
}

// NewSSDPUnicastServer returns a new SSDPUnicastServer.
func NewSSDPUnicastServer() *SSDPUnicastServer {
	ssdpPkt := &SSDPUnicastServer{}
	ssdpPkt.Socket = NewHTTPUSocket()
	ssdpPkt.Listener = nil
	return ssdpPkt
}

// Start starts this server.
func (self *SSDPUnicastServer) Start(port int) error {
	err := self.Socket.Bind(port)
	if err != nil {
		return err
	}
	go handleSSDPUnicastConnection(self)
	return nil
}

// Stop stops this server.
func (self *SSDPUnicastServer) Stop() error {
	err := self.Socket.Close()
	if err != nil {
		return err
	}
	return nil
}

func handleSSDPUnicastConnection(self *SSDPUnicastServer) {
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
			ssdpRes, _ := NewSSDPResponseFromPacket(ssdpPkt)
			self.Listener.DeviceResponseReceived(ssdpRes)
		}
	}
}
