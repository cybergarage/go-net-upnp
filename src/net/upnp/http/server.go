// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"fmt"
	"net"
	gohttp "net/http"
	"time"
)

const (
	TIMEOUT          = 60
	MAX_HEADER_BYTES = 1 << 20
)

// A SSDPListener represents a listener for Server.
type RequestListener interface {
	HTTPRequestReceived(httpReq gohttp.Request)
}

// A Server represents a Server.
type Server struct {
	*gohttp.Server
	Conn net.Listener
}

// NewServer returns a new Server.
func NewServer() *Server {
	Server := &Server{}
	return Server
}

// Start starts this server.
func (self *Server) Start(port int) error {
	self.Server = &gohttp.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        self,
		ReadTimeout:    TIMEOUT * time.Second,
		WriteTimeout:   TIMEOUT * time.Second,
		MaxHeaderBytes: MAX_HEADER_BYTES,
	}

	err := self.Server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// Stop stops this server.
func (self *Server) Stop() error {
	if self.Conn != nil {
		self.Conn.Close()
		self.Conn = nil
	}
	return nil
}

// ListenAndServe overides net/http to close the connection
func (self *Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", self.Addr)
	if err != nil {
		return err
	}
	self.Conn = ln
	return self.Server.Serve(ln.(*net.TCPListener))
}

// ServeHTTP is a handler
func (self *Server) ServeHTTP(gohttp.ResponseWriter, *gohttp.Request) {
}
