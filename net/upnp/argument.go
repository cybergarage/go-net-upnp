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

// SetString sets a string value into the specified argument.
func (arg *Argument) SetString(value string) error {
	arg.Value = value
	return nil
}

// GetString returns a string value into the specified argument.
func (arg *Argument) GetString() (string, error) {
	return arg.Value, nil
}

// SetInt sets a integer value into the specified argument.
func (arg *Argument) SetInt(value int) error {
	return arg.SetString(strconv.Itoa(value))
}

// GetInt return a integer value into the specified argument.
func (arg *Argument) GetInt() (int, error) {
	value, err := arg.GetString()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(value)
}

// SetFloat sets a integer value into the specified argument.
func (arg *Argument) SetFloat(value float64) error {
	return arg.SetString(fmt.Sprintf("%f", value))
}

// GetFloat return a integer value into the specified argument.
func (arg *Argument) GetFloat() (float64, error) {
	value, err := arg.GetString()
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// SetBool sets a boolean value into the specified argument.
func (arg *Argument) SetBool(value bool) error {
	ivalue := 0
	if value {
		ivalue = 1
	}
	return arg.SetInt(ivalue)
}

// GetBool return a boolean value into the specified argument.
func (arg *Argument) GetBool() (bool, error) {
	value, err := arg.GetString()
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
func (arg *Argument) isDirection(value string) bool {
	return (arg.Direction == value)
}

// IsInDirection returns true when the argument direction is in, otherwise false.
func (arg *Argument) IsInDirection() bool {
	return arg.isDirection(In)
}

// IsOutDirection returns true when the argument direction is out, otherwise false.
func (arg *Argument) IsOutDirection() bool {
	return arg.isDirection(Out)
}

// SetDirection sets a directional string of the specified diractional integer.
func (arg *Argument) SetDirection(dir int) bool {
	switch dir {
	case InDirection:
		arg.Direction = In
		return true
	case OutDirection:
		arg.Direction = Out
		return true
	}
	return false
}

// GetDirection returns a directional integer of the specified argument.
func (arg *Argument) GetDirection() int {
	if arg.IsInDirection() {
		return InDirection
	}
	if arg.IsOutDirection() {
		return OutDirection
	}
	return UnknownDirection
}
