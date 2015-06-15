// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"errors"
	"testing"
)

const (
	nullSSDPServerError = "SSDPServer is null"
)

func TestNewSSDPServer(t *testing.T) {
	sspdPkt := NewSSDPServer()
	if sspdPkt == nil {
		t.Error(errors.New(nullSSDPServerError))
	}
}
