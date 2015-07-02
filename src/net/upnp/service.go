// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A Service represents a icon.
type Service struct {
	XMLName           xml.Name        `xml:"service"`
	ServiceType       string          `xml:"serviceType"`
	SCPDURL           string          `xml:"SCPDURL"`
	ControlURL        string          `xml:"controlURL"`
	EventSubURL       string          `xml:"eventSubURL"`
	ServiceStateTable []StateVariable `xml:"serviceStateTable"`
}

// A Service represents a icon.
type ServiceList struct {
	XMLName  xml.Name  `xml:"iconList"`
	Services []Service `xml:"icon"`
}

// NewService returns a new Service.
func NewService() *Service {
	service := &Service{}
	return service
}
