// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import ()

// A Response represents a Response.
type ActionResponse struct {
}

// NewResponse returns a new Response.
func NewActionResponse() *ActionResponse {
	res := &ActionResponse{}
	return res
}
