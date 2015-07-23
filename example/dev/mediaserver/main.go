// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
mediaserver is a sample implementation of UPnP standard device, MediaServer:1.

        NAME
        mediaserver

        SYNOPSIS
        mediaserver [OPTIONS]

        DESCRIPTION
        mediaserver is a sample implmentation of UPnP Standardized DCP, MediaServer:1

        OPTIONS
        -v : *level* Enable verbose output.

        RETURN VALUE
          Return EXIT_SUCCESS or EXIT_FAILURE
*/
package media

import (
	"bufio"
	"os"

	"github.com/cybergarage/go-net-upnp/net/upnp/log"
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

	dev, err := NewMediaServer()
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
