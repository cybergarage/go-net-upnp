// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"fmt"
	"net"
	gohttp "net/http"
	"net/upnp/util"
	"time"
)

const (
	TIMEOUT          = 60
	MAX_HEADER_BYTES = 1 << 20
)

func GetServerName() string {
	return util.GetUserAgent()
}

// A SSDPListener represents a listener for Server.
type RequestListener interface {
	HTTPRequestReceived(*Request, ResponseWriter)
}

// A Server represents a Server.
type Server struct {
	*gohttp.Server
	Conn     net.Listener
	Listener RequestListener
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

	var err error
	self.Conn, err = net.Listen("tcp", self.Addr)
	if err != nil {
		return err
	}

	go self.Server.Serve(self.Conn.(*net.TCPListener))

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

// ServeHTTP is a handler
func (self *Server) ServeHTTP(res gohttp.ResponseWriter, req *gohttp.Request) {
	if self.Listener == nil {
		res.WriteHeader(gohttp.StatusInternalServerError)
		return
	}

	self.Listener.HTTPRequestReceived(NewRequestFromRequest(req), res)
}
