// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"io/ioutil"
	"math/rand"

	"net/upnp/http"
	"net/upnp/ssdp"
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

	Port     int            `xml:"-"`
	Listener DeviceListener `xml:"-"`

	Description         *DeviceDescription        `xml:"-"`
	ssdpMcastServerList *ssdp.MulticastServerList `xml:"-"`
	httpServer          *http.Server              `xml:"-"`
}

// NewDevice returns a new Device.
func NewDevice() *Device {
	dev := &Device{}

	dev.DeviceDescription = &DeviceDescription{}
	dev.Description = &DeviceDescription{}
	dev.ssdpMcastServerList = ssdp.NewMulticastServerList()
	dev.httpServer = http.NewServer()

	return dev
}

// NewDeviceFromSSDPRequest returns a device from the specified SSDP packet
func NewDeviceFromSSDPRequest(ssdpReq *ssdp.Request) (*Device, error) {

	descURL, err := ssdpReq.GetLocation()
	if err != nil {
		return nil, err
	}

	return NewDeviceFromDescriptionURL(descURL)
}

// NewDeviceFromDescriptionURL returns a device from the specified URL
func NewDeviceFromDescriptionURL(descURL string) (*Device, error) {
	res, err := http.Get(descURL)
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
	rootDev.DeviceDescription = rootDev.Description

	return rootDev, nil
}

// Start starts this control point.
func (self *Device) StartWithPort(port int) error {
	self.ssdpMcastServerList.Listener = self
	err := self.ssdpMcastServerList.Start()
	if err != nil {
		self.Stop()
		return err
	}

	self.httpServer.Listener = self
	err = self.httpServer.Start(port)
	if err != nil {
		self.Stop()
		return err
	}

	self.Port = port

	return nil
}

// LoadDescriptinString loads a device description string.
func (self *Device) LoadDescriptionString(desc string) error {
	err := xml.Unmarshal([]byte(desc), self.Description)
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

	descBytes, err := xml.MarshalIndent(root, "", "  ")
	if err != nil {
		return "", err
	}

	return string(descBytes), nil
}

// Start starts this control point.
func (self *Device) Start() error {
	port := rand.Intn(DEVICE_DEFAULT_PORT_RANGE) + DEVICE_DEFAULT_PORT_BASE
	return self.StartWithPort(port)
}

// Stop stops this control point.
func (self *Device) Stop() error {
	err := self.ssdpMcastServerList.Stop()
	if err != nil {
		return err
	}
	err = self.httpServer.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (self *Device) DeviceNotifyReceived(ssdpReq *ssdp.Request) {
	self.handleNotifyRequest(ssdpReq)

	if self.Listener != nil {
		self.Listener.DeviceNotifyReceived(ssdpReq)
	}
}

func (self *Device) DeviceSearchReceived(ssdpReq *ssdp.Request) {
	if self.Listener != nil {
		self.Listener.DeviceSearchReceived(ssdpReq)
	}

}

func (self *Device) HTTPRequestReceived(httpReq *http.Request, httpRes http.ResponseWriter) {
	if self.Listener != nil {
		self.Listener.HTTPRequestReceived(httpReq, httpRes)
		return
	}

	httpRes.WriteHeader(http.StatusInternalServerError)
}

func (self *Device) handleNotifyRequest(ssdpReq *ssdp.Request) {
	if !ssdpReq.IsDiscover() {
		return
	}
}
