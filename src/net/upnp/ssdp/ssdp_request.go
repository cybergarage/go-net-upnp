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

// NewSSDPPacket returns a new SSDPPacket.
func NewSSDPRequestFromBytes(bytes []byte) (*SSDPRequest, error) {
	ssdpPkt, err := NewSSDPPacketFromBytes(bytes)
	if err != nil {
		return nil, err
	}
	ssdpReq := NewSSDPRequest()
	ssdpReq.SSDPPacket = ssdpPkt
	return ssdpReq, nil
}
