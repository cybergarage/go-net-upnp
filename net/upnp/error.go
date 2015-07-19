// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"github.com/cybergarage/go-net-upnp/net/upnp/control"
)

// A Error represents a interface for UPnP error.
type Error interface {
	error
	GetCode() int
	GetDescription() string
}

const (
	ErrorInvalidAction                = control.ErrorInvalidAction
	ErrorInvalidArgs                  = control.ErrorInvalidArgs
	ErrorActionFailed                 = control.ErrorActionFailed
	ErrorArgumentValueInvalid         = control.ErrorArgumentValueInvalid
	ErrorArgumentValueOutOfRange      = control.ErrorArgumentValueOutOfRange
	ErrorOptionalActionNotImplemented = control.ErrorOptionalActionNotImplemented
	ErrorOutOfMemory                  = control.ErrorOutOfMemory
	ErrorHumanInterventionRequired    = control.ErrorHumanInterventionRequired
	ErrorStringArgumentTooLong        = control.ErrorStringArgumentTooLong
)

// NewErrorFromCode returns a new Error from the specified code.
func NewErrorFromCode(code int) Error {
	return control.NewUPnPErrorFromCode(code)
}
