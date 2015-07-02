// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/upnp"
	"net/upnp/ssdp"
	"os"
)

// A ControlPoint represents a ControlPoint.
type ControlPoint struct {
	*upnp.ControlPoint
	*ControlPointActionManager
}

// NewControlPoint returns a new Client.
func NewControlPoint() *ControlPoint {
	cp := &ControlPoint{}

	cp.ControlPoint = upnp.NewControlPoint()
	cp.ControlPointActionManager = NewControlPointActionManager()
	cp.ControlPoint.Listener = cp

	return cp
}

func (self *ControlPoint) DeviceNotifyReceived(ssdpReq *ssdp.Request) {
	os.Stdout.WriteString(fmt.Sprintf("%s\n", ssdpReq.String()))
}

func (self *ControlPoint) DeviceSearchReceived(ssdpReq *ssdp.Request) {
	os.Stdout.WriteString(fmt.Sprintf("%s\n", ssdpReq.String()))
}

func (self *ControlPoint) DoAction(key int) bool {
	return self.ControlPointActionManager.DoAction(self, key)
}
