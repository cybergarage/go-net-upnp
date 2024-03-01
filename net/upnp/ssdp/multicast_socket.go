// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

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
func (socket *MulticastSocket) Write(req *Request) (int, error) {
	return socket.HTTPMUSocket.Write(req.Bytes())
}
