// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
	"strings"

	"net/upnp"
)

const (
	errorActionControlInvalidInnerXML   = "invalid inner XML (%s)"
	errorActionControlInvalidActionName = "invalid action name (%s)"
)

// A ActionControl represents an action Control.
type ActionControl struct {
	Envelope struct {
		XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Body    struct {
			XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
			Innerxml string   `xml:",innerxml"`
			Action   Action
		}
	}
}

// A Action represents arguments in as SOAP action.
type ActionInnerXML struct {
	Innerxml string `xml:",innerxml"`
}

// NewActionControl returns a new Control.
func NewActionControl() *ActionControl {
	ctrl := &ActionControl{}
	ctrl.Envelope.Body.Action.Arguments = make([]*Argument, 0)
	return ctrl
}

// decodeEnvelopeXMLString parses an evnelope XML
func (self *ActionControl) decodeEnvelopeXMLString(envXML string) error {
	return xml.Unmarshal([]byte(envXML), &self.Envelope)
}

// decodeBodyInnerXMLString parses an innerXML of an action in body
func (self *ActionControl) decodeBodyInnerXMLString(bodyInnerXML string) error {
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
func (self *ActionControl) decodeActionInnerXMLString(actionInnerXML string) error {
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

// GetAction returns an actions in the SOAP Control.
func (self *ActionControl) GetAction() (*Action, error) {
	return &self.Envelope.Body.Action, nil
}

// SOAPContentBytes returns a SOAP content bytes.
func (self *ActionControl) SOAPContentBytes() ([]byte, error) {
	return xml.MarshalIndent(&self.Envelope, "", upnp.XmlMarshallIndent)
}

// SOAPContent returns a SOAP content string.
func (self *ActionControl) SOAPContentString() (string, error) {
	buf, err := self.SOAPContentBytes()
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
