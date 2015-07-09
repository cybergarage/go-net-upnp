// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package soap

import (
	"encoding/xml"
)

// A Request represents a Request.
type Request struct {
	XMLName xml.Name `xml:"s:Envelope"`
	Body    interface{}
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
func NewRequest() *Request {
	soapReq := &Request{}
	return soapReq
}

// GetAction returns an actions in the SOPA request.
func (self *Request) GetAction() (*Action, error) {
	//decoder := xml.NewDecoder(body)
	return nil, nil
}

func (req *Request) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := dec.Token()
		if token == nil {
			break
		}
		if err != nil {
			return err
		}
		if t, ok := token.(xml.StartElement); ok {
			var data string
			if err := dec.DecodeElement(&data, &t); err != nil {
				return err
			}
		}
	}

	return nil
}
