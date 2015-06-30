// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/upnp"
	"net/upnp/ssdp"
	"os"
	"fmt"
)
// A ControlPoint represents a ControlPoint.
type ControlPoint struct {
	*upnp.ControlPoint
}

// NewControlPoint returns a new Client.
func NewControlPoint() *ControlPoint {
	cp := &ControlPoint{}
	cp.ControlPoint = upnp.NewControlPoint()
	cp.ControlPoint.Listener = cp
	return cp
}

func (self *ControlPoint) DeviceNotifyReceived(ssdpPkt *ssdp.SSDPPacket) {
	os.Stdout.WriteString(fmt.Sprintf("%s\n", ssdpPkt.ToString()))
}
