// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"

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
	errorDeviceServiceNotFound          = "service (%s) is not found"
	errorDeviceBadLocationURL           = "location url is invalid (%s)"
	errorDeviceBadUrlBaseAndLocationURL = "URLBase and location url are invalid ('%s', '%s'). Couldn't get an absolute URL ('%s')"
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
		return nil, err
	}

	dev.SetLocationURL(descURL)

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
func (self *Device) createLocationURLForAddress(addr string) (string, error) {
	url := fmt.Sprintf("%s://%s:%d%s", DeviceProtocol, addr, self.Port, self.DescriptionURL)
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
	return nil, fmt.Errorf(errorDeviceServiceNotFound, serviceType)
}

// GetServiceById returns a service by the specified serviceId
func (self *Device) GetServiceById(serviceId string) (*Service, error) {
	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		if service.ServiceId == serviceId {
			return service, nil
		}
	}
	return nil, fmt.Errorf(errorDeviceServiceNotFound, serviceId)
}

// GetServiceByControlURL returns a service by the specified control URL
func (self *Device) GetServiceByControlURL(ctrlURL string) (*Service, error) {
	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		if service.ControlURL == ctrlURL {
			return service, nil
		}
	}
	return nil, fmt.Errorf(errorDeviceServiceNotFound, ctrlURL)
}

// GetServiceByEventSubURL returns a service by the specified event subscription URL
func (self *Device) GetServiceByEventSubURL(eventURL string) (*Service, error) {
	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		if service.EventSubURL == eventURL {
			return service, nil
		}
	}
	return nil, fmt.Errorf(errorDeviceServiceNotFound, eventURL)
}

func (self *Device) reviseParentObject() error {
	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		service.ParentDevice = self
	}

	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		service.reviseParentObject()
	}

	return nil
}

func (self *Device) reviseDescription() error {
	// check descriptionURL
	if len(self.DescriptionURL) <= 0 {
		self.DescriptionURL = DeviceDefaultDescriptionURL
	}

	// check UUID
	if len(self.UDN) <= 0 {
		self.SetUDN(util.CreateUUID())
	}

	// check description URLs in the service
	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		service.reviseDescription()
	}

	return nil
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

// GetAbsoluteURL return a absoulte URL of the specified path using URLBase or LocationURL.
func (self *Device) GetAbsoluteURL(path string) (*url.URL, error) {
	if 0 < len(self.URLBase) {
		url, err := util.GetAbsoluteURLFromBaseAndPath(self.URLBase, path)
		if err == nil {
			return url, err
		}
	}

	if 0 < len(self.LocationURL) {
		locationUrl, err := url.Parse(self.LocationURL)
		if err != nil {
			return nil, fmt.Errorf(errorDeviceBadLocationURL, self.LocationURL)
		}
		baseLocation := locationUrl.Scheme + "://" + locationUrl.Host
		url, err := util.GetAbsoluteURLFromBaseAndPath(baseLocation, path)
		if err == nil {
			return url, err
		}
	}

	url, err := util.GetAbsoluteURLFromBaseAndPath("", path)
	if err != nil {
		return nil, fmt.Errorf(errorDeviceBadUrlBaseAndLocationURL, self.URLBase, self.LocationURL, path)
	}

	return url, nil
}

// Start starts this control point.
func (self *Device) StartWithPort(port int) error {
	err := self.reviseParentObject()
	if err != nil {
		return err
	}

	err = self.reviseDescription()
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
