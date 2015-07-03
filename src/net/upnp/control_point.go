// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"math/rand"
	"net/upnp/ssdp"
)

// A ControlPointListener represents a listener for ControlPoint.
type ControlPointListener interface {
	ssdp.MulticastListener
	ssdp.UnicastListener
}

// A ControlPoint represents a ControlPoint.
type ControlPoint struct {
	Port            int
	RootDevices     []Device
	ssdpMcastServer *ssdp.MulticastServer
	ssdpUcastServer *ssdp.UnicastServer
	Listener        ControlPointListener
}

// NewControlPoint returns a new ControlPoint.
func NewControlPoint() *ControlPoint {
	cp := &ControlPoint{}
	cp.RootDevices = make([]Device, 0)
	cp.ssdpMcastServer = ssdp.NewMulticastServer()
	cp.ssdpUcastServer = ssdp.NewUnicastServer()
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

// Search sends a M-SEARCH request of the specified ST.
func (self *ControlPoint) Search(st string) error {
	ssdpReq, err := ssdp.NewSearchRequest(st)
	if err != nil {
		return err
	}

	ssdpSock, err := ssdp.NewMulticastSocket()
	if err != nil {
		return err
	}

	_, err = ssdpSock.Write(ssdpReq)

	return err
}

// SearchRootDevice sends a M-SEARCH request for root devices.
func (self *ControlPoint) SearchRootDevice() error {
	return self.Search(ssdp.ROOT_DEVICE)
}

func (self *ControlPoint) DeviceNotifyReceived(ssdpReq *ssdp.Request) {
	if self.Listener != nil {
		self.Listener.DeviceNotifyReceived(ssdpReq)
	}
}

func (self *ControlPoint) DeviceSearchReceived(ssdpReq *ssdp.Request) {
	if self.Listener != nil {
		self.Listener.DeviceSearchReceived(ssdpReq)
	}
}

func (self *ControlPoint) DeviceResponseReceived(ssdpRes *ssdp.Response) {
	if self.Listener != nil {
		self.Listener.DeviceResponseReceived(ssdpRes)
	}
}
