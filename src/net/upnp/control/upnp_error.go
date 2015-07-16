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
	XMLName          xml.Name `xml:"urn:schemas-upnp-org:control-1-0 UPnPError"`
	ErrorCode        int      `xml:"errorCode"`
	ErrorDescription string   `xml:"errorDescription"`
}

// NewUPnPError returns a new null UPnPError.
func NewUPnPError() *UPnPError {
	err := &UPnPError{}
	return err
}

// NewUPnPErrorFromCode returns a new UPnPError from the specified code.
func NewUPnPErrorFromCode(code int) *UPnPError {
	err := &UPnPError{
		ErrorCode:        code,
		ErrorDescription: errorCodeToString(code),
	}
	return err
}

func (self UPnPError) Error() string {
	return fmt.Sprintf(upnpErrorFormat, self.ErrorCode, self.ErrorDescription)
}
