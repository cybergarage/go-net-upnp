// Copyright 2015 The go-net-upnp Authors. All rights reserved.
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
	"fmt"
	"net/upnp/log"
	"os"
)

func handleInput(ctrlPoint *ControlPoint) {
	for {
		var key int
		fmt.Scanf("%c", &key)
		if !ctrlPoint.DoAction(key) {
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
