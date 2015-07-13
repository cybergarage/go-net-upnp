// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

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
func NewActionRequestFromSOAPString(soapReq string) (*ActionRequest, error) {
	req := NewActionRequest()
	err := req.decodeEnvelopeXMLString(soapReq)
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
