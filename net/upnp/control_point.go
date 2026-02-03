// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/cybergarage/go-logger/log"
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

// StartWithPort starts this control point using the specified port.
func (ctrl *ControlPoint) StartWithPort(port int) error {
	ctrl.ssdpMcastServerList.Listener = ctrl
	err := ctrl.ssdpMcastServerList.Start()
	if err != nil {
		ctrl.Stop()
		return err
	}

	ctrl.ssdpUcastServerList.Listener = ctrl
	err = ctrl.ssdpUcastServerList.Start(port)
	if err != nil {
		ctrl.Stop()
		return err
	}

	ctrl.Port = port

	return nil
}

// Start starts this control point.
func (ctrl *ControlPoint) Start() error {
	port := rand.Intn(ControlPointDefaultPortRange) + ControlPointDefaultPortBase
	return ctrl.StartWithPort(port)
}

// Stop stops this control point.
func (ctrl *ControlPoint) Stop() error {
	err := ctrl.ssdpMcastServerList.Stop()
	if err != nil {
		return err
	}
	return nil
}

// Search sends a M-SEARCH request of the specified ST.
func (ctrl *ControlPoint) Search(st string) error {
	return ctrl.ssdpUcastServerList.Search(st, ctrl.SearchMX)
}

// SearchRootDevice sends a M-SEARCH request for root devices.
func (ctrl *ControlPoint) SearchRootDevice() error {
	return ctrl.Search(ssdp.RootDevice)
}

// GetRootDevices returns found root devices.
func (ctrl *ControlPoint) GetRootDevices() []*Device {
	ctrl.Lock()

	devs := ctrl.rootDeviceMap.GetAllDevices()

	ctrl.Unlock()

	return devs
}

// GetRootDevicesByType returns found root devices of the specified device type.
func (ctrl *ControlPoint) GetRootDevicesByType(deviceType string) []*Device {
	ctrl.Lock()

	devs := ctrl.rootDeviceMap.GetDevicesByType(deviceType)

	ctrl.Unlock()

	return devs
}

// FindDeviceByTypeAndUDN returns a devices of the specified deviceType and UDN.
func (ctrl *ControlPoint) FindDeviceByTypeAndUDN(deviceType string, udn string) (*Device, bool) {
	ctrl.Lock()

	dev, ok := ctrl.rootDeviceMap.FindDeviceByTypeAndUDN(deviceType, udn)

	ctrl.Unlock()

	return dev, ok
}

// AddDevice adds a specified device.
func (ctrl *ControlPoint) addDevice(dev *Device) (bool, error) {
	ctrl.Lock()
	defer ctrl.Unlock()

	if ctrl.rootDeviceMap.HasDevice(dev) {
		log.Tracef("device (%s, %s) is already added", dev.DeviceType, dev.UDN)
		return false, nil
	}

	err := dev.LoadServiceDescriptions()
	if err != nil {
		return false, err
	}

	ok := ctrl.rootDeviceMap.AddDevice(dev)

	if ok {
		log.Tracef("device (%s, %s) is added", dev.DeviceType, dev.UDN)
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

func (ctrl *ControlPoint) DeviceNotifyReceived(ssdpReq *ssdp.Request) {
	usn, _ := ssdpReq.GetUSN()
	log.Tracef("notiry req : %s %s", usn, getFromToMessageFromSSDPPacket(ssdpReq.Packet))

	if ssdpReq.IsRootDevice() {
		newDev, err := NewDeviceFromSSDPRequest(ssdpReq)
		if err == nil {
			_, err = ctrl.addDevice(newDev)
		}
		if err != nil {
			log.Warnf("%s", err.Error())
		}
	}

	if ctrl.Listener != nil {
		ctrl.Listener.DeviceNotifyReceived(ssdpReq)
	}
}

func (ctrl *ControlPoint) DeviceSearchReceived(ssdpReq *ssdp.Request) {
	st, _ := ssdpReq.GetST()
	log.Tracef("search req : %s %s", st, getFromToMessageFromSSDPPacket(ssdpReq.Packet))

	if ctrl.Listener != nil {
		ctrl.Listener.DeviceSearchReceived(ssdpReq)
	}
}

func (ctrl *ControlPoint) DeviceResponseReceived(ssdpRes *ssdp.Response) {
	url, _ := ssdpRes.GetLocation()
	log.Tracef("search res : %s %s", url, getFromToMessageFromSSDPPacket(ssdpRes.Packet))

	newDev, err := NewDeviceFromSSDPResponse(ssdpRes)
	if err == nil {
		_, err = ctrl.addDevice(newDev)
	}
	if err != nil {
		log.Warnf("%s", err.Error())
	}

	if ctrl.Listener != nil {
		ctrl.Listener.DeviceResponseReceived(ssdpRes)
	}
}
