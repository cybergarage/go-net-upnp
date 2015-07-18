// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"testing"
)

type PacketSetHeaderFunc func(*Packet) func(string, string) error
type PacketGetHeaderFunc func(*Packet) func(string) (string, bool)

func TestNewPacket(t *testing.T) {
	NewPacket()
}
