// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A StateVariable represents a icon.
type StateVariable struct {
	XMLName           xml.Name          `xml:"stateVariable"`
	Name              string            `xml:"name"`
	DataType          string            `xml:"dataType"`
	DefaultValue      string            `xml:"defaultValue"`
	AllowedValueList  []AllowedValue    `xml:"allowedValueList"`
	AllowedValueRange AllowedValueRange `xml:"allowedValueRange"`
	SendEvents        string            `xml:"sendEvents,attr"`
	Multicast         string            `xml:"multicast,attr"`
}

// NewStateVariable returns a new StateVariable.
func NewStateVariable() *StateVariable {
	stat := &StateVariable{}
	return stat
}
