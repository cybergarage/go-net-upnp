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

// A Device represents a clinet.
type Device struct {
	XMLName          xml.Name  `xml:"device"`
	DeviceType       string    `xml:"deviceType"`
	FriendlyName     string    `xml:"friendlyName"`
	Manufacturer     string    `xml:"manufacturer"`
	ManufacturerURL  string    `xml:"manufacturerURL"`
	ModelDescription string    `xml:"modelDescription"`
	ModelName        string    `xml:"modelName"`
	ModelNumber      string    `xml:"modelNumber"`
	ModelURL         string    `xml:"modelURL"`
	SerialNumber     string    `xml:"serialNumber"`
	UDN              string    `xml:"UDN"`
	UPC              string    `xml:"UPC"`
	PresentationURL  string    `xml:"presentationURL"`
	IconList         []Icon    `xml:"iconList"`
	ServiceList      []Service `xml:"serviceList"`
	DeviceList       []Device  `xml:"deviceList"`

	Port     int
	Listener DeviceListener

	ssdpMcastServer *ssdp.MulticastServer
	httpServer      *http.Server
}

// NewDevice returns a new Device.
func NewDevice() *Device {
	dev := &Device{}
	dev.ssdpMcastServer = ssdp.NewMulticastServer()
	dev.httpServer = http.NewServer()
	return dev
}

// Start starts this control point.
func (self *Device) StartWithPort(port int) error {
	self.ssdpMcastServer.Listener = self
	err := self.ssdpMcastServer.Start()
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

// Start starts this control point.
func (self *Device) Start() error {
	port := rand.Intn(DEVICE_DEFAULT_PORT_RANGE) + DEVICE_DEFAULT_PORT_BASE
	return self.StartWithPort(port)
}

// Stop stops this control point.
func (self *Device) Stop() error {
	err := self.ssdpMcastServer.Stop()
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
