// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNullDevice(t *testing.T) {
	dev := NewDevice()

	err := dev.Start()
	if err != nil {
		t.Error(err)
	}

	if (dev.Port < DeviceDefaultPortBase) || (DeviceDefaultPortMax < dev.Port) {
		t.Errorf(errorTestDeviceInvalidPortRange, dev.Port, DeviceDefaultPortBase, DeviceDefaultPortMax)
	}

	url := fmt.Sprintf("http://localhost:%d/", dev.Port)
	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf(errorTestDeviceInvalidStatusCode, url, res.StatusCode, http.StatusInternalServerError)
	}

	err = dev.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestSampleDevice(t *testing.T) {
	dev, err := NewTestDevice()

	if err != nil {
		t.Error(err)
	}

	// start device

	err = dev.Start()
	if err != nil {
		t.Error(err)
	}

	// check device
	if (dev.Port < DeviceDefaultPortBase) || (DeviceDefaultPortMax < dev.Port) {
		t.Errorf(errorTestDeviceInvalidPortRange, dev.Port, DeviceDefaultPortBase, DeviceDefaultPortMax)
	}

	devDescURL := fmt.Sprintf("http://localhost:%d%s", dev.Port, dev.DescriptionURL)
	res, err := http.Get(devDescURL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf(errorTestDeviceInvalidStatusCode, devDescURL, res.StatusCode, http.StatusOK)
	}

	// stop device
	err = dev.Stop()
	if err != nil {
		t.Error(err)
	}
}
