// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"strings"
)

const (
	ResponseSuffix = "Response"
)

// A ActionResponse represents an action request.
type ActionResponse struct {
	*ActionControl
}

// NewActionResponse returns a new response.
func NewActionResponse() *ActionResponse {
	res := &ActionResponse{}
	res.ActionControl = NewActionControl()
	return res
}

// NewActionResponseFromSOAPString returns a new response.
func NewActionResponseFromSOAPBytes(soapRes []byte) (*ActionResponse, error) {
	res := NewActionResponse()
	err := res.decodeEnvelopeXMLBytes(soapRes)
	if err != nil {
		return nil, err
	}

	innerXMLBytes := res.Envelope.Body.Innerxml
	err = res.decodeBodyInnerXMLBytes([]byte(innerXMLBytes))
	if err != nil {
		return nil, err
	}

	// Fix 'ActionResponse' -> 'Action'
	res.Envelope.Body.Action.Name = strings.TrimSuffix(res.Envelope.Body.Action.Name, ResponseSuffix)

	return res, nil
}
