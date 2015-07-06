// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func printActionMessage(msg string) {
	os.Stdout.WriteString(fmt.Sprintf("%s\n", msg))
}

////////////////////////////////////////
// h : Help
////////////////////////////////////////

const (
	H_KEY  = 'h'
	H_DESC = "print this message"
)

func HelpAction(cp *ControlPoint) bool {
	for key, action := range cp.ControlPointActionManager.Commands {
		printActionMessage(fmt.Sprintf("'%c' : %s", key, action.Desc))
	}
	return true
}

////////////////////////////////////////
// q : Quit
////////////////////////////////////////

const (
	Q_KEY  = 'q'
	Q_DESC = "quit"
)

func QuitAction(cp *ControlPoint) bool {
	return false
}

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
// s : Print
////////////////////////////////////////

const (
	P_KEY  = 'p'
	P_DESC = "print found devices"
)

func PrintAction(cp *ControlPoint) bool {
	foundDevs := cp.GetRootDevices()

	printActionMessage(fmt.Sprintf("==== Devices (%d) ====", len(foundDevs)))
	for n, dev := range foundDevs {
		printActionMessage(fmt.Sprintf("[%d] '%s', '%s'", n, dev.FriendlyName, dev.DeviceType))
	}

	return true
}
