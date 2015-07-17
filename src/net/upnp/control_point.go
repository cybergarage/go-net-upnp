// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"

	"net/upnp/log"
	"net/upnp/ssdp"
	"net/upnp/util"
)

// A ControlPointListener represents a listener for ControlPoint.
type ControlPointListener interface {
	ssdp.MulticastListener
	ssdp.UnicastListener
}

// A ControlPoint represents a ControlPoint.
type ControlPoint struct {
	*sync.Mutex

	Port     int
	SearchMX int

	rootDeviceUdnMap    DeviceUdnMap
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

	cp.SearchMX = ControlPointDefaultSearchMX

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

	self.ssdpUcastServerList.Listener = self
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
	port := rand.Intn(ControlPointDefaultPortRange) + ControlPointDefaultPortBase
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
	return self.ssdpUcastServerList.Search(st, self.SearchMX)
}

// SearchRootDevice sends a M-SEARCH request for root devices.
func (self *ControlPoint) SearchRootDevice() error {
	return self.Search(ssdp.ROOT_DEVICE)
}

// GetRootDevices returns found root devices.
func (self *ControlPoint) GetRootDevices() []*Device {
	self.Lock()

	devCnt := len(self.rootDeviceUdnMap)
	devs := make([]*Device, devCnt)
	n := 0
	for _, dev := range self.rootDeviceUdnMap {
		devs[n] = dev
		n++
	}

	self.Unlock()

	return devs
}

// GetRootDevicesByType returns found root devices of the specified device type.
func (self *ControlPoint) GetRootDevicesByType(deviceType string) []*Device {
	devs := make([]*Device, 0)

	self.Lock()

	for _, dev := range self.rootDeviceUdnMap {
		if strings.Index(dev.DeviceType, deviceType) < 0 {
			continue
		}
		devs = append(devs, dev)
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
	defer self.Unlock()

	if self.rootDeviceUdnMap.HasDevice(dev) {
		log.Trace(fmt.Sprintf("device (%s) is already added", dev.UDN))
		return false
	}

	err := dev.LoadServiceDescriptions()
	if err != nil {
		return false
	}

	ok := self.rootDeviceUdnMap.AddDevice(dev)

	if ok {
		log.Trace(fmt.Sprintf("device (%s) is added", dev.UDN))
	}

	return ok
}

func getFromToMessageFromSSDPPacket(req *ssdp.Packet) string {
	fromAddr := req.From.String()
	toAddr := ""
	ifAddr, err := util.GetInterfaceAddress(req.Interface)
	if err == nil {
		toAddr = ifAddr
	}

	return fmt.Sprintf("(%s -> %s)", fromAddr, toAddr)
}

func (self *ControlPoint) DeviceNotifyReceived(ssdpReq *ssdp.Request) {
	usn, _ := ssdpReq.GetUSN()
	log.Trace(fmt.Sprintf("notiry req : %s %s", usn, getFromToMessageFromSSDPPacket(ssdpReq.Packet)))

	if ssdpReq.IsRootDevice() {
		newDev, err := NewDeviceFromSSDPRequest(ssdpReq)
		if err == nil {
			self.addDevice(newDev)
		} else {
			log.Warn(err)
		}
	}

	if self.Listener != nil {
		self.Listener.DeviceNotifyReceived(ssdpReq)
	}
}

func (self *ControlPoint) DeviceSearchReceived(ssdpReq *ssdp.Request) {
	st, _ := ssdpReq.GetST()
	log.Trace(fmt.Sprintf("search req : %s %s", st, getFromToMessageFromSSDPPacket(ssdpReq.Packet)))

	if self.Listener != nil {
		self.Listener.DeviceSearchReceived(ssdpReq)
	}
}

func (self *ControlPoint) DeviceResponseReceived(ssdpRes *ssdp.Response) {
	url, _ := ssdpRes.GetLocation()
	log.Trace(fmt.Sprintf("search res : %s %s", url, getFromToMessageFromSSDPPacket(ssdpRes.Packet)))

	newDev, err := NewDeviceFromSSDPResponse(ssdpRes)
	if err == nil {
		self.addDevice(newDev)
	} else {
		log.Warn(err)
	}

	if self.Listener != nil {
		self.Listener.DeviceResponseReceived(ssdpRes)
	}
}
