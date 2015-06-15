// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

type SSDPListener interface {
    DeviceNotifyReceived(ssdpPkt *SSDPPacket)
}

// A SSDPServer represents a packet of SSDP.
type SSDPServer struct {
	Socket *SSDPSocket
	Listeners []SSDPListener
}

// NewSSDPServer returns a new SSDPServer.
func NewSSDPServer() *SSDPServer {
	ssdpPkt := &SSDPServer{}
	ssdpPkt.Socket = NewSSDPSocket()
	ssdpPkt.Listeners = make([]SSDPListener, 0)
	return ssdpPkt
}

// Start starts this server.
func (self *SSDPServer) Start() (error) {
	err := self.Socket.Bind()
	if err != nil {
		return err
	}
	go handleSSDPConnection(self)
	return nil
}

// Stop stops this server.
func (self *SSDPServer) Stop() (error) {
	err := self.Socket.Close()
	if err != nil {
		return err
	}
	return nil
}

func handleSSDPConnection(self *SSDPServer) {
	for ;; {
		ssdpPkt, err := self.Socket.Read()
		if err != nil {
			break
		}
		
		for _, listener := range self.Listeners {
        	listener.DeviceNotifyReceived(ssdpPkt)
        }
	}
}
