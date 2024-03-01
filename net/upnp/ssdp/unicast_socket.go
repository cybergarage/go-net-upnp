// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

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
func (socket *UnicastSocket) WriteRequest(req *Request) (int, error) {
	return socket.HTTPUSocket.Write(ADDRESS, Port, req.Bytes())
}

// WriteBytes sends the specified bytes.
func (socket *UnicastSocket) WriteBytes(addr string, port int, b []byte) (int, error) {
	return socket.HTTPUSocket.Write(addr, port, b)
}

// WriteResponse sends the specified responst.
func (socket *UnicastSocket) WriteResponse(addr string, port int, res *Response) (int, error) {
	return socket.HTTPUSocket.Write(addr, port, res.Bytes())
}
