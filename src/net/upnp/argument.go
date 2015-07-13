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

	Value        string  `xml:"-"`
	ParentAction *Action `xml:"-"`
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

// SetString sets a value into the specified argument
func (self *Argument) SetString(value string) error {
	self.Value = value
	return nil
}

// GetString returns a value into the specified argument
func (self *Argument) GetString() (string, error) {
	return self.Value, nil
}

// isDirection returns true when the argument direction equals the specified value, otherwise false.
func (self *Argument) isDirection(value string) bool {
	return (self.Direction == value)
}

// IsInDirection returns true when the argument direction is in, otherwise false.
func (self *Argument) IsInDirection() bool {
	return self.isDirection(In)
}

// IsOutDirection returns true when the argument direction is out, otherwise false.
func (self *Argument) IsOutDirection() bool {
	return self.isDirection(Out)
}

// GetDirection returns a directional integer of the specified argument.
func (self *Argument) GetDirection() int {
	if self.IsInDirection() {
		return InDirection
	}
	if self.IsOutDirection() {
		return OutDirection
	}
	return UnknownDirection
}
