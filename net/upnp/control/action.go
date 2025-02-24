// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
)

const (
	soapActionSpace     = "u"
	soapActionPrefix    = soapActionSpace + xmlNsDelim
	soapActionSpaceAttr = xmlNs + xmlNsDelim + soapActionSpace
)

// A Action represents a SOAP action.
type Action struct {
	Name        string
	ServiceType string
	Arguments   []*Argument
}

func (action *Action) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = soapActionPrefix + action.Name
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: soapActionSpaceAttr}, Value: action.ServiceType},
	}

	e.EncodeToken(start)
	for n := 0; n < len(action.Arguments); n++ {
		arg := action.Arguments[n]
		argElem := xml.StartElement{Name: xml.Name{Local: arg.Name}}
		e.EncodeElement(arg.Value, argElem)
	}
	e.EncodeToken(start.End())

	return nil
}
