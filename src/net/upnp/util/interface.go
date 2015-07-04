// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"net"
	"strings"
)

func getInterfaceStringAddrs(addrs []net.Addr) (string, error) {
	for _, addr := range addrs {
		saddr := strings.Split(addr.String(), "/")
		if len(saddr) < 2 {
			continue
		}

		// Disabled IPv6 interface
		if 0 < strings.Index(saddr[0], ":") {
			continue
		}

		return saddr[0], nil
	}

	return "", errors.New("Available address not found")
}

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

		_, err = getInterfaceStringAddrs(addrs)
		if err != nil {
			continue
		}

		useIfs = append(useIfs, localIf)

	}

	return useIfs, err
}

func GetAvailableInterfaceAddresses() ([]string, error) {
	useIfStrAddrs := make([]string, 0)

	useIfs, err := GetAvailableInterfaces()
	if err != nil {
		return useIfStrAddrs, err
	}

	for _, useIf := range useIfs {
		useIfAddrs, err := useIf.Addrs()
		if err != nil {
			continue
		}

		useIfStrAddr, err := getInterfaceStringAddrs(useIfAddrs)
		if err != nil {
			continue
		}

		useIfStrAddrs = append(useIfStrAddrs, useIfStrAddr)

	}

	return useIfStrAddrs, err
}
