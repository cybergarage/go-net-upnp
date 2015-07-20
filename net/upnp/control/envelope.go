// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
)

const (
	soapEnvelopeNamespace = "http://schemas.xmlsoap.org/soap/envelope/"
	soapEnvelopeEncoding  = "http://schemas.xmlsoap.org/soap/encoding/"

	soapEnvelope           = "Envelope"
	soapEncodingStyle      = "encodingStyle"
	soapEnvelopeSpace      = "s"
	soapEnvelopePrefix     = soapEnvelopeSpace + xmlNsDelim
	soapEnvelopeSpaceAttr  = xmlNs + xmlNsDelim + soapEnvelopeSpace
	soapEnvelopeEncodeAttr = soapEnvelopeSpace + xmlNsDelim + soapEncodingStyle

	soapBody       = "Body"
	soapBodySpace  = "s"
	soapSoapPrefix = soapBodySpace + xmlNsDelim
)

// A Envelope represents an Envelope.
type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName  xml.Name `xml:"Body"`
		Innerxml string   `xml:",innerxml"`
		Action   Action
	}
}

func (self *Envelope) MarshalXML(e *xml.Encoder, env xml.StartElement) error {
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
