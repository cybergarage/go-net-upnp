// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"net/upnp/util"
)

// A UnicastServerList represents a packet of SSDP.
type UnicastServerList struct {
	Listener UnicastListener
	Servers  []*UnicastServer
}

// NewUnicastServerList returns a new UnicastServerList.
func NewUnicastServerList() *UnicastServerList {
	server := &UnicastServerList{}
	server.Servers = make([]*UnicastServer, 0)
	server.Listener = nil
	return server
}

// Start starts this server.
func (self *UnicastServerList) Start(port int) error {
	err := self.Stop()
	if err != nil {
		return err
	}

	addrs, err := util.GetAvailableInterfaceAddresses()
	if err != nil {
		return err
	}

	var lastErr error = nil

	self.Servers = make([]*UnicastServer, len(addrs))
	for n, addr := range addrs {
		server := NewUnicastServer()
		server.Listener = self.Listener
		err := server.Start(addr, port)
		if err != nil {
			lastErr = err
		}
		self.Servers[n] = server
	}

	return lastErr
}

// Stop stops this server.
func (self *UnicastServerList) Stop() error {
	var lastErr error = nil

	for _, server := range self.Servers {
		err := server.Stop()
		if err != nil {
			lastErr = err
		}
	}

	self.Servers = make([]*UnicastServer, 0)

	return lastErr
}

// Search sends a M-SEARCH request of the specified ST.
func (self *UnicastServerList) Search(st string) error {
	var lastErr error = nil

	for _, server := range self.Servers {
		err := server.Search(st)
		if err != nil {
			lastErr = err
		}
	}

	return lastErr
}
