// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"net/upnp/control"
)

// NewActionRequestFromAction returns a new Request.
func NewActionRequestFromAction(action *Action) (*control.ActionRequest, error) {
	req := control.NewActionRequest()

	req.Envelope.Body.Action.Name = action.Name

	service := action.ParentService
	if service != nil {
		req.Envelope.Body.Action.ServiceType = service.ServiceType
	}

	for n := 0; n < len(action.ArgumentList.Arguments); n++ {
		arg := &action.ArgumentList.Arguments[n]
		if arg.GetDirection() != InDirection {
			continue
		}
		reqArg := NewArgumentFromArgument(arg)
		req.Envelope.Body.Action.Arguments = append(req.Envelope.Body.Action.Arguments, reqArg)
	}

	return req, nil
}

// NewActionResponseFromAction returns a new Response.
func NewActionResponseFromAction(action *Action) (*control.ActionResponse, error) {
	res := control.NewActionResponse()

	// Fix 'Action' -> 'ActionResponse'
	res.Envelope.Body.Action.Name = action.Name + control.Response

	service := action.ParentService
	if service != nil {
		res.Envelope.Body.Action.ServiceType = service.ServiceType
	}

	for n := 0; n < len(action.ArgumentList.Arguments); n++ {
		arg := &action.ArgumentList.Arguments[n]
		if arg.GetDirection() != OutDirection {
			continue
		}
		resArg := NewArgumentFromArgument(arg)
		res.Envelope.Body.Action.Arguments = append(res.Envelope.Body.Action.Arguments, resArg)
	}

	return res, nil
}

// NewArgument returns a new argument.
func NewArgumentFromArgument(arg *Argument) *control.Argument {
	newArg := control.NewArgument()

	newArg.Name = arg.Name
	newArg.Value = arg.Value

	return newArg
}
