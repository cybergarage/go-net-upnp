// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnpctrl

import (
	"fmt"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-net-upnp/net/upnp"
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
	"github.com/cybergarage/go-net-upnp/net/upnp/util"
)

// A ControlPoint represents a ControlPoint.
type ControlPoint struct {
	*upnp.ControlPoint
	*ControlPointActionManager
}

// NewControlPoint returns a new Client.
func NewControlPoint() *ControlPoint {
	cp := &ControlPoint{}

	cp.ControlPoint = upnp.NewControlPoint()
	cp.ControlPointActionManager = NewControlPointActionManager()
	cp.ControlPoint.Listener = cp

	return cp
}

func printMessage(msg string) {
	log.Tracef(msg)
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
	//os.Stdout.WriteString(fmt.Sprintf("%s\n", req.String()))
}

func (self *ControlPoint) DeviceSearchReceived(req *ssdp.Request) {
	st, _ := req.GetST()
	printMessage(fmt.Sprintf("search req : %s %s", st, GetFromToMessageFromSSDPPacket(req.Packet)))
	//os.Stdout.WriteString(fmt.Sprintf("%s\n", req.String()))
}

func (self *ControlPoint) DeviceResponseReceived(res *ssdp.Response) {
	url, _ := res.GetLocation()
	printMessage(fmt.Sprintf("search res : %s %s", url, GetFromToMessageFromSSDPPacket(res.Packet)))
	//os.Stdout.WriteString(fmt.Sprintf("%s\n", res.String()))
}

func (self *ControlPoint) DoAction(key int) bool {
	return self.ControlPointActionManager.DoAction(self, key)
}
