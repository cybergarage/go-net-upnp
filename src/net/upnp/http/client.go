// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	gohttp "net/http"
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
	res, err := client.Client.Do(req.Request)
	if err != nil {
		return nil, err
	}
	return NewResponseFromResponse(res), nil
}
