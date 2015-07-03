// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A Action represents a UPnP action.
type Action struct {
	XMLName      xml.Name     `xml:"action"`
	Name         string       `xml:"name"`
	ArgumentList ArgumentList `xml:"argumentList"`
}

// A ActionList represents a UPnP action list.
type ActionList struct {
	XMLName xml.Name `xml:"actionList"`
	Actions []Action `xml:"action"`
}

// NewAction returns a new Action.
func NewAction() *Action {
	icon := &Action{}
	return icon
}
