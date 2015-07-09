// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A Argument represents a UPnP argument.
type Argument struct {
	XMLName              xml.Name `xml:"argument"`
	Name                 string   `xml:"name"`
	Direction            string   `xml:"direction"`
	RelatedStateVariable string   `xml:"relatedStateVariable"`

	Value string `xml:"-"`
}

// A ArgumentList represents a UPnP argument list.
type ArgumentList struct {
	XMLName   xml.Name   `xml:"argumentList"`
	Arguments []Argument `xml:"argument"`
}

// NewArgument returns a new Argument.
func NewArgument() *Argument {
	arg := &Argument{}
	return arg
}
