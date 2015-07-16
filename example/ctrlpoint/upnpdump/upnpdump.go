// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
upnpdump dumps prints all devices in the local network.

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
	"os"
	"time"

	"net/upnp"
	"net/upnp/log"
)

func main() {
	//logger := log.NewStdoutLogger(log.LoggerLevelTrace)
	//log.SetSharedLogger(logger)

	// Start a control point

	ctrlPoint := upnp.NewControlPoint()
	err := ctrlPoint.Start()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer ctrlPoint.Stop()

	// Search root devices

	err = ctrlPoint.SearchRootDevice()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Sleep until all search responses are received

	time.Sleep(time.Duration(ctrlPoint.SearchMX) * time.Second)

	// Print basic descriptions of found devices

	if len(ctrlPoint.GetRootDevices()) == 0 {
		fmt.Printf("Device not found !!")
		os.Exit(0)
	}

	for n, dev := range ctrlPoint.GetRootDevices() {
		devURL := dev.LocationURL

		presentationURL := dev.PresentationURL
		if 0 < len(presentationURL) {
			url, err := dev.GetAbsoluteURL(presentationURL)
			if err == nil {
				devURL = url.String()
			}
		}

		fmt.Printf("[%d] '%s', '%s', %s\n", n, dev.FriendlyName, dev.DeviceType, devURL)
	}

	os.Exit(0)
}
