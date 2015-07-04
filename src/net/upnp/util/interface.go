// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"net"
	"strings"
)

func GetAvailableInterfaces() ([]net.Interface, error) {
	useIfs := make([]net.Interface, 0)

	localIfs, err := net.Interfaces()
	if err != nil {
		return useIfs, err
	}

	for _, localIf := range localIfs {
		if (localIf.Flags & net.FlagLoopback) != 0 {
			continue
		}
		if (localIf.Flags & net.FlagUp) == 0 {
			continue
		}
		if (localIf.Flags & net.FlagMulticast) == 0 {
			continue
		}
		addrs, err := localIf.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			saddr := strings.Split(addr.String(), "/")
			if len(saddr) < 2 {
				continue
			}

			// Disabled IPv6 interface
			if 0 < strings.Index(saddr[0], ":") {
				continue
			}

			useIfs = append(useIfs, localIf)
			break
		}

	}

	return useIfs, err
}
