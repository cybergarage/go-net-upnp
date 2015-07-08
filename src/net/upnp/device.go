// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"

	"net/upnp/http"
	"net/upnp/ssdp"
	"net/upnp/util"
)

// A DeviceListener represents a listener for Device.
type DeviceListener interface {
	ssdp.MulticastListener
	http.RequestListener
}

// A Device represents a UPnP device.
type Device struct {
	*DeviceDescription
	SpecVersion SpecVersion `xml:"-"`
	URLBase     string      `xml:"-"`

	Port           int            `xml:"-"`
	Listener       DeviceListener `xml:"-"`
	LocationURL    string         `xml:"-"`
	DescriptionURL string         `xml:"-"`

	ssdpMcastServerList *ssdp.MulticastServerList `xml:"-"`
	httpServer          *http.Server              `xml:"-"`
}

const (
	errorDeviceServiceNotFound = "service (%s) is not found"
)

// NewDevice returns a new Device.
func NewDevice() *Device {
	dev := &Device{}

	dev.DeviceDescription = &DeviceDescription{}

	return dev
}

// NewDeviceFromSSDPRequest returns a device from the specified SSDP packet
func NewDeviceFromSSDPRequest(ssdpReq *ssdp.Request) (*Device, error) {

	descURL, err := ssdpReq.GetLocation()
	if err != nil {
		return nil, err
	}

	dev, err := NewDeviceFromDescriptionURL(descURL)

	if err != nil {
		dev.SetLocationURL(descURL)
	}

	return dev, err
}

// NewDeviceFromSSDPRequest returns a device from the specified SSDP packet
func NewDeviceFromSSDPResponse(ssdpRes *ssdp.Response) (*Device, error) {

	descURL, err := ssdpRes.GetLocation()
	if err != nil {
		return nil, err
	}

	dev, err := NewDeviceFromDescriptionURL(descURL)
	if err != nil {
		dev.SetLocationURL(descURL)
	}

	return dev, err
}

// NewDeviceFromDescriptionURL returns a device from the specified URL
func NewDeviceFromDescriptionURL(descURL string) (*Device, error) {
	res, err := http.Get(descURL)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	devDescBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return NewDeviceFromDescription(string(devDescBytes))
}

// NewDeviceFromDescription returns a device from the specified string
func NewDeviceFromDescription(devDesc string) (*Device, error) {
	root, err := NewDeviceDescriptionRootFromString(devDesc)
	if err != nil {
		return nil, err
	}

	rootDev := &root.Device
	rootDev.SpecVersion = rootDev.SpecVersion
	rootDev.URLBase = rootDev.URLBase

	return rootDev, nil
}

// SetLocationURL set a location URL of SSDP packet.
func (self *Device) SetLocationURL(url string) error {
	self.LocationURL = url
	return nil
}

// CreateLocationURL return a location URL for SSDP packet.
func (self *Device) CreateLocationURLForAddress(addr string) (string, error) {
	url := fmt.Sprint("%s://%s:%s:%s", DeviceProtocol, addr, self.Port, self.DescriptionURL)
	return url, nil
}

// LoadDescriptinString loads a device description string.
func (self *Device) LoadDescriptionString(desc string) error {
	err := xml.Unmarshal([]byte(desc), self)
	if err != nil {
		return err
	}

	self.DeviceDescription = self.DeviceDescription

	return nil
}

// DescriptionString returns a descrition string.
func (self *Device) DescriptionString() (string, error) {
	root, err := NewDeviceDescriptionRootFromDevice(self)
	if err != nil {
		return "", err
	}

	descBytes, err := xml.MarshalIndent(root, "", XmlMarshallIndent)
	if err != nil {
		return "", err
	}

	return string(descBytes), nil
}

// SetUDN sets a the specified UUID with a prefix.
func (self *Device) SetUDN(uuid string) error {
	self.UDN = fmt.Sprintf("%s%s", DeviceUUIDPrefix, uuid)
	return nil
}

// GetServiceByType returns a service by the specified serviceType
func (self *Device) GetServiceByType(serviceType string) (*Service, error) {
	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		if service.ServiceType == serviceType {
			return service, nil
		}
	}
	return nil, errors.New(fmt.Sprintf(errorDeviceServiceNotFound, serviceType))
}

// GetServiceById returns a service by the specified serviceId
func (self *Device) GetServiceById(serviceId string) (*Service, error) {
	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		if service.ServiceId == serviceId {
			return service, nil
		}
	}
	return nil, errors.New(fmt.Sprintf(errorDeviceServiceNotFound, serviceId))
}

func (self *Device) reviseDescription() error {
	// check UUID
	if len(self.UDN) <= 0 {
		self.SetUDN(util.CreateUUID())
	}

	return nil
}

// Start starts this control point.
func (self *Device) StartWithPort(port int) error {
	err := self.reviseDescription()
	if err != nil {
		return err
	}

	self.ssdpMcastServerList = ssdp.NewMulticastServerList()
	self.ssdpMcastServerList.Listener = self
	err = self.ssdpMcastServerList.Start()
	if err != nil {
		self.Stop()
		return err
	}

	self.httpServer = http.NewServer()
	self.httpServer.Listener = self
	err = self.httpServer.Start(port)
	if err != nil {
		self.Stop()
		return err
	}

	self.Port = port

	return nil
}

// Start starts this control point.
func (self *Device) Start() error {
	port := rand.Intn(DeviceDefaultPortRange) + DeviceDefaultPortBase
	return self.StartWithPort(port)
}

// Stop stops this control point.
func (self *Device) Stop() error {
	var lastErr error = nil

	err := self.ssdpMcastServerList.Stop()
	if err != nil {
		lastErr = err
	}
	self.ssdpMcastServerList = nil

	err = self.httpServer.Stop()
	if err != nil {
		lastErr = err
	}
	self.httpServer = nil

	return lastErr
}

// selectAvailableInterfaceForAddr return a interface from the specified address.
func (self *Device) selectAvailableInterfaceForAddr(fromAddr string) (string, error) {
	ifi, err := util.GetAvailableInterfaceForAddr(fromAddr)
	if err != nil {
		return "", nil
	}

	ifAddr, err := util.GetInterfaceAddress(ifi)
	if err != nil {
		return "", nil
	}

	return ifAddr, err
}
