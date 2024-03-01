// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
)

func (dev *Device) DeviceNotifyReceived(ssdpReq *ssdp.Request) {
	if dev.SSDPListener != nil {
		dev.SSDPListener.DeviceNotifyReceived(ssdpReq)
	}
}

func (dev *Device) postResponseMessge(ssdpReq *ssdp.Request) error {
	fromAddr := ssdpReq.From.IP.String()
	fromPort := ssdpReq.From.Port

	ifAddr, err := dev.selectAvailableInterfaceForAddr(fromAddr)
	if err != nil {
		return err
	}

	locationURL, err := dev.createLocationURLForAddress(ifAddr)
	if err != nil {
		return err
	}

	ssdpRes := ssdp.NewResponse()
	ssdpRes.SetLocation(locationURL.String())

	sock := ssdp.NewUnicastSocket()
	_, err = sock.WriteResponse(fromAddr, fromPort, ssdpRes)

	return err
}

func (dev *Device) handleDiscoverRequest(ssdpReq *ssdp.Request) {
	if ssdpReq.IsRootDevice() {
		dev.postResponseMessge(ssdpReq)
		return
	}

	st, err := ssdpReq.GetST()
	if err != nil {
		return
	}

	if dev.HasDeviceType(st) || dev.HasServiceType(st) {
		dev.postResponseMessge(ssdpReq)
		return
	}
}

func (dev *Device) DeviceSearchReceived(ssdpReq *ssdp.Request) {
	if ssdpReq.IsDiscover() {
		dev.handleDiscoverRequest(ssdpReq)
	}

	if dev.SSDPListener != nil {
		dev.SSDPListener.DeviceSearchReceived(ssdpReq)
	}
}
