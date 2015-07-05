// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import ()

// DeviceMap manages devices by UDN
type DeviceUdnMap map[string]*Device

// DeviceUdnMap returns a new device map.
func NewDeviceUdnMap() *DeviceUdnMap {
	devMap := make(DeviceUdnMap)
	return &devMap
}

// AddDevice adds a specified device.
func (self *DeviceUdnMap) AddDevice(dev *Device) bool {
	udn := dev.Description.UDN
	if len(udn) <= 0 {
		return false
	}

	(*self)[udn] = dev

	return true
}

// HasDevice check whether the specified device is added.
func (self *DeviceUdnMap) FindDeviceByUDN(udn string) (*Device, bool) {
	if len(udn) <= 0 {
		return nil, false
	}

	dev, ok := (*self)[udn]
	if !ok {
		return nil, false
	}

	return dev, true
}

// HasDevice check whether the specified device is added.
func (self *DeviceUdnMap) HasDeviceByUDN(udn string) bool {
	_, ok := self.FindDeviceByUDN(udn)
	return ok
}
