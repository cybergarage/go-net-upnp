// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"fmt"
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
	return nil
}

// ServeHTTP is a handler
func (self *Server) ServeHTTP(gohttp.ResponseWriter, *gohttp.Request) {
}
