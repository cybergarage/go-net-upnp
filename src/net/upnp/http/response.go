// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	gohttp "net/http"
)

// A Response represents a Response.
type Response struct {
	*gohttp.Response
}

// NewResponse returns a new Response.
func NewResponseFromResponse(req *gohttp.Response) *Response {
	httpReq := &Response{Response: req}
	return httpReq
}

// NewResponse returns a new Response.
func NewResponse(req *gohttp.Response) *Response {
	return NewResponseFromResponse(&gohttp.Response{})
}
