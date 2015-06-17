// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"net/upnp/ssdp"
)

// A ControlPoint represents a clinet.
type ControlPoint struct {
	RootDevices []Device
	SSDPServer  *ssdp.SSDPServer
}

// NewControlPoint returns a new Client.
func NewControlPoint() *ControlPoint {
	cp := &ControlPoint{}
	cp.RootDevices = make([]Device, 0)
	cp.SSDPServer = ssdp.NewSSDPServer()
	return cp
}

// Start starts this control point.
func (self *ControlPoint) Start() error {
	self.SSDPServer.Listener = self
	err := self.SSDPServer.Start()
	if err != nil {
		return err
	}
	return nil
}

// Stop stops this control point.
func (self *ControlPoint) Stop() error {
	err := self.SSDPServer.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (self *ControlPoint) DeviceNotifyReceived(ssdpPkt *ssdp.SSDPPacket) {
}
