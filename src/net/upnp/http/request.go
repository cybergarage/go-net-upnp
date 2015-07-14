// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"io"
	gohttp "net/http"
)

// A Request represents a Request.
type Request struct {
	*gohttp.Request
}

// NewRequest returns a new Request.
func NewRequestFromRequest(req *gohttp.Request) *Request {
	httpReq := &Request{Request: req}
	return httpReq
}

// NewRequest returns a new Request.
func NewRequest(method, urlStr string, body io.Reader) (*Request, error) {
	req, err := gohttp.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}
	httpReq := NewRequestFromRequest(req)
	return httpReq, nil
}
