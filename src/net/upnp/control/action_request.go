// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import ()

// A ActionRequest represents an action request.
type ActionRequest struct {
	*ActionControl
}

// NewActionRequest returns a new Request.
func NewActionRequest() *ActionRequest {
	req := &ActionRequest{}
	req.ActionControl = NewActionControl()
	return req
}

// NewActionRequestFromSOAPString returns a new Request.
func NewActionRequestFromSOAPBytes(soapReq []byte) (*ActionRequest, error) {
	req := NewActionRequest()
	err := req.decodeEnvelopeXMLBytes(soapReq)
	if err != nil {
		return nil, err
	}

	InnerXMLBytes := req.Envelope.Body.Innerxml
	err = req.decodeBodyInnerXMLBytes([]byte(InnerXMLBytes))
	if err != nil {
		return nil, err
	}

	return req, nil
}
