// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"net"
	"strings"
)

const (
	errorAvailableAddressNotFound = "Available address not found"
	errorAvailableInterfaceFound  = "Available interface not found"
)

func IsIPv6Address(addr string) bool {
	if 0 < strings.Index(addr, ":") {
		return true
	}

	return false
}

func GetInterfaceAddress(ifi net.Interface) (string, error) {
	addrs, err := ifi.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		saddr := strings.Split(addr.String(), "/")
		if len(saddr) < 2 {
			continue
		}

		// Disabled IPv6 interface
		if IsIPv6Address(saddr[0]) {
			continue
		}

		return saddr[0], nil
	}

	return "", errors.New(errorAvailableAddressNotFound)
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

		_, addrErr := GetInterfaceAddress(localIf)
		if addrErr != nil {
			continue
		}

		useIfs = append(useIfs, localIf)
	}

	if len(useIfs) <= 0 {
		return useIfs, errors.New(errorAvailableInterfaceFound)
	}

	return useIfs, err
}
