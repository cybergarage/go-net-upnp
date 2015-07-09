// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"strings"
)

const (
	Response = "Response"
)

// A ActionResponse represents an action request.
type ActionResponse struct {
	*ActionControl
}

// NewRequest returns a new Request.
func NewActionResponse() *ActionResponse {
	req := &ActionResponse{}
	req.ActionControl = NewActionControl()
	return req
}

// NewRequest returns a new Request.
func NewActionResponseFromSOAPString(reqStr string) (*ActionResponse, error) {
	res := NewActionResponse()
	err := res.decodeEnvelopeXMLString(reqStr)
	if err != nil {
		return nil, err
	}

	innerXMLString := res.Envelope.Body.Innerxml
	err = res.decodeBodyInnerXMLString(innerXMLString)
	if err != nil {
		return nil, err
	}

	// Fix 'ActionResponse' -> 'Action'
	res.Envelope.Body.Action.Name = strings.TrimSuffix(res.Envelope.Body.Action.Name, Response)

	return res, nil
}
