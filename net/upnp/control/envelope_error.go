// Copyright 2015 Satoshi Konno. All rights reserved.
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
			FaultCode   string   `xml:"faultcode"`
			FaultString string   `xml:"faultstring"`
			Detail      struct {
				XMLName   xml.Name  `xml:"detail"`
				UPnPError UPnPError `xml:"urn:schemas-upnp-org:control-1-0 UPnPError"`
			}
		}
	}
}

const (
	soapFault       = "Fault"
	soapFaultSpace  = "s"
	soapFaultPrefix = soapFaultSpace + xmlNsDelim

	soapFaultCode        = "faultcode"
	soapFaultCodeDefault = "s:Client"

	soapFaultString        = "faultstring"
	soapFaultStringDefault = "UPnPError"

	soapDetail = "detail"
)

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

	// <Fault>

	fault := xml.StartElement{Name: xml.Name{Local: (soapFaultPrefix + soapFault)}}
	e.EncodeToken(fault)

	faultCode := xml.StartElement{Name: xml.Name{Local: soapFaultCode}}
	e.EncodeElement(self.Body.Fault.FaultCode, faultCode)

	faultStr := xml.StartElement{Name: xml.Name{Local: soapFaultString}}
	e.EncodeElement(self.Body.Fault.FaultString, faultStr)

	// <detail>

	detail := xml.StartElement{Name: xml.Name{Local: soapDetail}}
	e.EncodeToken(detail)

	// <UPnPError>

	upnpErr := &self.Body.Fault.Detail.UPnPError
	upnpErr.MarshalXML(e, xml.StartElement{})

	// </detail>

	e.EncodeToken(detail.End())

	// </Fault>

	e.EncodeToken(fault.End())

	// </Body>

	e.EncodeToken(body.End())

	// </Envelope>

	e.EncodeToken(env.End())

	return nil
}
