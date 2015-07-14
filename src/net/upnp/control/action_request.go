// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"net/upnp"
)

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

// NewActionRequestFromAction returns a new Request.
func NewActionRequestFromAction(action *upnp.Action) (*ActionRequest, error) {
	req := NewActionRequest()

	req.Envelope.Body.Action.Name = action.Name

	service := action.ParentService
	if service != nil {
		req.Envelope.Body.Action.ServiceType = service.ServiceType
	}

	for n := 0; n < len(action.ArgumentList.Arguments); n++ {
		arg := &action.ArgumentList.Arguments[n]
		if arg.GetDirection() != upnp.InDirection {
			continue
		}
		reqArg := NewArgumentFromArgument(arg)
		req.Envelope.Body.Action.Arguments = append(req.Envelope.Body.Action.Arguments, reqArg)
	}

	return req, nil
}
