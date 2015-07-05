// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"math/rand"
	"sync"

	"net/upnp/ssdp"
)

// A ControlPointListener represents a listener for ControlPoint.
type ControlPointListener interface {
	ssdp.MulticastListener
	ssdp.UnicastListener
}

// A ControlPoint represents a ControlPoint.
type ControlPoint struct {
	*sync.Mutex

	Port int

	rootDeviceUdnMap    *DeviceUdnMap
	ssdpMcastServerList *ssdp.MulticastServerList
	ssdpUcastServerList *ssdp.UnicastServerList
	Listener            ControlPointListener
}

// NewControlPoint returns a new ControlPoint.
func NewControlPoint() *ControlPoint {
	cp := &ControlPoint{}

	cp.Mutex = &sync.Mutex{}
	cp.rootDeviceUdnMap = NewDeviceUdnMap()
	cp.ssdpMcastServerList = ssdp.NewMulticastServerList()
	cp.ssdpUcastServerList = ssdp.NewUnicastServerList()

	return cp
}

// Start starts this control point.
func (self *ControlPoint) StartWithPort(port int) error {
	self.ssdpMcastServerList.Listener = self
	err := self.ssdpMcastServerList.Start()
	if err != nil {
		self.Stop()
		return err
	}

	err = self.ssdpUcastServerList.Start(port)
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
	err := self.ssdpMcastServerList.Stop()
	if err != nil {
		return err
	}
	return nil
}

// Search sends a M-SEARCH request of the specified ST.
func (self *ControlPoint) Search(st string) error {
	return self.ssdpUcastServerList.Search(st)
}

// SearchRootDevice sends a M-SEARCH request for root devices.
func (self *ControlPoint) SearchRootDevice() error {
	return self.Search(ssdp.ROOT_DEVICE)
}

// GetRootDevices returns found root devices.
func (self *ControlPoint) GetRootDevices() []*Device {
	self.Lock()

	devCnt := len(*self.rootDeviceUdnMap)
	devs := make([]*Device, devCnt)
	n := 0
	for _, dev := range *self.rootDeviceUdnMap {
		devs[n] = dev
		n++
	}
	self.Unlock()

	return devs
}

// FindDeviceByUSN returns a devices of the specified UDN
func (self *ControlPoint) FindDeviceByUDN(udn string) (*Device, bool) {
	self.Lock()
	dev, ok := self.rootDeviceUdnMap.FindDeviceByUDN(udn)
	self.Unlock()
	return dev, ok
}

// AddDevice adds a specified device.
func (self *ControlPoint) addDevice(dev *Device) bool {
	self.Lock()
	ok := self.rootDeviceUdnMap.AddDevice(dev)
	self.Unlock()
	return ok
}

func (self *ControlPoint) addDeviceFromSSDPPacket(ssdpReq *ssdp.Request) bool {
	usn, err := ssdpReq.GetUSN()
	if err != nil {
		return false
	}

	_, ok := self.FindDeviceByUDN(usn)
	if ok {
		return false
	}

	newDev, err := NewDeviceFromSSDPRequest(ssdpReq)
	if err != nil {
		return false
	}

	return self.addDevice(newDev)
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
