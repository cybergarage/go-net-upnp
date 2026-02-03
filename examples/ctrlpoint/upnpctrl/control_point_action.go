// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type ContolPointActionFunc func(*ControlPoint) bool

// ControlPointAction represents an executable command for a control point.
type ControlPointAction struct {
	Key  int
	Desc string
	Func ContolPointActionFunc
}

// NewControlPointAction returns a new ControlPointAction.
func NewControlPointAction(key int, desc string, actionFunc ContolPointActionFunc) *ControlPointAction {
	cpAction := &ControlPointAction{key, desc, actionFunc}
	return cpAction
}
