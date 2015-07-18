// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/cybergarage/go-net-upnp/net/upnp"
)

const (
	InternetGatewayDeviceType = "urn:schemas-upnp-org:device:InternetGatewayDevice:1"
	WANConnectionDevice       = "urn:schemas-upnp-org:device:WANConnectionDevice:1"

	WANIPConnectionServiceType          = "urn:schemas-upnp-org:service:WANIPConnection:1"
	WANCommonInterfaceConfigServiceType = "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1"

	GetExternalIPAddressAction = "GetExternalIPAddress"
	NewExternalIPAddress       = "NewExternalIPAddress"

	PortMappingNumberOfEntriesAction = "PortMappingNumberOfEntries"

	GetGenericPortMappingEntryAction = "GetGenericPortMappingEntry"
	NewPortMappingIndex              = "NewPortMappingIndex"

	GetTotalBytesReceivedAction = "GetTotalBytesReceived"
	NewTotalBytesReceived       = "NewTotalBytesReceived"

	GetTotalBytesSentAction = "GetTotalBytesSent"
	NewTotalBytesSent       = "NewTotalBytesSent"
)

type GatewayDevice struct {
	*upnp.Device
}

func NewGatewayDevice(dev *upnp.Device) *GatewayDevice {
	gwDev := &GatewayDevice{Device: dev}
	return gwDev
}

func (self *GatewayDevice) GetWANIPConnectionServiceAction(name string) (*upnp.Action, error) {
	wanConDev, err := self.GetEmbeddedDeviceByType(WANConnectionDevice)
	if err != nil {
		return nil, err
	}

	service, err := wanConDev.GetServiceByType(WANIPConnectionServiceType)
	if err != nil {
		return nil, err
	}

	return service.GetActionByName(name)
}

func (self *GatewayDevice) GetExternalIPAddress() (string, error) {
	action, err := self.GetWANIPConnectionServiceAction(GetExternalIPAddressAction)
	if err != nil {
		return "", err
	}

	err = action.Post()
	if err != nil {
		return "", err
	}

	return action.GetArgumentString(NewExternalIPAddress)
}
