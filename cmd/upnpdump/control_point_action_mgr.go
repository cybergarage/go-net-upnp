// Copyright 2015 Satoshi Konno. All rights reserved.
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
	actionMgr.initDefaultActions()
	return actionMgr
}

func (self *ControlPointActionManager) initDefaultActions() error {
	self.Commands = make(map[int]*ControlPointAction)
	self.AddAction(Q_KEY, Q_DESC, QuitAction)
	self.AddAction(H_KEY, H_DESC, HelpAction)
	self.AddAction(S_KEY, S_DESC, SearchAction)
	return nil
}

func (self *ControlPointActionManager) AddAction(key int, desc string, actionFunc ContolPointActionFunc) error {
	action := NewControlPointAction(key, desc, actionFunc)
	self.Commands[key] = action
	return nil
}

func (self *ControlPointActionManager) DoAction(ctrlPoint *ControlPoint, key int) bool {
	action, ok := ctrlPoint.Commands[key]
	if !ok {
		action, _ = ctrlPoint.Commands[H_KEY]
	}

	return action.Func(ctrlPoint)
}
