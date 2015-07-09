// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import ()

// A ActionRequest represents an action request.
type ActionRequest struct {
	*ActionControl
}

// NewRequest returns a new Request.
func NewActionRequest() *ActionRequest {
	req := &ActionRequest{}
	req.ActionControl = NewActionControl()
	return req
}

// NewRequest returns a new Request.
func NewActionRequestFromSOAPString(reqStr string) (*ActionRequest, error) {
	req := NewActionRequest()
	err := req.decodeEnvelopeXMLString(reqStr)
	if err != nil {
		return nil, err
	}

	innerXMLString := req.Envelope.Body.Innerxml
	err = req.decodeBodyInnerXMLString(innerXMLString)
	if err != nil {
		return nil, err
	}

	return req, nil
}
