// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"fmt"
	"io"
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
	errorDeviceBadURLBaseAndLocationURL = "URLBase and location url are invalid ('%s', '%s'). Couldn't get an absolute URL ('%s')"
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

	devDescBytes, err := io.ReadAll(res.Body)
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
	return &root.Device, nil
}

// SetLocationURL set a location URL of SSDP packet.
func (dev *Device) SetLocationURL(url string) error {
	dev.LocationURL = url
	return nil
}

// CreateLocationURL returns a location URL for SSDP packet.
func (dev *Device) createLocationURLForAddress(addr string) (*url.URL, error) {
	locationBase := fmt.Sprintf("%s://%s:%d", DeviceProtocol, addr, dev.Port)
	url, err := util.GetAbsoluteURLFromBaseAndPath(locationBase, dev.DescriptionURL)
	if err != nil {
		return nil, err
	}
	return url, nil
}

// HasDeviceType returns true if the device or the embedded device type is the specified type, otherwise false.
func (dev *Device) HasDeviceType(deviceType string) bool {
	if dev.DeviceType == deviceType {
		return true
	}

	for n := 0; n < len(dev.DeviceList.Devices); n++ {
		dev := &dev.DeviceList.Devices[n]
		if dev.HasDeviceType(deviceType) {
			return true
		}
	}

	return false
}

// HasServiceType returns true if the device or the embedded device has the specified service type, otherwise false.
func (dev *Device) HasServiceType(serviceType string) bool {
	for n := 0; n < len(dev.ServiceList.Services); n++ {
		service := &dev.ServiceList.Services[n]
		if service.ServiceType == serviceType {
			return true
		}
	}

	for n := 0; n < len(dev.DeviceList.Devices); n++ {
		dev := &dev.DeviceList.Devices[n]
		if dev.HasServiceType(serviceType) {
			return true
		}
	}

	return false
}

// GetRootDevice returns the root device.
func (dev *Device) GetRootDevice() *Device {
	rootDev := dev
	for rootDev.ParentDevice != nil {
		rootDev = rootDev.ParentDevice
	}
	return rootDev
}

// LoadDescriptionBytes loads a device description string.
func (dev *Device) LoadDescriptionBytes(descBytes []byte) error {
	err := xml.Unmarshal(descBytes, dev)
	if err != nil {
		return err
	}
	return nil
}

