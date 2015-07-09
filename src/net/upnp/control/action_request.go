// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

// A ActionRequest represents an action request.
type ActionRequest struct {
	Envelope struct {
		Body struct {
			Action Action
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
	return req
}

// NewRequest returns a new Request.
func NewActionRequestFromSOAPString(reqStr string) (*ActionRequest, error) {
	req := NewActionRequest()
	err := req.decodeXMLBytes([]byte(reqStr))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetAction returns an actions in the SOPA request.
func (self *ActionRequest) GetAction() (*Action, error) {
	//decoder := xml.NewDecoder(body)
	return nil, nil
}

func (req *ActionRequest) decodeXMLBytes(xmlBytes []byte) error {
	decoder := xml.NewDecoder(bytes.NewReader(xmlBytes))

	var value string
	token, err := decoder.Token()
	for n := 0; err != io.EOF; n++ {
		fmt.Printf("token[%d] = %p\n", n, token)

		if token == nil {
			break
		}
		/*
			if err != nil {
				return err
			}
		*/

		//var value string

		//token.(type).String()
		switch elem := token.(type) {
		case xml.StartElement:
			decoder.DecodeElement(&value, &elem)
			fmt.Printf("startElem (%s) = %s\n", elem.Name, value)
		case xml.EndElement:
			fmt.Printf("EndElement (%s) = %s\n", elem.Name, value)
		}

		token, _ = decoder.Token()
	}

	return nil
}
