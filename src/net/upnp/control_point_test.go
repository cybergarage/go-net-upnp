// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"errors"
	"testing"
)

const (
	nullControlPointError = "ControlPoint is null"
)

func TesttNewClient(t *testing.T) {
	cp := NewControlPoint()
	if cp == nil {
		t.Error(errors.New(nullControlPointError))
	}
}
