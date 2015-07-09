// Copyright 2015 Satoshi Konno. All rights reserved.
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

	xmlSepFunc := func(r rune) bool {
		switch r {
		case '<', '>', ' ':
			return true
		}
		return false
	}

	innerXML := req.Envelope.Body.Innerxml
	params := strings.FieldsFunc(innerXML, xmlSepFunc)

	for n, param := range params {
		fmt.Printf("params[%d] =  %s\n", n, param)
	}

	if len(params) < 2 {
		return nil, errors.New(fmt.Sprintf(errorActionRequestInvalidInnerXML, innerXML))
	}

	names := strings.Split(params[0], ":")
	if len(names) < 2 {
		return nil, errors.New(fmt.Sprintf(errorActionRequestInvalidActionName, params[0]))
	}
	req.Envelope.Body.Action.Name = names[1]

	for n := 2; (n + 2) < len(params); n += 3 {
		fmt.Printf("n = %d %d\n", n, len(params))
		switch n {
		case 0:
		default:
			arg := &Argument{Name: params[n], Value: params[n+1]}
			req.Envelope.Body.Action.Arguments = append(req.Envelope.Body.Action.Arguments, arg)
			fmt.Printf("[%d] %s %s (%p) %d\n", n, arg.Name, arg.Value, req.Envelope.Body.Action.Arguments, len(req.Envelope.Body.Action.Arguments))
		}
	}

	return req, nil
}

// GetAction returns an actions in the SOPA request.
func (self *ActionRequest) GetAction() (*Action, error) {
	return &self.Envelope.Body.Action, nil
}
