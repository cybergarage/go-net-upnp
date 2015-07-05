// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
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
	Description *DeviceDescription

	Port     int
	Listener DeviceListener

	ssdpMcastServerList *ssdp.MulticastServerList
	httpServer          *http.Server
}

// A DeviceDescription represents a UPnP device description.
type DeviceDescription struct {
	XMLName          xml.Name    `xml:"device"`
	DeviceType       string      `xml:"deviceType"`
	FriendlyName     string      `xml:"friendlyName"`
	Manufacturer     string      `xml:"manufacturer"`
	ManufacturerURL  string      `xml:"manufacturerURL"`
	ModelDescription string      `xml:"modelDescription"`
	ModelName        string      `xml:"modelName"`
	ModelNumber      string      `xml:"modelNumber"`
	ModelURL         string      `xml:"modelURL"`
	SerialNumber     string      `xml:"serialNumber"`
	UDN              string      `xml:"UDN"`
	UPC              string      `xml:"UPC"`
	PresentationURL  string      `xml:"presentationURL"`
	IconList         IconList    `xml:"iconList"`
	ServiceList      ServiceList `xml:"serviceList"`
	DeviceList       DeviceList  `xml:"deviceList"`
}

// A DeviceList represents a ServiceList.
type DeviceList struct {
	XMLName xml.Name `xml:"deviceList"`
	Devices []Device `xml:"device"`
}

// NewDevice returns a new Device.
func NewDevice() *Device {
	dev := &Device{}
	dev.Description = &DeviceDescription{}
	dev.ssdpMcastServerList = ssdp.NewMulticastServerList()
	dev.httpServer = http.NewServer()
	return dev
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
	return nil
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
