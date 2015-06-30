// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

type SSDPResponse struct {
	*SSDPPacket
}

// NewSSDPPacket returns a new SSDPPacket.
func NewSSDPResponse() *SSDPResponse {
	ssdpRes := &SSDPResponse{}
	ssdpRes.SSDPPacket = NewSSDPPacket()
	return ssdpRes
}

// NewSSDPPacket returns a new SSDPPacket.
func NewSSDPResponseFromBytes(bytes []byte) (*SSDPResponse, error) {
	ssdpPkt, err := NewSSDPPacketFromBytes(bytes)
	if err != nil {
		return nil, err
	}
	ssdpRes := NewSSDPResponse()
	ssdpRes.SSDPPacket = ssdpPkt
	return ssdpRes, nil
}
