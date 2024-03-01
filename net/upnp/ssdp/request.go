// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

type Request struct {
	*Packet
}

// NewPacket returns a new Request.
func NewRequest() *Request {
	ssdpReq := &Request{}
	ssdpReq.Packet = NewPacket()
	return ssdpReq
}

// NewRequestFromBytes returns a new Request from the specified bytes.
func NewRequestFromBytes(bytes []byte) (*Request, error) {
	ssdpPkt, err := NewPacketFromBytes(bytes)
	if err != nil {
		return nil, err
	}
	ssdpReq := NewRequest()
	ssdpReq.Packet = ssdpPkt
	return ssdpReq, nil
}

// NewRequestFromString returns a new Request from the specified string.
func NewRequestFromString(packet string) (*Request, error) {
	return NewRequestFromBytes([]byte(packet))
}

// NewRequestFromPacket returns a new Request from the specified packet.
func NewRequestFromPacket(packet *Packet) (*Request, error) {
	ssdpReq := &Request{}
	ssdpReq.Packet = packet
	return ssdpReq, nil
}

// NewSearchRequest a new Request from the specified bytes.
func NewSearchRequest(st string, mx int) (*Request, error) {
	ssdpReq := NewRequest()

	ssdpReq.SetMethod(MSearch)
	ssdpReq.SetHost(MulticastAddress)
	ssdpReq.SetST(st)
	ssdpReq.SetMX(mx)
	ssdpReq.SetMAN(Discover)

	return ssdpReq, nil
}

func (req *Request) IsDiscover() bool {
	return req.IsHeaderString(MAN, Discover)
}

func (req *Request) IsRootDevice() bool {
	return req.IsHeaderString(ST, RootDevice)
}

func (req *Request) IsAlive() bool {
	return req.IsHeaderString(NTS, NTSAlive)
}

func (req *Request) IsByeBye() bool {
	return req.IsHeaderString(NTS, NTSByeBye)
}

func (req *Request) IsUpdate() bool {
	return req.IsHeaderString(NTS, NTSUpdate)
}
