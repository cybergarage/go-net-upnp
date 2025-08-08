// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	errorDeviceCouldNotAdded = "device (%s:%s) couldn't be added"
	errorDeviceMapSize       = "device map size is invalid = %d: expected %d"
	errorDeviceNotFound      = "device (%s:%s) is not found"
)

func TestNewDeviceMap(t *testing.T) {
	NewDeviceMap()
}

func TestDeviceMapCount(t *testing.T) {
	devMap := NewDeviceMap()
	if devMap.Size() != 0 {
		t.Errorf(errorDeviceMapSize, devMap.Size(), 0)
	}

	// Create devices

	typeCnt := (rand.Int() % 10) + 10
	types := make([]string, typeCnt)
	for n := range typeCnt {
		types[n] = fmt.Sprintf("type%d", n)
	}

	udnCnt := (rand.Int() % 10) + 10
	udns := make([]string, udnCnt)
	for n := range udnCnt {
		udns[n] = fmt.Sprintf("udn%d", n)
	}

	devCnt := typeCnt * udnCnt
	devs := make([]*Device, devCnt)
	n := 0
	for _, ty := range types {
		for _, udn := range udns {
			dev := NewDevice()
			dev.DeviceType = ty
			dev.UDN = udn
			devs[n] = dev
			n++
		}
	}

	// Add devices

	for n, dev := range devs {
		if devMap.Size() != n {
			t.Errorf(errorDeviceMapSize, devMap.Size(), n)
		}
		ok := devMap.AddDevice(devs[n])
		if !ok {
			t.Errorf(errorDeviceCouldNotAdded, dev.DeviceType, dev.UDN)
		}
		if devMap.Size() != (n + 1) {
			t.Errorf(errorDeviceMapSize, devMap.Size(), (n + 1))
		}
	}

	if devMap.Size() != len(devs) {
		t.Errorf(errorDeviceMapSize, devMap.Size(), len(devs))
	}

	// GetAllDevices

	mapDevs := devMap.GetAllDevices()
	if devMap.Size() != len(mapDevs) {
		t.Errorf(errorDeviceMapSize, devMap.Size(), len(mapDevs))
	}

	// GetDevicesByType

	for _, ty := range types {
		mapDevs := devMap.GetDevicesByType(ty)
		if udnCnt != len(mapDevs) {
			t.Errorf(errorDeviceMapSize, udnCnt, len(mapDevs))
		}
	}

	// Find devices

	for _, ty := range types {
		for _, udn := range udns {
			_, ok := devMap.FindDeviceByTypeAndUDN(ty, udn)
			if !ok {
				t.Errorf(errorDeviceNotFound, ty, udn)
			}
		}
	}
}
