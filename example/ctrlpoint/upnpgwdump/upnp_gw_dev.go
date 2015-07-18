// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/upnp"
)

type GatewayDevice struct {
	*upnp.Device
}

func NewGatewayDevice(dev *upnp.Device) *GatewayDevice {
	gwDev := &GatewayDevice{Device: dev}
	return gwDev
}
