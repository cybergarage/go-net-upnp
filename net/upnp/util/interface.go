// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"net"
	"strings"
)

const (
	errorAvailableAddressNotFound = "available address not found"
	errorAvailableInterfaceFound  = "available interface not found"
)

func IsIPv6Address(addr string) bool {
	return 0 < strings.Index(addr, ":")
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

	if len(useIfs) == 0 {
		return useIfs, errors.New(errorAvailableInterfaceFound)
	}

	return useIfs, err
}

func getMatchAddressBlockCount(ifAddr string, targetAddr string) int {
	const addrSep = "."
	targetAddrs := strings.Split(targetAddr, addrSep)
	ifAddrs := strings.Split(ifAddr, addrSep)

	if len(targetAddrs) != len(ifAddrs) {
		return -1
	}

	addrSize := len(targetAddrs)
	for n := 0; n < len(targetAddrs); n++ {
		if targetAddrs[n] != ifAddrs[n] {
			return n
		}
	}

	return addrSize
}

func GetAvailableInterfaceForAddr(fromAddr string) (net.Interface, error) {
	ifis, err := GetAvailableInterfaces()
	if err != nil {
		return net.Interface{}, err
	}

	switch len(ifis) {
	case 0:
		return net.Interface{}, errors.New(errorAvailableInterfaceFound)
	case 1:
		return ifis[0], nil
	}

	ifAddrs := make([]string, len(ifis))
	for n := 0; n < len(ifAddrs); n++ {
		ifAddrs[n], _ = GetInterfaceAddress(ifis[n])
	}

	selIf := ifis[0]
	selIfMatchBlocks := getMatchAddressBlockCount(fromAddr, ifAddrs[0])
	for n := 0; n < len(ifAddrs); n++ {
		matchBlocks := getMatchAddressBlockCount(fromAddr, ifAddrs[n])
		if matchBlocks < selIfMatchBlocks {
			continue
		}
		selIf = ifis[n]
		selIfMatchBlocks = matchBlocks
	}

	return selIf, nil
}
