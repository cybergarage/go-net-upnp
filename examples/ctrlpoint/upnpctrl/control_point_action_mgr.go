// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// ControlPointActionManager manages available ControlPointAction commands.
type ControlPointActionManager struct {
	Commands map[int]*ControlPointAction
}

// NewControlPointActionManager returns a new ControlPointActionManager.
func NewControlPointActionManager() *ControlPointActionManager {
	actionMgr := &ControlPointActionManager{}
	actionMgr.initDefaultActions()
	return actionMgr
}

func (cp *ControlPointActionManager) initDefaultActions() error {
	cp.Commands = make(map[int]*ControlPointAction)

	cp.AddAction(Q_KEY, Q_DESC, QuitAction)
	cp.AddAction(H_KEY, H_DESC, HelpAction)
	cp.AddAction(S_KEY, S_DESC, SearchAction)
	cp.AddAction(P_KEY, P_DESC, PrintAction)
	cp.AddAction(A_KEY, A_DESC, PostAction)

	return nil
}

func (cp *ControlPointActionManager) AddAction(key int, desc string, actionFunc ContolPointActionFunc) error {
	action := NewControlPointAction(key, desc, actionFunc)
	cp.Commands[key] = action
	return nil
}

func (cp *ControlPointActionManager) DoAction(ctrlPoint *ControlPoint, key int) bool {
	action, ok := ctrlPoint.Commands[key]
	if !ok {
		action = ctrlPoint.Commands[H_KEY]
	}

	return action.Func(ctrlPoint)
}
