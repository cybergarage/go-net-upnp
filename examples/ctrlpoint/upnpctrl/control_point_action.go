// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnpctrl

type ContolPointActionFunc func(*ControlPoint) bool

// A ControlPoint represents a ControlPoint.
type ControlPointAction struct {
	Key  int
	Desc string
	Func ContolPointActionFunc
}

// NewControlPoint returns a new Client.
func NewControlPointAction(key int, desc string, actionFunc ContolPointActionFunc) *ControlPointAction {
	cpAction := &ControlPointAction{key, desc, actionFunc}
	return cpAction
}
