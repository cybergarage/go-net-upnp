// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"io"
	gohttp "net/http"

	"net/upnp/util"
)

// A Request represents a Request.
type Request struct {
	*gohttp.Request
}

// NewRequest returns a new Request.
func NewRequestFromRequest(req *gohttp.Request) *Request {
	httpReq := &Request{Request: req}
	httpReq.Header.Add(UserAgent, util.GetUserAgent())
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

// NewSOAPRequest returns a new Request.
func NewSOAPRequest(urlStr string, soapAction string, body io.Reader) (*Request, error) {
	httpReq, err := NewRequest(POST, urlStr, body)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add(ContentType, ContentTypeXML)
	httpReq.Header.Add(SoapAction, soapAction)

	return httpReq, nil
}
