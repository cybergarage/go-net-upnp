// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A Service represents a UPnP service.
type Service struct {
	*ServiceDescription

	XMLName     xml.Name `xml:"service"`
	ServiceType string   `xml:"serviceType"`
	ServiceId   string   `xml:"serviceId"`
	SCPDURL     string   `xml:"SCPDURL"`
	ControlURL  string   `xml:"controlURL"`
	EventSubURL string   `xml:"eventSubURL"`

	Description *ServiceDescription
}

// A ServiceDescription represents a UPnP service description.
type ServiceDescription struct {
	XMLName           xml.Name          `xml:"scpd"`
	ServiceStateTable ServiceStateTable `xml:"serviceStateTable"`
	ActionList        ActionList        `xml:"actionList"`
}

// A ServiceList represents a UPnP serviceList.
type ServiceList struct {
	XMLName  xml.Name  `xml:"serviceList"`
	Services []Service `xml:"service"`
}

// NewService returns a new Service.
func NewService() *Service {
	service := &Service{}
	service.ServiceDescription = &ServiceDescription{}
	service.Description = &ServiceDescription{}
	return service
}

// LoadDescriptinString loads a device description string.
func (self *Service) LoadDescriptionString(desc string) error {
	err := xml.Unmarshal([]byte(desc), self.Description)
	if err != nil {
		return err
	}

	self.ServiceDescription = self.Description

	return nil
}
