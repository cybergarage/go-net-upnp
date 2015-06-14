// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"errors"
	"testing"
)

const (
	nullDeviceError = "Device is null"
)

func TestNewDevice(t *testing.T) {
	dev := NewDevice()
	if dev == nil {
		t.Error(errors.New(nullDeviceError))
	}
}
