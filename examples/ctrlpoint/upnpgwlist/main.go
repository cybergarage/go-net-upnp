// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
upnpgwdump dumps prints all internet gatway devices, InternetGatewayDevice:1, in the local network.

	NAME
	upnpgwdump

	SYNOPSIS
	upnpdump [OPTIONS]

	DESCRIPTION
	upnpgwdump is a utility to dump SSDP messages.


	OPTIONS
	-v [0 | 1] : Enable verbose output.

	EXIT STATUS
	  Return EXIT_SUCCESS or EXIT_FAILURE

	EXAMPLES
	  The following is how to enable the verbose output
	    upnpgwdump -v 1
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-net-upnp/net/upnp"
)

func printGatewayDevice(n int, dev *GatewayDevice) {
	fmt.Printf("[%d] %s (%s)\n", n, dev.FriendlyName, dev.LocationURL)

	// ExternalIPAddress

	addr, err := dev.GetExternalIPAddress()
	if err == nil {
		fmt.Printf("  External IP address = %s\n", addr)
	}

	// GetTotalBytesReceived

	recvBytes, err := dev.GetTotalBytesReceived()
	if err == nil {
		fmt.Printf("  Total Bytes Received = %s\n", recvBytes)
	}

	// GetTotalBytesSent

	sentBytes, err := dev.GetTotalBytesSent()
	if err == nil {
		fmt.Printf("  Total Bytes Sent = %s\n", sentBytes)
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

	gwDevs := ctrlPoint.GetRootDevicesByType(InternetGatewayDeviceType)
	if len(gwDevs) == 0 {
		fmt.Printf("Internet gateway device is not found !!\n")
		os.Exit(0)
	}

	for n, dev := range gwDevs {
		gwDev := NewGatewayDevice(dev)
		printGatewayDevice(n, gwDev)
	}

	os.Exit(0)
}
