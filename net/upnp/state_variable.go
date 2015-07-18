// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A StateVariable represents a UPnP state variable.
type StateVariable struct {
	XMLName           xml.Name          `xml:"stateVariable"`
	Name              string            `xml:"name"`
	DataType          string            `xml:"dataType"`
	DefaultValue      string            `xml:"defaultValue"`
	AllowedValueList  AllowedValueList  `xml:"allowedValueList"`
	AllowedValueRange AllowedValueRange `xml:"allowedValueRange"`
	SendEvents        string            `xml:"sendEvents,attr"`
	Multicast         string            `xml:"multicast,attr"`
	ParentService     *Service          `xml:"-"`
}

// A StateVariable represents a UPnP state variable.
type ServiceStateTable struct {
	XMLName        xml.Name        `xml:"serviceStateTable"`
	StateVariables []StateVariable `xml:"stateVariable"`
}

// NewStateVariable returns a new StateVariable.
func NewStateVariable() *StateVariable {
	stat := &StateVariable{}
	return stat
}
