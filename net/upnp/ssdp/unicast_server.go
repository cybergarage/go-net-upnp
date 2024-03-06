// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"errors"
	"net"

	"github.com/cybergarage/go-logger/log"
)

// A UnicastListener represents a listener for UnicastServer.
type UnicastListener interface {
	DeviceResponseReceived(ssdpRes *Response)
}

// A UnicastServer represents a packet of SSDP.
type UnicastServer struct {
	Socket    *UnicastSocket
	Listener  UnicastListener
	Interface net.Interface
}

// NewUnicastServer returns a new UnicastServer.
func NewUnicastServer() *UnicastServer {
	server := &UnicastServer{}
	server.Socket = NewUnicastSocket()
	server.Listener = nil
	return server
}

// Start starts this server.
func (server *UnicastServer) Start(ifi net.Interface, port int) error {
	err := server.Socket.Bind(ifi, port)
	if err != nil {
		return err
	}
	server.Interface = ifi
	go handleSSDPUnicastConnection(server)
	return nil
}

// Stop stops this server.
func (server *UnicastServer) Stop() error {
	err := server.Socket.Close()
	if err != nil {
		return err
	}
	return nil
}

// Search sends a M-SEARCH request of the specified ST.
func (server *UnicastServer) Search(st string, mx int) error {
	ssdpReq, err := NewSearchRequest(st, mx)
	if err != nil {
		return err
	}

	_, err = server.Socket.WriteRequest(ssdpReq)

	return err
}

func handleSSDPUnicastConnection(server *UnicastServer) {
	for {
		ssdpPkt, err := server.Socket.Read()
		if err != nil {
			if !errors.Is(err, net.ErrClosed) {
				log.Error(err)
			}
			break
		}

		if server.Listener != nil {
			ssdpRes, _ := NewResponseFromPacket(ssdpPkt)
			server.Listener.DeviceResponseReceived(ssdpRes)
		}
	}
}
