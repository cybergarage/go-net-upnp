// Copyright 2015 Satoshi Konno. All rights reserved.
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

// WriteRequest sends the specified request.
func (self *UnicastSocket) WriteRequest(req *Request) (int, error) {
	return self.HTTPUSocket.Write(ADDRESS, PORT, req.Bytes())
}

// WriteBytes sends the specified bytes.
func (self *UnicastSocket) WriteBytes(addr string, port int, b []byte) (int, error) {
	return self.HTTPUSocket.Write(addr, port, b)
}

// WriteResponse sends the specified responst.
func (self *UnicastSocket) WriteResponse(addr string, port int, res *Response) (int, error) {
	return self.HTTPUSocket.Write(addr, port, res.Bytes())
}
