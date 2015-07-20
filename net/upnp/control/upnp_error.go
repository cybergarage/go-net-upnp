// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
	"fmt"
)

const (
	upnpErrorFormat = "UPnP Error : [%d] %s"
)

const (
	ErrorInvalidAction                = 401
	ErrorInvalidArgs                  = 402
	ErrorActionFailed                 = 501
	ErrorArgumentValueInvalid         = 600
	ErrorArgumentValueOutOfRange      = 601
	ErrorOptionalActionNotImplemented = 602
	ErrorOutOfMemory                  = 603
	ErrorHumanInterventionRequired    = 604
	ErrorStringArgumentTooLong        = 605
)

const (
	soapUPnPError          = "UPnPError"
	soapUPnPErrorNamespace = "urn:schemas-upnp-org:control-1-0"

	soapUPnPErrorCode = "errorCode"
	soapUPnPErrorDesc = "errorDescription"
)

func errorCodeToString(code int) string {
	errMsgs := map[int]string{
		ErrorInvalidAction:                "Invalid Action",
		ErrorInvalidArgs:                  "Invalid Args",
		ErrorActionFailed:                 "Action Failed",
		ErrorArgumentValueInvalid:         "Argument Value Invalid",
		ErrorArgumentValueOutOfRange:      "Argument Value Out of Range",
		ErrorOptionalActionNotImplemented: "Optional Action Not Implemented",
		ErrorOutOfMemory:                  "Out of Memory",
		ErrorHumanInterventionRequired:    "Human Intervention Required",
		ErrorStringArgumentTooLong:        "String Argument Too Long",
	}

	msg, ok := errMsgs[code]
	if !ok {
		return ""
	}

	return msg
}

type UPnPError struct {
	XMLName     xml.Name `xml:"urn:schemas-upnp-org:control-1-0 UPnPError"`
	Code        int      `xml:"errorCode"`
	Description string   `xml:"errorDescription"`
}

// NewUPnPError returns a new null UPnPError.
func NewUPnPError() *UPnPError {
	err := &UPnPError{}
	return err
}

// NewUPnPErrorFromCode returns a new UPnPError from the specified code.
func NewUPnPErrorFromCode(code int) *UPnPError {
	err := &UPnPError{
		Code:        code,
		Description: errorCodeToString(code),
	}
	return err
}

func (self *UPnPError) GetCode() int {
	return self.Code
}

func (self *UPnPError) GetDescription() string {
	return self.Description
}

func (self *UPnPError) Error() string {
	return fmt.Sprintf(upnpErrorFormat, self.Code, self.Description)
}

func (self *UPnPError) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = soapUPnPError
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: xmlNs}, Value: soapUPnPErrorNamespace},
	}

	e.EncodeToken(start)

	// errorCode

	errCode := xml.StartElement{Name: xml.Name{Local: soapUPnPErrorCode}}
	e.EncodeElement(self.Code, errCode)

	// errorDescripton

	errDesc := xml.StartElement{Name: xml.Name{Local: soapUPnPErrorDesc}}
	e.EncodeElement(self.Description, errDesc)

	e.EncodeToken(start.End())

	return nil
}
