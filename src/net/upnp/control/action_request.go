// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

const (
	errorActionRequestInvalidInnerXML   = "invalid inner XML (%s)"
	errorActionRequestInvalidActionName = "invalid action name (%s)"
)

// A ActionRequest represents an action request.
type ActionRequest struct {
	Envelope struct {
		XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Body    struct {
			XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
			Innerxml string   `xml:",innerxml"`
			Action   Action   `xml:"-"`
		}
	}
}

// A Action represents a SOAP action.
type Action struct {
	Name      string
	Arguments []*Argument
}

// A Action represents arguments in as SOAP action.
type Argument struct {
	Name  string
	Value string
}

// NewRequest returns a new Request.
func NewActionRequest() *ActionRequest {
	req := &ActionRequest{}
	req.Envelope.Body.Action.Arguments = make([]*Argument, 0)
	return req
}

// NewRequest returns a new Request.
func NewActionRequestFromSOAPString(reqStr string) (*ActionRequest, error) {
	req := NewActionRequest()
	err := xml.Unmarshal([]byte(reqStr), &req.Envelope)
	if err != nil {
		return nil, err
	}

	innerXMLString := req.Envelope.Body.Innerxml
	err = req.decodeInnerXML(innerXMLString)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// decodeInnerXML parses the innerXML
func (self *ActionRequest) decodeInnerXML(innerXML string) error {
	xmlSepFunc := func(r rune) bool {
		switch r {
		case '<', '>', ' ':
			return true
		}
		return false
	}

	params := strings.FieldsFunc(innerXML, xmlSepFunc)

	if len(params) < 2 {
		return errors.New(fmt.Sprintf(errorActionRequestInvalidInnerXML, innerXML))
	}

	names := strings.Split(params[0], ":")
	if len(names) < 2 {
		return errors.New(fmt.Sprintf(errorActionRequestInvalidActionName, params[0]))
	}
	self.Envelope.Body.Action.Name = names[1]

	for n := 2; (n + 2) < len(params); n += 3 {
		switch n {
		case 0:
		default:
			arg := &Argument{Name: params[n], Value: params[n+1]}
			self.Envelope.Body.Action.Arguments = append(self.Envelope.Body.Action.Arguments, arg)
		}
	}

	return nil
}

// GetAction returns an actions in the SOPA request.
func (self *ActionRequest) GetAction() (*Action, error) {
	return &self.Envelope.Body.Action, nil
}
