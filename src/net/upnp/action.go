// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"errors"
	"fmt"
)

const (
	errorActionArgumentNotFound = "argument (%s) is not found"
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
	action := &Action{}
	return action
}

// GetArgumentByName returns an argument by the specified name
func (self *Action) GetArgumentByName(name string) (*Argument, error) {
	for n := 0; n < len(self.ArgumentList.Arguments); n++ {
		arg := &self.ArgumentList.Arguments[n]
		if arg.Name == name {
			return arg, nil
		}
	}
	return nil, errors.New(fmt.Sprintf(errorActionArgumentNotFound, name))
}
