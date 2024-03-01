// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"github.com/cybergarage/go-net-upnp/net/upnp/util"
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
func (servers *MulticastServerList) Start() error {
	err := servers.Stop()
	if err != nil {
		return err
	}

	ifis, err := util.GetAvailableInterfaces()
	if err != nil {
		return err
	}

	var lastErr error

	servers.Servers = make([]*MulticastServer, len(ifis))
	for n, ifi := range ifis {
		server := NewMulticastServer()
		server.Listener = servers.Listener
		err := server.Start(ifi)
		if err != nil {
			lastErr = err
		}
		servers.Servers[n] = server
	}

	return lastErr
}

// Stop stops this server.
func (servers *MulticastServerList) Stop() error {
	var lastErr error

	for _, server := range servers.Servers {
		err := server.Stop()
		if err != nil {
			lastErr = err
		}
	}

	servers.Servers = make([]*MulticastServer, 0)

	return lastErr
}
