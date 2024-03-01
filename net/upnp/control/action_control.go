// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"bytes"
	"encoding/xml"
)

const (
	errorActionControlInvalidInnerXML   = "invalid inner XML (%s)"
	errorActionControlInvalidActionName = "invalid action name (%s)"
)

// A ActionControl represents an action Control.
type ActionControl struct {
	Envelope Envelope `xml:"Envelope"`
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
func (action *ActionControl) decodeEnvelopeXMLBytes(envXML []byte) error {
	return xml.Unmarshal([]byte(envXML), &action.Envelope)
}

// decodeBodyInnerXMLBytes parses an innerXML of an action in body
func (action *ActionControl) decodeBodyInnerXMLBytes(bodyInnerXML []byte) error {
	reader := bytes.NewReader(bodyInnerXML)
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
			action.Envelope.Body.Action.Name = actionName
			var actionArgs ActionInnerXML
			if err := decorder.DecodeElement(&actionArgs, &elem); err != nil {
				return err
			}
			err := action.decodeActionInnerXMLBytes([]byte(actionArgs.Innerxml))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// decodeActionInnerXMLBytes parses an innerXML of arguments in action
func (action *ActionControl) decodeActionInnerXMLBytes(actionInnerXML []byte) error {
	reader := bytes.NewReader(actionInnerXML)
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
			action.Envelope.Body.Action.Arguments = append(action.Envelope.Body.Action.Arguments, arg)
		}
	}

	return nil
}

// GetAction returns an actions in the SOAP Control.
func (action *ActionControl) GetAction() (*Action, error) {
	return &action.Envelope.Body.Action, nil
}

// SOAPContent returns a SOAP content string.
func (action *ActionControl) SOAPContentString() (string, error) {
	buf, err := xml.MarshalIndent(&action.Envelope, "", xmlMarshallIndent)
	if err != nil {
		return "", err
	}
	return xmlHeader + string(buf), nil
}
