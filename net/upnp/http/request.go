// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"io"
	"net/url"
	"strings"

	gohttp "net/http"

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
	httpReq.Header.Add(SoapAction, soapAction)

	return httpReq, nil
}

func (self *Request) IsSOAPRequest() bool {
	_, ok := self.Header[SoapAction]
	return ok
}

func (self *Request) GetSOAPServiceActionName() (string, bool) {
	soapAction := self.Header.Get(SoapAction)
	if len(soapAction) <= 0 {
		return "", false
	}

	actions := strings.Split(soapAction, SoapActionDelim)
	if len(actions) < 2 {
		return "", false
	}

	return actions[len(actions)-1], true
}
