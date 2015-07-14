// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
)

// A ErrorResponse represents an error response.
type ErrorResponse struct {
	Envelope struct {
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
}

type UPnPError struct {
	XMLName          xml.Name `xml:"urn:schemas-upnp-org:control-1-0 UPnPError"`
	ErrorCode        int      `xml:"errorCode"`
	ErrorDescription string   `xml:"errorDescription"`
}

// NewErrorResponse returns a new error response.
func NewErrorResponse() *ErrorResponse {
	res := &ErrorResponse{}
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

// GetUPnPError returns an UPnP error.
func (self *ErrorResponse) GetUPnPError() (*UPnPError, error) {
	return &self.Envelope.Body.Fault.Detail.UPnPError, nil
}
