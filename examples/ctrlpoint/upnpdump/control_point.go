// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnpdump

import (
	"fmt"
	"os"

	"github.com/cybergarage/go-net-upnp/net/upnp"
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
	"github.com/cybergarage/go-net-upnp/net/upnp/util"
)

// A ControlPoint represents a ControlPoint.
type ControlPoint struct {
	*upnp.ControlPoint
}

// NewControlPoint returns a new Client.
func NewControlPoint() *ControlPoint {
	cp := &ControlPoint{}

	cp.ControlPoint = upnp.NewControlPoint()
	cp.ControlPoint.Listener = cp

	return cp
}

func printMessage(msg string) {
	os.Stdout.WriteString(fmt.Sprintf("%s\n", msg))
}

func GetFromToMessageFromSSDPPacket(req *ssdp.Packet) string {
	fromAddr := req.From.String()
	toAddr := ""
	ifAddr, err := util.GetInterfaceAddress(req.Interface)
	if err == nil {
		toAddr = ifAddr
	}

	return fmt.Sprintf("(%s -> %s)", fromAddr, toAddr)
}

func (self *ControlPoint) DeviceNotifyReceived(req *ssdp.Request) {
	usn, _ := req.GetUSN()
	printMessage(fmt.Sprintf("notiry req : %s %s", usn, GetFromToMessageFromSSDPPacket(req.Packet)))
}

func (self *ControlPoint) DeviceSearchReceived(req *ssdp.Request) {
	st, _ := req.GetST()
	printMessage(fmt.Sprintf("search req : %s %s", st, GetFromToMessageFromSSDPPacket(req.Packet)))
}

func (self *ControlPoint) DeviceResponseReceived(res *ssdp.Response) {
	url, _ := res.GetLocation()
	printMessage(fmt.Sprintf("search res : %s %s", url, GetFromToMessageFromSSDPPacket(res.Packet)))
}
