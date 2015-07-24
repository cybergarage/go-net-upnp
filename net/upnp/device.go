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

	"github.com/cybergarage/go-net-upnp/net/upnp/http"
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
	"github.com/cybergarage/go-net-upnp/net/upnp/util"
)

// A DeviceHTTPListener represents a listener for HTTP requests.
type DeviceHTTPListener interface {
	http.RequestListener
}

// A DeviceSSDPListener represents a listener for SSDP requests.
type DeviceSSDPListener interface {
	ssdp.MulticastListener
}

// A DeviceActionListener represents a listener for action request.
type DeviceActionListener interface {
	ActionRequestReceived(*Action) Error
}

// A DeviceListener represents a listener for all requests.
type DeviceListener interface {
	DeviceHTTPListener
	DeviceSSDPListener
	DeviceActionListener
}

// A Device represents a UPnP device.
type Device struct {
	*DeviceDescription
	SpecVersion SpecVersion `xml:"-"`
	URLBase     string      `xml:"-"`

	ParentDevice   *Device              `xml:"-"`
	Port           int                  `xml:"-"`
	HTTPListener   DeviceHTTPListener   `xml:"-"`
	SSDPListener   DeviceSSDPListener   `xml:"-"`
	ActionListener DeviceActionListener `xml:"-"`
	LocationURL    string               `xml:"-"`
	DescriptionURL string               `xml:"-"`

	ssdpMcastServerList *ssdp.MulticastServerList `xml:"-"`
	httpServer          *http.Server              `xml:"-"`
}

const (
	errorDeviceServiceNotFound          = "service (%s) is not found"
	errorDeviceEmbeddedDeviceNotFound   = "embedded device (%s) is not found"
	errorDeviceBadLocationURL           = "location url is invalid (%s)"
	errorDeviceBadUrlBaseAndLocationURL = "URLBase and location url are invalid ('%s', '%s'). Couldn't get an absolute URL ('%s')"
	errorDeviceBadDescriptionURL        = "DescriptionURL (%s) is bad response (%d)"
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
		return nil, fmt.Errorf(errorDeviceBadDescriptionURL, descURL, res.StatusCode)
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

// CreateLocationURL returns a location URL for SSDP packet.
func (self *Device) createLocationURLForAddress(addr string) (*url.URL, error) {
	locationBase := fmt.Sprintf("%s://%s:%d", DeviceProtocol, addr, self.Port)
	url, err := util.GetAbsoluteURLFromBaseAndPath(locationBase, self.DescriptionURL)
	if err != nil {
		return nil, err
	}
	return url, nil
}

// HasDeviceType returns true if the device or the embedded device type is the specified type, otherwise false.
func (self *Device) HasDeviceType(deviceType string) bool {
	if self.DeviceType == deviceType {
		return true
	}

	for n := 0; n < len(self.DeviceList.Devices); n++ {
		dev := &self.DeviceList.Devices[n]
		if dev.HasDeviceType(deviceType) {
			return true
		}
	}

	return false
}

// HasServiceType returns true if the device or the embedded device has the specified service type, otherwise false.
func (self *Device) HasServiceType(serviceType string) bool {
	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		if service.ServiceType == serviceType {
			return true
		}
	}

	for n := 0; n < len(self.DeviceList.Devices); n++ {
		dev := &self.DeviceList.Devices[n]
		if dev.HasServiceType(serviceType) {
			return true
		}
	}

	return false
}

// GetRootDevice returns the root device.
func (self *Device) GetRootDevice() *Device {
	rootDev := self
	for rootDev.ParentDevice != nil {
		rootDev = rootDev.ParentDevice
	}
	return rootDev
}

// LoadDescriptionBytes loads a device description string.
func (self *Device) LoadDescriptionBytes(descBytes []byte) error {
	err := xml.Unmarshal(descBytes, self)
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

	descBytes, err := xml.MarshalIndent(root, "", xmlMarshallIndent)
	if err != nil {
		return "", err
	}

	return string(descBytes), nil
}

// LoadServiceDescriptions loads service descriptions.
func (self *Device) LoadServiceDescriptions() error {
	var lastErr error

	for n := 0; n < len(self.ServiceList.Services); n++ {
		service := &self.ServiceList.Services[n]
		lastErr = service.LoadDescriptionFromSCPDURL()
	}

	// Embedded devices

	for n := 0; n < len(self.DeviceList.Devices); n++ {
		dev := &self.DeviceList.Devices[n]
		lastErr = dev.LoadServiceDescriptions()
	}

	return lastErr
}

// SetUDN sets a the specified UUID with a prefix.
func (self *Device) SetUDN(uuid string) error {
	self.UDN = fmt.Sprintf("%s%s", DeviceUUIDPrefix, uuid)
	return nil
}

// GetEmbeddedDevices returns all embedded devices
func (self *Device) GetEmbeddedDevices() []*Device {
	devCnt := len(self.DeviceList.Devices)
	devs := make([]*Device, devCnt)
	for n := 0; n < devCnt; n++ {
		devs[n] = &self.DeviceList.Devices[n]
	}
	return devs
}

// GetEmbeddedDeviceByType returns a embedded device by the specified deviceType
func (self *Device) GetEmbeddedDeviceByType(deviceType string) (*Device, error) {
	devCnt := len(self.DeviceList.Devices)
	for n := 0; n < devCnt; n++ {
		dev := &self.DeviceList.Devices[n]
		if dev.DeviceType == deviceType {
			return dev, nil
		}
	}
	return nil, fmt.Errorf(errorDeviceEmbeddedDeviceNotFound, deviceType)
}

// GetServices returns all services
func (self *Device) GetServices() []*Service {
	servicCnt := len(self.ServiceList.Services)
	services := make([]*Service, servicCnt)
	for n := 0; n < servicCnt; n++ {
		services[n] = &self.ServiceList.Services[n]
	}
	return services
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

	// Embedded devices

	for n := 0; n < len(self.DeviceList.Devices); n++ {
		dev := &self.DeviceList.Devices[n]
		dev.ParentDevice = self
		dev.reviseParentObject()
	}

	return nil
}

// TODO : Support embedded devices
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
	rootDev := self.GetRootDevice()

	if 0 < len(rootDev.URLBase) {
		url, err := util.GetAbsoluteURLFromBaseAndPath(rootDev.URLBase, path)
		if err == nil {
			return url, err
		}
	}

	if 0 < len(rootDev.LocationURL) {
		locationUrl, err := url.Parse(rootDev.LocationURL)
		if err != nil {
			return nil, fmt.Errorf(errorDeviceBadLocationURL, rootDev.LocationURL)
		}
		baseLocation := locationUrl.Scheme + "://" + locationUrl.Host
		url, err := util.GetAbsoluteURLFromBaseAndPath(baseLocation, path)
		if err == nil {
			return url, err
		}
	}

	url, err := util.GetAbsoluteURLFromBaseAndPath("", path)
	if err != nil {
		return nil, fmt.Errorf(errorDeviceBadUrlBaseAndLocationURL, rootDev.URLBase, rootDev.LocationURL, path)
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
