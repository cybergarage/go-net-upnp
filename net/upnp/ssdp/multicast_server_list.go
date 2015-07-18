// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"../util"
)

// A MulticastServerList represents a packet of SSDP.
type MulticastServerList struct {
	Listener MulticastListener
	Servers  []*MulticastServer
}

// NewMulticastServerList returns a new MulticastServerList.
func NewMulticastServerList() *MulticastServerList {
	server := &MulticastServerList{}
	server.Servers = make([]*MulticastServer, 0)
	server.Listener = nil
	return server
}

// Start starts this server.
func (self *MulticastServerList) Start() error {
	err := self.Stop()
	if err != nil {
		return err
	}

	ifis, err := util.GetAvailableInterfaces()
	if err != nil {
		return err
	}

	var lastErr error = nil

	self.Servers = make([]*MulticastServer, len(ifis))
	for n, ifi := range ifis {
		server := NewMulticastServer()
		server.Listener = self.Listener
		err := server.Start(ifi)
		if err != nil {
			lastErr = err
		}
		self.Servers[n] = server
	}

	return lastErr
}

// Stop stops this server.
func (self *MulticastServerList) Stop() error {
	var lastErr error = nil

	for _, server := range self.Servers {
		err := server.Stop()
		if err != nil {
			lastErr = err
		}
	}

	self.Servers = make([]*MulticastServer, 0)

	return lastErr
}
