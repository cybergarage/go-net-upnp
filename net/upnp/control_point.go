// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/cybergarage/go-net-upnp/net/upnp/log"
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
	"github.com/cybergarage/go-net-upnp/net/upnp/util"
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

	rootDeviceMap       *DeviceMap
	ssdpMcastServerList *ssdp.MulticastServerList
	ssdpUcastServerList *ssdp.UnicastServerList
	Listener            ControlPointListener
}

// NewControlPoint returns a new ControlPoint.
func NewControlPoint() *ControlPoint {
	cp := &ControlPoint{}

	cp.Mutex = &sync.Mutex{}
	cp.rootDeviceMap = NewDeviceMap()
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

	devs := self.rootDeviceMap.GetAllDevices()

	self.Unlock()

	return devs
}

// GetRootDevicesByType returns found root devices of the specified device type.
func (self *ControlPoint) GetRootDevicesByType(deviceType string) []*Device {
	self.Lock()

	devs := self.rootDeviceMap.GetDevicesByType(deviceType)

	self.Unlock()

	return devs
}

// FindDeviceByTypeAndUDN returns a devices of the specified deviceType and UDN
func (self *ControlPoint) FindDeviceByTypeAndUDN(deviceType string, udn string) (*Device, bool) {
	self.Lock()

	dev, ok := self.rootDeviceMap.FindDeviceByTypeAndUDN(deviceType, udn)

	self.Unlock()

	return dev, ok
}

// AddDevice adds a specified device.
func (self *ControlPoint) addDevice(dev *Device) (bool, error) {
	self.Lock()
	defer self.Unlock()

	if self.rootDeviceMap.HasDevice(dev) {
		log.Trace(fmt.Sprintf("device (%s, %s) is already added", dev.DeviceType, dev.UDN))
		return false, nil
	}

	err := dev.LoadServiceDescriptions()
	if err != nil {
		return false, err
	}

	ok := self.rootDeviceMap.AddDevice(dev)

	if ok {
		log.Trace(fmt.Sprintf("device (%s, %s) is added", dev.DeviceType, dev.UDN))
	}

	return ok, nil
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
			_, err = self.addDevice(newDev)
		}
		if err != nil {
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
		_, err = self.addDevice(newDev)
	}
	if err != nil {
		log.Warn(err)
	}

	if self.Listener != nil {
		self.Listener.DeviceResponseReceived(ssdpRes)
	}
}
