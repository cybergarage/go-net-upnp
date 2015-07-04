// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import ()

// A UnicastSocket represents a packet of SSDP.
type UnicastSocket struct {
	*HTTPUSocket
}

// NewUnicastSocket returns a new UnicastSocket.
func NewUnicastSocket() *UnicastSocket {
	ssdpSock := &UnicastSocket{}
	ssdpSock.HTTPUSocket = NewHTTPUSocket()
	return ssdpSock
}

// Write sends the specified bytes.
func (self *UnicastSocket) Write(req *Request) (int, error) {
	return self.HTTPUSocket.Write(ADDRESS, PORT, req.Bytes())
}
