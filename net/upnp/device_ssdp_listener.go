// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
)

func (self *Device) DeviceNotifyReceived(ssdpReq *ssdp.Request) {
	if self.SSDPListener != nil {
		self.SSDPListener.DeviceNotifyReceived(ssdpReq)
	}
}

func (self *Device) postResponseMessge(ssdpReq *ssdp.Request) error {
	fromAddr := ssdpReq.From.IP.String()
	fromPort := ssdpReq.From.Port

	ifAddr, err := self.selectAvailableInterfaceForAddr(fromAddr)
	if err != nil {
		return err
	}

	locationURL, err := self.createLocationURLForAddress(ifAddr)
	if err != nil {
		return err
	}

	ssdpRes := ssdp.NewResponse()
	ssdpRes.SetLocation(locationURL.String())

	sock := ssdp.NewUnicastSocket()
	_, err = sock.WriteResponse(fromAddr, fromPort, ssdpRes)

	return err
}

func (self *Device) handleDiscoverRequest(ssdpReq *ssdp.Request) {
	if ssdpReq.IsRootDevice() {
		self.postResponseMessge(ssdpReq)
		return
	}

	st, err := ssdpReq.GetST()
	if err != nil {
		return
	}

	if self.HasDeviceType(st) || self.HasServiceType(st) {
		self.postResponseMessge(ssdpReq)
		return
	}
}

func (self *Device) DeviceSearchReceived(ssdpReq *ssdp.Request) {
	if ssdpReq.IsDiscover() {
		self.handleDiscoverRequest(ssdpReq)
	}

	if self.SSDPListener != nil {
		self.SSDPListener.DeviceSearchReceived(ssdpReq)
	}
}
