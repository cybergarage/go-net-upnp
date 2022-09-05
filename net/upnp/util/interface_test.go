// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"testing"
)

func TestNewMulticastServerList(t *testing.T) {
	ifs, err := GetAvailableInterfaces()
	if err != nil {
		t.Error(err)
	}
	if len(ifs) == 0 {
		t.Errorf("available interface is not found")
	}
}
