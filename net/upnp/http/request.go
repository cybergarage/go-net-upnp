// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"io"
	gohttp "net/http"
	"net/url"
	"strings"

	"github.com/cybergarage/go-net-upnp/net/upnp/util"
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
func NewSOAPRequest(url *url.URL, soapAction string, body io.Reader) (*Request, error) {
	httpReq, err := NewRequest(POST, url.String(), body)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add(ContentType, ContentTypeXML)
	httpReq.Header.Add(SOAPAction, soapAction)

	return httpReq, nil
}

func (req *Request) IsSOAPRequest() bool {
	_, ok := req.Header[SOAPAction]
	return ok
}

func (req *Request) GetSOAPServiceActionName() (string, bool) {
	soapAction := req.Header.Get(SOAPAction)
	if len(soapAction) == 0 {
		return "", false
	}

	actions := strings.Split(soapAction, SOAPActionDelim)
	if len(actions) < 2 {
		return "", false
	}

	return actions[len(actions)-1], true
}
