// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
upntctrl browses UPnP devices in the local network, and post any actions into the found devices.

	NAME
	upntctrl

	SYNOPSIS
	upnpdump [OPTIONS]

	DESCRIPTION
	upnpdump is a utility to dump SSDP messages.


	OPTIONS
	-v [0 | 1] : Enable verbose output.

	EXIT STATUS
	  Return EXIT_SUCCESS or EXIT_FAILURE

	EXAMPLES
	  The following is how to enable the verbose output
	    upntctrl -v 1
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cybergarage/go-logger/log"
)

const (
	errorNoInput = "no input"
)

var gKeyboardReader *bufio.Reader

func GetKeyboardReader() *bufio.Reader {
	return gKeyboardReader
}

func ReadKeyboardLine() (string, error) {
	b, _, err := GetKeyboardReader().ReadLine()
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func ReadKeyboardKey() (byte, error) {
	b, _, err := GetKeyboardReader().ReadLine()
	if err != nil {
		return 0, err
	}

	if len(b) == 0 {
		return 0, fmt.Errorf(errorNoInput)
	}

	return b[0], nil
}

func handleInput(ctrlPoint *ControlPoint) {
	for {
		key, err := ReadKeyboardKey()
		if err != nil {
			key = H_KEY
		}
		if !ctrlPoint.DoAction(int(key)) {
			return
		}
	}
}

func main() {
	// Set command line options

	verbose := flag.Int("v", 0, "Enable verbose mode [0|1]")
	flag.Usage = func() {
		cmd := strings.Split(os.Args[0], "/")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", cmd[len(cmd)-1])
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if 0 < *verbose {
		log.SetDefault(log.NewStdoutLogger(log.LevelTrace))
	}

	ctrlPoint := NewControlPoint()
	err := ctrlPoint.Start()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer ctrlPoint.Stop()

	err = ctrlPoint.SearchRootDevice()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	gKeyboardReader = bufio.NewReader(os.Stdin)
	handleInput(ctrlPoint)

	os.Exit(0)
}
