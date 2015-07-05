// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"testing"
)

const (
	testDeviceDesecription = "<root></root>"
)

func TestParseDeviceDescription(t *testing.T) {
	devRoot := DeviceDescriptionRoot{}
	err := xml.Unmarshal([]byte(testDeviceDesecription), &devRoot)
	if err != nil {
		t.Error(err)
	}
}
