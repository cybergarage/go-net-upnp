// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

// DeviceMap manages devices by UDN.
type DeviceMap map[string]map[string]*Device

// NewDeviceMap returns a new device map.
func NewDeviceMap() *DeviceMap {
	devMap := make(DeviceMap)
	return &devMap
}

// AddDevice adds a specified device.
func (devMap *DeviceMap) AddDevice(dev *Device) bool {
	if dev == nil {
		return false
	}

	devUdn := dev.UDN
	if len(devUdn) == 0 {
		return false
	}

	devType := dev.DeviceType
	if len(devType) == 0 {
		return false
	}

	devTypeMap, ok := (*devMap)[devType]
	if !ok {
		devTypeMap = make(map[string]*Device)
		(*devMap)[devType] = devTypeMap
	}

	devTypeMap[devUdn] = dev

	return true
}

// Size() returns all device count.
func (devMap *DeviceMap) Size() int {
	devCnt := 0

	for _, typeDevs := range *devMap {
		devCnt += len(typeDevs)
	}

	return devCnt
}

// GetAllDevices returns all devices.
func (devMap *DeviceMap) GetAllDevices() []*Device {
	devs := make([]*Device, 0)

	for _, typeDevs := range *devMap {
		for _, dev := range typeDevs {
			devs = append(devs, dev)
		}
	}

	return devs
}

// GetDevicesByType returns only devices of the specified device type.
func (devMap *DeviceMap) GetDevicesByType(deviceType string) []*Device {
	devs := make([]*Device, 0)

	typeDevs, ok := (*devMap)[deviceType]
	if !ok {
		return devs
	}

	for _, dev := range typeDevs {
		devs = append(devs, dev)
	}

	return devs
}

// FindDeviceByTypeAndUDN find a device of the specified device type and udn.
func (devMap *DeviceMap) FindDeviceByTypeAndUDN(deviceType string, udn string) (*Device, bool) {
	if len(deviceType) == 0 || len(udn) == 0 {
		return nil, false
	}

	devs, ok := (*devMap)[deviceType]
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
func (devMap *DeviceMap) HasDeviceByTypeAndUDN(deviceType string, udn string) bool {
	_, ok := devMap.FindDeviceByTypeAndUDN(deviceType, udn)
	return ok
}

// HasDevice adds a specified device.
func (devMap *DeviceMap) HasDevice(dev *Device) bool {
	if dev == nil {
		return false
	}
	return devMap.HasDeviceByTypeAndUDN(dev.DeviceType, dev.UDN)
}
