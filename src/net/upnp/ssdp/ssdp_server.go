// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

// A SSDPServer represents a packet of SSDP.
type SSDPServer struct {
	IsRunnable chan bool
}

// NewSSDPServer returns a new SSDPServer.
func NewSSDPServer() *SSDPServer {
	ssdpPkt := &SSDPServer{}
	ssdpPkt.IsRunnable = make(chan bool)
	return ssdpPkt
}

// Start starts this server.
func (self *SSDPServer) Start() (err error) {
	go func() {
		for {
			select {
			case <-self.IsRunnable:
				return
			default:
				// Do other stuff
			}
		}
	}()
	return nil
}

// Stop stops this server.
func (self *SSDPServer) Stop() (err error) {
	self.IsRunnable <- false
	return nil
}
