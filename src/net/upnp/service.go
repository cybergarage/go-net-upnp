// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A Service represents a UPnP service.
type Service struct {
	XMLName     xml.Name `xml:"service"`
	ServiceType string   `xml:"serviceType"`
	ServiceId   string   `xml:"serviceId"`
	SCPDURL     string   `xml:"SCPDURL"`
	ControlURL  string   `xml:"controlURL"`
	EventSubURL string   `xml:"eventSubURL"`

	description       *ServiceDescription `xml:"-"`
	ServiceStateTable *ServiceStateTable  `xml:"-"`
	ActionList        *ActionList         `xml:"-"`
}

// NewService returns a new Service.
func NewService() *Service {
	service := &Service{}
	service.description = &ServiceDescription{}
	service.ServiceStateTable = &ServiceStateTable{}
	service.ActionList = &ActionList{}
	return service
}

// NewServiceFromDescription returns a service from the specified descrition string
func NewServiceFromDescription(serviceDesc string) (*Service, error) {
	service := NewService()

	err := service.LoadDescriptionString(serviceDesc)
	if err != nil {
		return nil, err
	}

	return service, err
}

// LoadDescriptinString loads a device description string.
func (self *Service) LoadDescriptionString(desc string) error {
	err := xml.Unmarshal([]byte(desc), self.description)
	if err != nil {
		return err
	}

	self.ServiceStateTable = &self.description.ServiceStateTable
	self.ActionList = &self.description.ActionList

	return nil
}

// DescriptionString returns a descrition string.
func (self *Service) DescriptionString() (string, error) {
	descBytes, err := xml.MarshalIndent(self.description, "", XML_MARSHALL_INDENT)
	if err != nil {
		return "", err
	}

	return string(descBytes), nil
}
