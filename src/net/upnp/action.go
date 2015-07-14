// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"

	"net/upnp/control"
	"net/upnp/http"
)

const (
	errorActionArgumentNotFound = "argument (%s) is not found"
)

// A Action represents a UPnP action.
type Action struct {
	XMLName       xml.Name     `xml:"action"`
	Name          string       `xml:"name"`
	ArgumentList  ArgumentList `xml:"argumentList"`
	ParentService *Service     `xml:"-"`
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

func (self *Action) reviseParentObject() error {
	for n := 0; n < len(self.ArgumentList.Arguments); n++ {
		arg := &self.ArgumentList.Arguments[n]
		arg.ParentAction = self
	}

	return nil
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

// SetArgumentString sets a value into the specified argument
func (self *Action) SetArgumentString(name string, value string) error {
	arg, err := self.GetArgumentByName(name)
	if err != nil {
		return err
	}
	return arg.SetString(value)
}

// GetArgumentString return a value into the specified argument
func (self *Action) GetArgumentString(name string) (string, error) {
	arg, err := self.GetArgumentByName(name)
	if err != nil {
		return "", err
	}
	return arg.GetString()
}

// Post sends the specified arguments into the deveice.
func (self *Action) Post() error {
	req, err := NewActionRequestFromAction(self)
	if err != nil {
		return err
	}

	soapReqBytes, err := req.SOAPContentBytes()
	if err != nil {
		return err
	}

	httpReq, err := http.NewSOAPRequest("", "", bytes.NewBuffer(soapReqBytes))
	if err != nil {
		return err
	}

	httpClient, err := http.NewClient()
	if err != nil {
		return err
	}

	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		return err
	}

	statusCode := httpRes.StatusCode
	defer httpRes.Body.Close()
	soapResBytes, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}

	if statusCode == http.StatusOK {
		_, err = control.NewActionResponseFromSOAPBytes(soapResBytes)
		if err != nil {
			return err
		}

	} else {
		_, err := control.NewErrorResponseFromSOAPBytes(soapResBytes)
		if err != nil {
			return err
		}
	}

	return nil
}
