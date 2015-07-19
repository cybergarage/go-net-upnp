// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"fmt"
	"strconv"
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

// SetString sets a string value into the specified argument
func (self *Argument) SetString(value string) error {
	self.Value = value
	return nil
}

// GetString returns a string value into the specified argument
func (self *Argument) GetString() (string, error) {
	return self.Value, nil
}

// SetInt sets a integer value into the specified argument
func (self *Argument) SetInt(value int) error {
	return self.SetString(strconv.Itoa(value))
}

// GetInt return a integer value into the specified argument
func (self *Argument) GetInt() (int, error) {
	value, err := self.GetString()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(value)
}

// SetFloat sets a integer value into the specified argument
func (self *Argument) SetFloat(value float64) error {
	return self.SetString(fmt.Sprint("%f", value))
}

// GetFloat return a integer value into the specified argument
func (self *Argument) GetFloat() (float64, error) {
	value, err := self.GetString()
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// SetBool sets a boolean value into the specified argument
func (self *Argument) SetBool(value bool) error {
	ivalue := 0
	if value {
		ivalue = 1
	}
	return self.SetInt(ivalue)
}

// GetBool return a boolean value into the specified argument
func (self *Argument) GetBool() (bool, error) {
	value, err := self.GetString()
	if err != nil {
		return false, err
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return b, nil
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

// SetDirection sets a directional string of the specified diractional integer.
func (self *Argument) SetDirection(dir int) bool {
	switch dir {
	case InDirection:
		self.Direction = In
		return true
	case OutDirection:
		self.Direction = Out
		return true
	}
	return false
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
