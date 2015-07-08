// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package soap

import (
	"encoding/xml"
)

// A Request represents a Request.
type Request struct {
	XMLName xml.Name `xml:"s:Envelope"`
	Body    interface{}
}

// NewRequest returns a new Request.
func NewRequest() *Request {
	soapReq := &Request{}
	return soapReq
}

// Read reads from the current opend socket.
func (self *Request) GetAction() string {
	//decoder := xml.NewDecoder(body)
	return ""
}
