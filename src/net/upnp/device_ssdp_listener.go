// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"net/upnp/ssdp"
)

func (self *Device) DeviceNotifyReceived(ssdpReq *ssdp.Request) {
	if self.Listener != nil {
		self.Listener.DeviceNotifyReceived(ssdpReq)
	}
}

func (self *Device) postResponseMessge(ssdpReq *ssdp.Request) {
	from := ssdpReq.From

	ssdpRes := ssdp.NewResponse()

}

func (self *Device) handleDiscoverRequest(ssdpReq *ssdp.Request) {
	if ssdpReq.IsRootDevice() {
		self.postResponseMessge(ssdpReq)
		return
	}
}

func (self *Device) DeviceSearchReceived(ssdpReq *ssdp.Request) {
	if ssdpReq.IsDiscover() {
		self.handleDiscoverRequest(ssdpReq)
	}

	if self.Listener != nil {
		self.Listener.DeviceSearchReceived(ssdpReq)
	}
}