// DescriptionString returns a descrition string.
func (dev *Device) DescriptionString() (string, error) {
	root, err := NewDeviceDescriptionRootFromDevice(dev)
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
func (dev *Device) LoadServiceDescriptions() error {
	var lastErr error

	for n := 0; n < len(dev.ServiceList.Services); n++ {
		service := &dev.ServiceList.Services[n]
		lastErr = service.LoadDescriptionFromSCPDURL()
	}

	// Embedded devices

	for n := 0; n < len(dev.DeviceList.Devices); n++ {
		dev := &dev.DeviceList.Devices[n]
		lastErr = dev.LoadServiceDescriptions()
	}

	return lastErr
}

// SetUDN sets a the specified UUID with a prefix.
func (dev *Device) SetUDN(uuid string) error {
	dev.UDN = fmt.Sprintf("%s%s", DeviceUUIDPrefix, uuid)
	return nil
}

// GetEmbeddedDevices returns all embedded devices
func (dev *Device) GetEmbeddedDevices() []*Device {
	devCnt := len(dev.DeviceList.Devices)
	devs := make([]*Device, devCnt)
	for n := 0; n < devCnt; n++ {
		devs[n] = &dev.DeviceList.Devices[n]
	}
	return devs
}

// GetEmbeddedDeviceByType returns a embedded device by the specified deviceType
func (dev *Device) GetEmbeddedDeviceByType(deviceType string) (*Device, error) {
	devCnt := len(dev.DeviceList.Devices)
	for n := 0; n < devCnt; n++ {
		dev := &dev.DeviceList.Devices[n]
		if dev.DeviceType == deviceType {
			return dev, nil
		}
	}
	return nil, fmt.Errorf(errorDeviceEmbeddedDeviceNotFound, deviceType)
}

// GetServices returns all services
func (dev *Device) GetServices() []*Service {
	servicCnt := len(dev.ServiceList.Services)
	services := make([]*Service, servicCnt)
	for n := 0; n < servicCnt; n++ {
		services[n] = &dev.ServiceList.Services[n]
	}
	return services
}

// GetServiceByType returns a service by the specified serviceType
func (dev *Device) GetServiceByType(serviceType string) (*Service, error) {
	for n := 0; n < len(dev.ServiceList.Services); n++ {
		service := &dev.ServiceList.Services[n]
		if service.ServiceType == serviceType {
			return service, nil
		}
	}
	return nil, fmt.Errorf(errorDeviceServiceNotFound, serviceType)
}

// GetServiceByID returns a service by the specified serviceId
func (dev *Device) GetServiceByID(serviceID string) (*Service, error) {
	for n := 0; n < len(dev.ServiceList.Services); n++ {
		service := &dev.ServiceList.Services[n]
		if service.ServiceID == serviceID {
			return service, nil
		}
	}
	return nil, fmt.Errorf(errorDeviceServiceNotFound, serviceID)
}

// GetServiceByControlURL returns a service by the specified control URL
func (dev *Device) GetServiceByControlURL(ctrlURL string) (*Service, error) {
	for n := 0; n < len(dev.ServiceList.Services); n++ {
		service := &dev.ServiceList.Services[n]
		if service.ControlURL == ctrlURL {
			return service, nil
		}
	}
	return nil, fmt.Errorf(errorDeviceServiceNotFound, ctrlURL)
}

// GetServiceByEventSubURL returns a service by the specified event subscription URL
func (dev *Device) GetServiceByEventSubURL(eventURL string) (*Service, error) {
	for n := 0; n < len(dev.ServiceList.Services); n++ {
		service := &dev.ServiceList.Services[n]
		if service.EventSubURL == eventURL {
			return service, nil
		}
	}
	return nil, fmt.Errorf(errorDeviceServiceNotFound, eventURL)
}

func (dev *Device) reviseParentObject() error {
	for n := 0; n < len(dev.ServiceList.Services); n++ {
		service := &dev.ServiceList.Services[n]
		service.ParentDevice = dev
	}

	for n := 0; n < len(dev.ServiceList.Services); n++ {
		service := &dev.ServiceList.Services[n]
		service.reviseParentObject()
	}

	// Embedded devices

	for n := 0; n < len(dev.DeviceList.Devices); n++ {
		dev := &dev.DeviceList.Devices[n]
		dev.ParentDevice = dev
		dev.reviseParentObject()
	}

	return nil
}

// TODO : Support embedded devices
func (dev *Device) reviseDescription() error {
	// check descriptionURL
	if len(dev.DescriptionURL) == 0 {
		dev.DescriptionURL = DeviceDefaultDescriptionURL
	}

	// check UUID
	if len(dev.UDN) == 0 {
		dev.SetUDN(util.CreateUUID())
	}

	// check description URLs in the service
	for n := 0; n < len(dev.ServiceList.Services); n++ {
		service := &dev.ServiceList.Services[n]
		service.reviseDescription()
	}

	return nil
}

// selectAvailableInterfaceForAddr return a interface from the specified address.
func (dev *Device) selectAvailableInterfaceForAddr(fromAddr string) (string, error) {
	ifi, err := util.GetAvailableInterfaceForAddr(fromAddr)
	if err != nil {
		return "", err
	}

	ifAddr, err := util.GetInterfaceAddress(ifi)
	if err != nil {
		return "", err
	}

	return ifAddr, nil
}

// GetAbsoluteURL return a absolute URL of the specified path using URLBase or LocationURL.
func (dev *Device) GetAbsoluteURL(path string) (*url.URL, error) {
	rootDev := dev.GetRootDevice()

	if 0 < len(rootDev.URLBase) {
		url, err := util.GetAbsoluteURLFromBaseAndPath(rootDev.URLBase, path)
		if err == nil {
			return url, nil
		}
	}

	if 0 < len(rootDev.LocationURL) {
		locationURL, err := url.Parse(rootDev.LocationURL)
		if err != nil {
			return nil, fmt.Errorf(errorDeviceBadLocationURL, rootDev.LocationURL)
		}
		baseLocation := locationURL.Scheme + "://" + locationURL.Host
		url, err := util.GetAbsoluteURLFromBaseAndPath(baseLocation, path)
		if err == nil {
			return url, nil
		}
	}

	url, err := util.GetAbsoluteURLFromBaseAndPath("", path)
	if err != nil {
		return nil, fmt.Errorf(errorDeviceBadURLBaseAndLocationURL, rootDev.URLBase, rootDev.LocationURL, path)
	}

	return url, nil
}

// Start starts this control point.
func (dev *Device) StartWithPort(port int) error {
	err := dev.reviseParentObject()
	if err != nil {
		return err
	}

	err = dev.reviseDescription()
	if err != nil {
		return err
	}

	dev.ssdpMcastServerList = ssdp.NewMulticastServerList()
	dev.ssdpMcastServerList.Listener = dev
	err = dev.ssdpMcastServerList.Start()
	if err != nil {
		dev.Stop()
		return err
	}

	dev.httpServer = http.NewServer()
	dev.httpServer.Listener = dev
	err = dev.httpServer.Start(port)
	if err != nil {
		dev.Stop()
		return err
	}

	dev.Port = port

	return nil
}

// Start starts this control point.
func (dev *Device) Start() error {
	port := rand.Intn(DeviceDefaultPortRange) + DeviceDefaultPortBase
	return dev.StartWithPort(port)
}

// Stop stops this control point.
func (dev *Device) Stop() error {
	var lastErr error

	err := dev.ssdpMcastServerList.Stop()
	if err != nil {
		lastErr = err
	}
	dev.ssdpMcastServerList = nil

	err = dev.httpServer.Stop()
	if err != nil {
		lastErr = err
	}
	dev.httpServer = nil

	return lastErr
}
