// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
lightdev is a sample implementation for UPnP device.

        NAME
        lightdev

        SYNOPSIS
        lightdev [OPTIONS]

        DESCRIPTION
        lightdev is a sample implmentation of UPnP Standardized DCP, BinaryLight:1

        OPTIONS
        -v : *level* Enable verbose output.

        RETURN VALUE
          Return EXIT_SUCCESS or EXIT_FAILURE
*/

package main

import (
	"bufio"
	"os"

	"net/upnp"
	"net/upnp/log"
)

func handleInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		key, err := reader.ReadByte()
		if err != nil {
			continue
		}
		if key == 'q' {
			return
		}
	}
}

func main() {
	logger := log.NewStdoutLogger(log.LoggerLevelTrace)
	log.SetSharedLogger(logger)

	dev, err := upnp.NewDeviceFromDescription(BinaryLightDeviceDescription)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	err = dev.Start()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer dev.Stop()

	handleInput()

	os.Exit(0)
}
