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
	"net/upnp/log"
	"os"
)

func handleInput(ctrlPoint *ControlPoint) {
	for {
		reader := bufio.NewReader(os.Stdin)
		key, err := reader.ReadByte()
		if err != nil {
			continue
		}
		if !ctrlPoint.DoAction(int(key)) {
			return
		}
	}
}

func main() {
	logger := log.NewStdoutLogger(log.LoggerLevelTrace)
	log.SetSharedLogger(logger)

	ctrlPoint := NewControlPoint()
	err := ctrlPoint.Start()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer ctrlPoint.Stop()

	handleInput(ctrlPoint)

	os.Exit(0)
}
