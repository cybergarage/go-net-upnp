// Copyright 2015 Satoshi Konno. All rights reserved.
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

// NewSSDPResponseFromBytes returns a new SSDPPacket from the specified bytes.
func NewSSDPResponseFromBytes(bytes []byte) (*SSDPResponse, error) {
	ssdpPkt, err := NewSSDPPacketFromBytes(bytes)
	if err != nil {
		return nil, err
	}
	ssdpRes := NewSSDPResponse()
	ssdpRes.SSDPPacket = ssdpPkt
	return ssdpRes, nil
}

// NewSSDPRequestFromString returns a new SSDPPacket from the specified string.
func NewSSDPResponseFromString(packet string) (*SSDPResponse, error) {
	return NewSSDPResponseFromBytes([]byte(packet))
}
