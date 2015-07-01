// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import ()

// A ControlPoint represents a ControlPoint.
type ControlPointActionManager struct {
	Commands map[int]*ControlPointAction
}

// NewControlPoint returns a new Client.
func NewControlPointActionManager() *ControlPointActionManager {
	actionMgr := &ControlPointActionManager{}
	actionMgr.Commands = make(map[int]*ControlPointAction)

	actionMgr.Commands[Q_KEY] = NewControlPointAction(Q_KEY, Q_DESC, QuitAction)
	actionMgr.Commands[H_KEY] = NewControlPointAction(H_KEY, H_DESC, HelpAction)

	return actionMgr
}

func (self *ControlPointActionManager) DoAction(ctrlPoint *ControlPoint, key int) bool {
	action, ok := ctrlPoint.Commands[key]
	if !ok {
		action, _ = ctrlPoint.Commands[H_KEY]
	}

	return action.Func(ctrlPoint)
}
