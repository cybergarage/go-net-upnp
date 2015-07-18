// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import ()

// A MulticastSocket represents a socket of SSDP.
type MulticastSocket struct {
	*HTTPMUSocket
}

// NewMulticastSocket returns a new MulticastSocket.
func NewMulticastSocket() (*MulticastSocket, error) {
	ssdpSock := &MulticastSocket{}
	ssdpSock.HTTPMUSocket = NewHTTPMUSocket()
	return ssdpSock, nil
}

// Write sends the specified bytes.
func (self *MulticastSocket) Write(req *Request) (int, error) {
	return self.HTTPMUSocket.Write(req.Bytes())
}
