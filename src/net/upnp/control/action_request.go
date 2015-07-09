// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
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

// A Action represents arguments in as SOAP action.
type ActionInnerXML struct {
	Innerxml string `xml:",innerxml"`
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
	err = req.decodeBodyInnerXMLString(innerXMLString)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// decodeBodyInnerXMLString parses an innerXML of an action in body
func (self *ActionRequest) decodeBodyInnerXMLString(bodyInnerXML string) error {
	reader := strings.NewReader(bodyInnerXML)
	decorder := xml.NewDecoder(reader)

	for {
		token, err := decorder.Token()
		if token == nil {
			break
		}
		if err != nil {
			return err
		}
		switch elem := token.(type) {
		case xml.StartElement:
			actionName := elem.Name.Local
			self.Envelope.Body.Action.Name = actionName
			var actionArgs ActionInnerXML
			if err := decorder.DecodeElement(&actionArgs, &elem); err != nil {
				return err
			}
			err := self.decodeActionInnerXMLString(actionArgs.Innerxml)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// decodeActionInnerXMLString parses an innerXML of arguments in action
func (self *ActionRequest) decodeActionInnerXMLString(actionInnerXML string) error {
	reader := strings.NewReader(actionInnerXML)
	decorder := xml.NewDecoder(reader)

	for {
		token, err := decorder.Token()
		if token == nil {
			break
		}
		if err != nil {
			return err
		}
		switch elem := token.(type) {
		case xml.StartElement:
			argName := elem.Name.Local
			var argValue string
			if err := decorder.DecodeElement(&argValue, &elem); err != nil {
				return err
			}
			arg := &Argument{Name: argName, Value: argValue}
			self.Envelope.Body.Action.Arguments = append(self.Envelope.Body.Action.Arguments, arg)
		}
	}

	return nil
}

// GetAction returns an actions in the SOPA request.
func (self *ActionRequest) GetAction() (*Action, error) {
	return &self.Envelope.Body.Action, nil
}
