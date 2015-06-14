// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"errors"
	"testing"
)

const (
	nullSSDPPacketError = "SSDPPacket is null"
)

func TestNewSSDPPacket(t *testing.T) {
	sspdPkt := NewSSDPPacket()
	if sspdPkt == nil {
		t.Error(errors.New(nullSSDPPacketError))
	}
}
