// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"time"

	"github.com/cybergarage/go-net-upnp/net/upnp/http"
)

type Response struct {
	*Packet
}

// NewPacket returns a new Packet.
func NewResponse() *Response {
	ssdpRes := &Response{}
	ssdpRes.Packet = NewPacket()

	ssdpRes.SetStatusCode(http.StatusOK)
	ssdpRes.SetServer(http.GetServerName())
	ssdpRes.SetEXT("")
	ssdpRes.SetDate(time.Now().Format(time.RFC1123))

	return ssdpRes
}

// NewResponseFromBytes returns a new Packet from the specified bytes.
func NewResponseFromBytes(bytes []byte) (*Response, error) {
	ssdpPkt, err := NewPacketFromBytes(bytes)
	if err != nil {
		return nil, err
	}
	ssdpRes := NewResponse()
	ssdpRes.Packet = ssdpPkt
	return ssdpRes, nil
}

// NewSSDPRequestFromString returns a new Packet from the specified string.
func NewResponseFromString(packet string) (*Response, error) {
	return NewResponseFromBytes([]byte(packet))
}

// NewResponseFromPacket returns a new Packet from the specified packet.
func NewResponseFromPacket(packet *Packet) (*Response, error) {
	ssdpRes := &Response{}
	ssdpRes.Packet = packet
	return ssdpRes, nil
}
