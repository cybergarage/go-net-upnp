// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

type Request struct {
	*Packet
}

// NewPacket returns a new Packet.
func NewRequest() *Request {
	ssdpReq := &Request{}
	ssdpReq.Packet = NewPacket()
	return ssdpReq
}

// NewRequestFromBytes returns a new Packet from the specified bytes.
func NewRequestFromBytes(bytes []byte) (*Request, error) {
	ssdpPkt, err := NewPacketFromBytes(bytes)
	if err != nil {
		return nil, err
	}
	ssdpReq := NewRequest()
	ssdpReq.Packet = ssdpPkt
	return ssdpReq, nil
}

// NewRequestFromString returns a new Packet from the specified string.
func NewRequestFromString(packet string) (*Request, error) {
	return NewRequestFromBytes([]byte(packet))
}

// NewRequestFromPacket returns a new Packet from the specified packet.
func NewRequestFromPacket(packet *Packet) (*Request, error) {
	ssdpReq := &Request{}
	ssdpReq.Packet = packet
	return ssdpReq, nil
}
