// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
upnpsearch dumps prints all devices in the local network.

	NAME
	upnpdump

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
	    upnpsearch -v 1
*/
package upnpsearch

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-net-upnp/net/upnp"
)

func printDevice(n int, dev *upnp.Device) {
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
		log.SetSharedLogger(log.NewStdoutLogger(log.LevelTrace))
	}

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
		fmt.Printf("UPnP device is not found !!\n")
		os.Exit(0)
	}

	for n, dev := range ctrlPoint.GetRootDevices() {
		printDevice(n, dev)
	}

	os.Exit(0)
}
