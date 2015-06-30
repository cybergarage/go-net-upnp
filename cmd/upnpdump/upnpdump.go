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
	"os"
	"fmt"
)

func outputError(err error) {
	os.Stderr.WriteString(fmt.Sprintf("%s\n", err.Error()))
}

func main() {
	ctrlPoint := NewControlPoint()
	err := ctrlPoint.Start()
	if err != nil {
		outputError(err)
		os.Exit(1)
	}
	defer ctrlPoint.Stop()

	os.Exit(0)
}
