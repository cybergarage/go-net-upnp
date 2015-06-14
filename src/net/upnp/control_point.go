// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

// A ControlPoint represents a clinet.
type ControlPoint struct {
	RootDevices []Device
}

// NewControlPoint returns a new Client.
func NewControlPoint() *ControlPoint {
	cp := &ControlPoint{}
	cp.RootDevices = make([]Device, 0)
	return cp
}
