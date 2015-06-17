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
  "net/upnp"
	"os"
)

func main() {
	os.Exit(0)
}
