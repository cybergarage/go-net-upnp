// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"io/ioutil"

	"net/upnp/http"
	"net/upnp/ssdp"
)

// A DeviceDescriptionRoot represents a root UPnP device description.
type DeviceDescriptionRoot struct {
	XMLName     xml.Name    `xml:"root"`
	SpecVersion SpecVersion `xml:"specVersion"`
	URLBase     string      `xml:"URLBase"`
	Device      Device      `xml:"device"`
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
	root := DeviceDescriptionRoot{}
	err := xml.Unmarshal([]byte(devDesc), &root)
	if err != nil {
		return nil, err
	}

	rootDev := &root.Device
	rootDev.SpecVersion = rootDev.SpecVersion
	rootDev.URLBase = rootDev.URLBase
	rootDev.DeviceDescription = rootDev.Description

	return rootDev, nil
}
