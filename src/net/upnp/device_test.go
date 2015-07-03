// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"fmt"
	gohttp "net/http"
	"testing"
)

func TestNewDevice(t *testing.T) {
	dev := NewDevice()

	err := dev.Start()
	if err != nil {
		t.Error(err)
	}

	res, err := gohttp.Get(fmt.Sprintf("http://localhost:%d/", dev.Port))
	if err != nil {
		t.Error(err)
	}

	if (dev.Port < DEVICE_DEFAULT_PORT_BASE) || (DEVICE_DEFAULT_PORT_MAX < dev.Port) {
		t.Errorf("got invalid port = [%d] : expected : [%d]~[%d]", dev.Port, DEVICE_DEFAULT_PORT_BASE, DEVICE_DEFAULT_PORT_MAX)
	}

	if res.StatusCode != gohttp.StatusInternalServerError {
		t.Errorf("got invalid test code = [%d] : expected : [%d]", res.StatusCode, gohttp.StatusInternalServerError)
	}

	err = dev.Stop()
	if err != nil {
		t.Error(err)
	}
}
