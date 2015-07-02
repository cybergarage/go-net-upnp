// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

type SSDPRequest struct {
	*SSDPPacket
}

// NewSSDPPacket returns a new SSDPPacket.
func NewSSDPRequest() *SSDPRequest {
	ssdpReq := &SSDPRequest{}
	ssdpReq.SSDPPacket = NewSSDPPacket()
	return ssdpReq
}

// NewSSDPRequestFromBytes returns a new SSDPPacket from the specified bytes.
func NewSSDPRequestFromBytes(bytes []byte) (*SSDPRequest, error) {
	ssdpPkt, err := NewSSDPPacketFromBytes(bytes)
	if err != nil {
		return nil, err
	}
	ssdpReq := NewSSDPRequest()
	ssdpReq.SSDPPacket = ssdpPkt
	return ssdpReq, nil
}

// NewSSDPRequestFromString returns a new SSDPPacket from the specified string.
func NewSSDPRequestFromString(packet string) (*SSDPRequest, error) {
	return NewSSDPRequestFromBytes([]byte(packet))
}

// NewSSDPRequestFromPacket returns a new SSDPPacket from the specified packet.
func NewSSDPRequestFromPacket(packet *SSDPPacket) (*SSDPRequest, error) {
	ssdpReq := &SSDPRequest{}
	ssdpReq.SSDPPacket = packet
	return ssdpReq, nil
}
