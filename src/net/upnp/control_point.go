// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"math/rand"
	"net/upnp/ssdp"
)

// A ControlPoint represents a listener for ControlPoint.
type ControlPointListener interface {
	ssdp.SSDPMulticastListener
	ssdp.SSDPUnicastListener
}

// A ControlPoint represents a ControlPoint.
type ControlPoint struct {
	Port            int
	RootDevices     []Device
	ssdpMcastServer *ssdp.SSDPMulticastServer
	ssdpUcastServer *ssdp.SSDPUnicastServer
	Listener        ControlPointListener
}

// NewControlPoint returns a new ControlPoint.
func NewControlPoint() *ControlPoint {
	cp := &ControlPoint{}
	cp.RootDevices = make([]Device, 0)
	cp.ssdpMcastServer = ssdp.NewSSDPMulticastServer()
	cp.ssdpUcastServer = ssdp.NewSSDPUnicastServer()
	return cp
}

// Start starts this control point.
func (self *ControlPoint) StartWithPort(port int) error {
	self.ssdpMcastServer.Listener = self
	err := self.ssdpMcastServer.Start()
	if err != nil {
		self.Stop()
		return err
	}

	err = self.ssdpUcastServer.Start(port)
	if err != nil {
		self.Stop()
		return err
	}

	self.Port = port

	return nil
}

// Start starts this control point.
func (self *ControlPoint) Start() error {
	port := rand.Intn(CONTROLPOINT_DEFAULT_PORT_RANGE) + CONTROLPOINT_DEFAULT_PORT_BASE
	return self.StartWithPort(port)
}

// Stop stops this control point.
func (self *ControlPoint) Stop() error {
	err := self.ssdpMcastServer.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (self *ControlPoint) DeviceNotifyReceived(ssdpReq *ssdp.SSDPRequest) {
	if self.Listener != nil {
		self.Listener.DeviceNotifyReceived(ssdpReq)
	}
}

func (self *ControlPoint) DeviceResponseReceived(ssdpRes *ssdp.SSDPResponse) {
	if self.Listener != nil {
		self.Listener.DeviceResponseReceived(ssdpRes)
	}
}
