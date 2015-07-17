// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import ()

// DeviceMap manages devices by UDN
type DeviceMap map[string]map[string]*Device

// DeviceMap returns a new device map.
func NewDeviceMap() DeviceMap {
	devMap := make(DeviceMap)
	return devMap
}

// AddDevice adds a specified device.
func (self *DeviceMap) AddDevice(dev *Device) bool {
	if dev == nil {
		return false
	}

	udn := dev.UDN
	if len(udn) <= 0 {
		return false
	}

	deviceType := dev.DeviceType
	if len(deviceType) <= 0 {
		return false
	}

	(*self)[deviceType][udn] = dev

	return true
}

// GetAllDevices returns all devices.
func (self *DeviceMap) GetAllDevices() []*Device {
	devs := make([]*Device, 0)

	for _, typeDevs := range *self {
		for _, dev := range typeDevs {
			devs = append(devs, dev)
		}
	}

	return devs
}

// GetDevicesByType returns only devices of the specified device type.
func (self *DeviceMap) GetDevicesByType(deviceType string) []*Device {
	devs := make([]*Device, 0)

	typeDevs, ok := (*self)[deviceType]
	if !ok {
		return devs
	}

	for _, dev := range typeDevs {
		devs = append(devs, dev)
	}

	return devs
}

// FindDeviceByTypeAndUDN find a device of the specified device type and udn.
func (self *DeviceMap) FindDeviceByTypeAndUDN(deviceType string, udn string) (*Device, bool) {
	if len(deviceType) <= 0 || len(udn) <= 0 {
		return nil, false
	}

	devs, ok := (*self)[deviceType]
	if !ok {
		return nil, false
	}

	dev, ok := devs[udn]
	if !ok {
		return nil, false
	}

	return dev, true
}

// HasDeviceByTypeAndUDN check whether a device of the specified device type and udn exits.
func (self *DeviceMap) HasDeviceByTypeAndUDN(deviceType string, udn string) bool {
	_, ok := self.FindDeviceByTypeAndUDN(deviceType, udn)
	return ok
}

// HasDevice adds a specified device.
func (self *DeviceMap) HasDevice(dev *Device) bool {
	if dev == nil {
		return false
	}
	return self.HasDeviceByTypeAndUDN(dev.DeviceType, dev.UDN)
}
