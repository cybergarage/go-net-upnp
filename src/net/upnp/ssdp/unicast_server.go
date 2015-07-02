// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"net/upnp/log"
)

// A UnicastListener represents a listener for UnicastServer.
type UnicastListener interface {
	DeviceResponseReceived(ssdpRes *Response)
}

// A UnicastServer represents a packet of SSDP.
type UnicastServer struct {
	Socket   *HTTPUSocket
	Listener UnicastListener
}

// NewUnicastServer returns a new UnicastServer.
func NewUnicastServer() *UnicastServer {
	ssdpPkt := &UnicastServer{}
	ssdpPkt.Socket = NewHTTPUSocket()
	ssdpPkt.Listener = nil
	return ssdpPkt
}

// Start starts this server.
func (self *UnicastServer) Start(port int) error {
	err := self.Socket.Bind(port)
	if err != nil {
		return err
	}
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

func handleSSDPUnicastConnection(self *UnicastServer) {
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
			ssdpRes, _ := NewResponseFromPacket(ssdpPkt)
			self.Listener.DeviceResponseReceived(ssdpRes)
		}
	}
}
