// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	gohttp "net/http"

	"github.com/cybergarage/go-net-upnp/net/upnp/util"
)

// A Client represents a Client.
type Client struct {
	*gohttp.Client
}

// NewRequest returns a new Request.
func NewClient() (*Client, error) {
	client := &Client{}
	client.Client = &gohttp.Client{}
	return client, nil
}

func (client *Client) Do(req *Request) (*Response, error) {
	if ua := req.Header.Get(UserAgent); ua == "" {
		req.Header.Set(UserAgent, util.GetUserAgent())
	}
	res, err := client.Client.Do(req.Request)
	if err != nil {
		return nil, err
	}
	return NewResponseFromResponse(res), nil
}
