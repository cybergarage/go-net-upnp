// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

////////////////////////////////////////
// s : Search
////////////////////////////////////////

const (
	S_KEY  = 's'
	S_DESC = "search root devices"
)

func SearchAction(cp *ControlPoint) bool {
	err := cp.SearchRootDevice()
	if err != nil {
		return false
	}
	return true
}

////////////////////////////////////////
// h : Help
////////////////////////////////////////

const (
	H_KEY  = 'h'
	H_DESC = "print this message"
)

func QuitAction(cp *ControlPoint) bool {
	return false
}

////////////////////////////////////////
// q : Quit
////////////////////////////////////////

const (
	Q_KEY  = 'q'
	Q_DESC = "quit"
)

func HelpAction(cp *ControlPoint) bool {
	for key, action := range cp.ControlPointActionManager.Commands {
		os.Stderr.WriteString(fmt.Sprintf("'%c' : %s\n", key, action.Desc))

	}
	return true
}
