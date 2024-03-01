// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"fmt"
	"net"
	gohttp "net/http"
	"time"

	"github.com/cybergarage/go-net-upnp/net/upnp/util"
)

const (
	DefaultTimeout = 60
	MaxHeaderBytes = 1 << 20
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
func (server *Server) Start(port int) error {
	server.Server = &gohttp.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        server,
		ReadTimeout:    DefaultTimeout * time.Second,
		WriteTimeout:   DefaultTimeout * time.Second,
		MaxHeaderBytes: MaxHeaderBytes,
	}

	var err error
	server.Conn, err = net.Listen("tcp", server.Addr)
	if err != nil {
		return err
	}

	go server.Server.Serve(server.Conn.(*net.TCPListener))

	return nil
}

// Stop stops this server.
func (server *Server) Stop() error {
	if server.Conn != nil {
		server.Conn.Close()
		server.Conn = nil
	}
	return nil
}

// ServeHTTP is a handler
func (server *Server) ServeHTTP(res gohttp.ResponseWriter, req *gohttp.Request) {
	if server.Listener == nil {
		res.WriteHeader(gohttp.StatusInternalServerError)
		return
	}

	server.Listener.HTTPRequestReceived(NewRequestFromRequest(req), res)
}
