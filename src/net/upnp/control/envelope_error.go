// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
)

// A ErrorEnvelope represents an Envelope for error response.
type ErrorEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
		Fault   struct {
			XMLName     xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
			Faultcode   string   `xml:"faultcode"`
			Faultstring string   `xml:"faultstring"`
			Detail      struct {
				XMLName   xml.Name  `xml:"detail"`
				UPnPError UPnPError `xml:"urn:schemas-upnp-org:control-1-0 UPnPError"`
			}
		}
	}
}

/*
func (self *ErrorEnvelope) MarshalXML(e *xml.Encoder, env xml.StartElement) error {
	// <Envelope>

	env.Name.Local = soapEnvelopePrefix + soapEnvelope
	env.Attr = []xml.Attr{
		{Name: xml.Name{Local: soapEnvelopeSpaceAttr}, Value: soapEnvelopeNamespace},
		{Name: xml.Name{Local: soapEnvelopeEncodeAttr}, Value: soapEnvelopeEncoding},
	}
	e.EncodeToken(env)

	// <Body>

	body := xml.StartElement{Name: xml.Name{Local: (soapSoapPrefix + soapBody)}}
	e.EncodeToken(body)

	// <Action>

	action := &self.Body.Action
	action.MarshalXML(e, xml.StartElement{})

	// </Body>

	e.EncodeToken(body.End())

	// </Envelope>

	e.EncodeToken(env.End())

	return nil
}
*/
