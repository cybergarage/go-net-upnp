// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/cybergarage/go-net-upnp/net/upnp/http"
)

const (
	defaultServiceScpdURL    = "/service/scpd/%s.xml"
	defaultServiceControlURL = "/service/control/%s"
	defaultServiceEventURL   = "/service/event/%s"
)

const (
	errorServiceDescriptionNotFound = "service (%s) description is not found"
	errorServiceHasNoActions        = "service (%s) has no actions"
	errorServiceActionNotFound      = "action (%s) is not found in the service (%s)"
	errorServiceHasNoParentDevice   = "service (%s) has no parent device"
	errorServiceBadSCPDURL          = "SCPDURL (%s) is bad response (%d)"
)

// A Service represents a UPnP service.
type Service struct {
	XMLName     xml.Name `xml:"service"`
	ServiceType string   `xml:"serviceType"`
	ServiceID   string   `xml:"serviceId"`
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

// NewServiceFromDescription returns a service from the specified descrition string.
func NewServiceFromDescriptionBytes(descBytes []byte) (*Service, error) {
	service := NewService()

	err := service.LoadDescriptionBytes(descBytes)
	if err != nil {
		return nil, err
	}

	return service, err
}

// LoadDescriptionBytes loads a device description string.
func (service *Service) LoadDescriptionBytes(descBytes []byte) error {
	service.description = &ServiceDescription{}

	err := xml.Unmarshal(descBytes, service.description)
	if err != nil {
		return err
	}

	service.ServiceStateTable = &service.description.ServiceStateTable
	service.ActionList = &service.description.ActionList

	err = service.reviseParentObject()
	if err != nil {
		return err
	}

	return nil
}

// LoadDescriptinString loads a device description string.
func (service *Service) LoadDescriptionFromSCPDURL() error {
	// Some services has no SCPDURL such as Panasonic AiSEG001
	if len(service.SCPDURL) == 0 {
		return nil
	}

	scpdURL, err := service.GetAbsoluteSCPDURL()
	if err != nil {
		return err
	}

	res, err := http.Get(scpdURL.String())
	if err != nil {
		return fmt.Errorf("%w (%s)", err, scpdURL)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(errorServiceBadSCPDURL, scpdURL.String(), res.StatusCode)
	}

	scpdBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = service.LoadDescriptionBytes(scpdBytes)
	if err != nil {
		return fmt.Errorf("%w (%s)", err, scpdURL)
	}

	return res.Body.Close()
}

// DescriptionString returns a descrition string.
func (service *Service) DescriptionString() (string, error) {
	descBytes, err := xml.MarshalIndent(service.description, "", xmlMarshallIndent)
	if err != nil {
		return "", err
	}

	return string(descBytes), nil
}

func (service *Service) getShortServiceType() string {
	serviceTypes := strings.Split(service.ServiceType, ":")
	if len(serviceTypes) <= 1 {
		return service.ServiceID
	}
	return serviceTypes[len(serviceTypes)-2]
}

func (service *Service) reviseParentObject() error {
	if service.ActionList != nil {
		for n := 0; n < len(service.ActionList.Actions); n++ {
			action := &service.ActionList.Actions[n]
			action.ParentService = service
			action.reviseParentObject()
		}
	}

	if service.ServiceStateTable != nil {
		for n := 0; n < len(service.ServiceStateTable.StateVariables); n++ {
			statVar := &service.ServiceStateTable.StateVariables[n]
			statVar.ParentService = service
		}
	}

	return nil
}

func (service *Service) reviseDescription() error {
	shortServiceID := service.getShortServiceType()

	// check description URLs

	if len(service.SCPDURL) == 0 {
		service.SCPDURL = fmt.Sprintf(defaultServiceScpdURL, shortServiceID)
	}

	if len(service.ControlURL) == 0 {
		service.ControlURL = fmt.Sprintf(defaultServiceControlURL, shortServiceID)
	}

	if len(service.EventSubURL) == 0 {
		service.EventSubURL = fmt.Sprintf(defaultServiceEventURL, shortServiceID)
	}

	return nil
}

func (service *Service) isDescriptionURL(path string) bool {
	return path == service.SCPDURL
}

func (service *Service) isControlURL(path string) bool {
	return path == service.ControlURL
}

func (service *Service) isEventSubURL(path string) bool {
	return path == service.EventSubURL
}

func (service *Service) getAbsoluteURL(path string) (*url.URL, error) {
	if service.ParentDevice == nil {
		return nil, fmt.Errorf(errorServiceHasNoParentDevice, service.ServiceType)
	}
	return service.ParentDevice.GetAbsoluteURL(path)
}

func (service *Service) GetAbsoluteSCPDURL() (*url.URL, error) {
	return service.getAbsoluteURL(service.SCPDURL)
}

func (service *Service) GetAbsoluteControlURL() (*url.URL, error) {
	return service.getAbsoluteURL(service.ControlURL)
}

func (service *Service) GetAbsoluteEventSubURL() (*url.URL, error) {
	return service.getAbsoluteURL(service.EventSubURL)
}

// GetActions returns all actions.
func (service *Service) GetActions() []*Action {
	actionCnt := len(service.ActionList.Actions)
	actions := make([]*Action, actionCnt)
	for n := 0; n < actionCnt; n++ {
		actions[n] = &service.ActionList.Actions[n]
	}
	return actions
}

// GetActionByName returns an action by the specified name.
func (service *Service) GetActionByName(name string) (*Action, error) {
	if service.description == nil {
		return nil, fmt.Errorf(errorServiceDescriptionNotFound, service.XMLName.Local)
	}

	if len(service.ActionList.Actions) == 0 {
		return nil, fmt.Errorf(errorServiceHasNoActions, service.XMLName.Local)
	}

	for n := 0; n < len(service.ActionList.Actions); n++ {
		action := &service.ActionList.Actions[n]
		if action.Name == name {
			return action, nil
		}
	}

	return nil, fmt.Errorf(errorServiceActionNotFound, name, service.ServiceType)
}
