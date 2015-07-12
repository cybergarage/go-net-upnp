// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

const (
	defaultServiceScpdURL    = "/service/scpd/%s.xml"
	defaultServiceControlURL = "/service/control/%s"
	defaultServiceEventURL   = "/service/event/%s"
)

const (
	errorServiceDescriptionNotFound = "action (%s) is not found. service (%s) description is null."
	errorServiceHanNoActions        = "action (%s) is not found. service (%s) has no actions."
	errorServiceActionNotFound      = "action (%s) is not found in the service (%s)"
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
	ParentDevice      *Device             `xml:"-"`
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
	self.description = &ServiceDescription{}

	err := xml.Unmarshal([]byte(desc), self.description)
	if err != nil {
		return err
	}

	self.ServiceStateTable = &self.description.ServiceStateTable
	self.ActionList = &self.description.ActionList

	err = self.reviseParentObject()
	if err != nil {
		return err
	}

	return nil
}

// DescriptionString returns a descrition string.
func (self *Service) DescriptionString() (string, error) {
	descBytes, err := xml.MarshalIndent(self.description, "", XmlMarshallIndent)
	if err != nil {
		return "", err
	}

	return string(descBytes), nil
}

func (self *Service) getShortServiceType() string {
	serviceTypes := strings.Split(self.ServiceType, ":")
	if len(serviceTypes) <= 1 {
		return self.ServiceId
	}
	return serviceTypes[len(serviceTypes)-2]
}

func (self *Service) reviseParentObject() error {
	if self.ActionList != nil {
		for n := 0; n < len(self.ActionList.Actions); n++ {
			action := &self.ActionList.Actions[n]
			action.ParentService = self
			action.reviseParentObject()
		}
	}

	if self.ServiceStateTable != nil {
		for n := 0; n < len(self.ServiceStateTable.StateVariables); n++ {
			statVar := &self.ServiceStateTable.StateVariables[n]
			statVar.ParentService = self
		}
	}

	return nil
}

func (self *Service) reviseDescription() error {
	shortServiceId := self.getShortServiceType()

	// check description URLs

	if len(self.SCPDURL) <= 0 {
		self.SCPDURL = fmt.Sprintf(defaultServiceScpdURL, shortServiceId)
	}

	if len(self.ControlURL) <= 0 {
		self.ControlURL = fmt.Sprintf(defaultServiceControlURL, shortServiceId)
	}

	if len(self.EventSubURL) <= 0 {
		self.EventSubURL = fmt.Sprintf(defaultServiceEventURL, shortServiceId)
	}

	return nil
}

func (self *Service) isDescriptionURL(path string) bool {
	if path == self.SCPDURL {
		return true
	}
	return false
}

func (self *Service) isControlURL(path string) bool {
	if path == self.ControlURL {
		return true
	}
	return false
}

func (self *Service) isEventSubURL(path string) bool {
	if path == self.EventSubURL {
		return true
	}
	return false
}

// GetActionByName returns an action by the specified name
func (self *Service) GetActionByName(name string) (*Action, error) {
	if self.description == nil {
		return nil, errors.New(fmt.Sprintf(errorServiceDescriptionNotFound, name, self.ServiceType))
	}

	if len(self.ActionList.Actions) <= 0 {
		return nil, errors.New(fmt.Sprintf(errorServiceHanNoActions, name, self.ServiceType))
	}

	for n := 0; n < len(self.ActionList.Actions); n++ {
		action := &self.ActionList.Actions[n]
		if action.Name == name {
			return action, nil
		}
	}
	return nil, errors.New(fmt.Sprintf(errorServiceActionNotFound, name, self.ServiceType))
}
