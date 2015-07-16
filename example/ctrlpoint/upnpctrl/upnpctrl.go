// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
upnpdump dumps SSDP messages in the local network.

        NAME
        upnpdump

        SYNOPSIS
        upnpdump [OPTIONS]

        DESCRIPTION
        upnpdump is a utility to dump SSDP messages.


        OPTIONS
        -v : *level* Enable verbose output.

        RETURN VALUE
          Return EXIT_SUCCESS or EXIT_FAILURE
*/

package main

import (
	"bufio"
	"fmt"
	"os"

	"net/upnp/log"
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
			continue
		}
		if !ctrlPoint.DoAction(int(key)) {
			return
		}
	}
}

func main() {
	/*logger := */ log.NewStdoutLogger(log.LoggerLevelTrace)
	//log.SetSharedLogger(logger)

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
