// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
upnpdump prints SSDP packets in the local network.

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
	    upntctrl -v 1
*/
package upnpdump

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cybergarage/go-logger/log"
)

func printHelp() {
	fmt.Printf("s : (s)earch root devices\n")
	fmt.Printf("q : (q)uit\n")
}

func handleInput(ctrlPoint *ControlPoint) {
	kb := bufio.NewReader(os.Stdin)

	for {
		keys, _, _ := kb.ReadLine()
		if len(keys) <= 0 {
			printHelp()
			continue
		}
		switch keys[0] {
		case 'q':
			return
		case 's':
			ctrlPoint.SearchRootDevice()
		default:
			printHelp()
		}
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

	handleInput(ctrlPoint)

	os.Exit(0)
}
