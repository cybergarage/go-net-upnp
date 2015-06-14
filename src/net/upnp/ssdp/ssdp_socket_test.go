// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"errors"
	"testing"
)

const (
	nullSSDPSocketError = "SSDPSocket is null"
)

func TestNewSSDPSocket(t *testing.T) {
	sspdPkt := NewSSDPSocket()
	if sspdPkt == nil {
		t.Error(errors.New(nullSSDPSocketError))
	}
}
