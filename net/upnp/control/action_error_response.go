// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
)

// A ErrorResponse represents an error response.
type ErrorResponse struct {
	Envelope ErrorEnvelope `xml:"Envelope"`
}

// NewErrorResponse returns a new error response.
func NewErrorResponse() *ErrorResponse {
	res := &ErrorResponse{}

	res.Envelope.Body.Fault.FaultCode = soapFaultCodeDefault
	res.Envelope.Body.Fault.FaultString = soapFaultStringDefault

	return res
}

// NewErrorResponseFromSOAPString returns a  new error response.
func NewErrorResponseFromSOAPBytes(soapStr []byte) (*ErrorResponse, error) {
	res := NewErrorResponse()

	err := xml.Unmarshal(soapStr, &res.Envelope)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// NewErrorResponseFromUPnPError returns a new error response.
func NewErrorResponseFromUPnPError(upnpError *UPnPError) *ErrorResponse {
	res := NewErrorResponse()
	res.Envelope.Body.Fault.Detail.UPnPError = *upnpError
	return res
}

// GetUPnPError returns an UPnP error.
func (self *ErrorResponse) GetUPnPError() (*UPnPError, error) {
	return &self.Envelope.Body.Fault.Detail.UPnPError, nil
}

// SOAPContent returns a SOAP content string.
func (self *ErrorResponse) SOAPContentString() (string, error) {
	buf, err := xml.MarshalIndent(&self.Envelope, "", xmlMarshallIndent)
	if err != nil {
		return "", err
	}
	return xmlHeader + string(buf), nil
}
