// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"net/upnp"
	"net/upnp/ssdp"
	"net/upnp/util"
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
	os.Stdout.WriteString(fmt.Sprintf("%s\n", msg))
}

func GetFromToMessageFromSSDPRequest(req *ssdp.Request) string {
	fromAddr := req.From.String()
	toAddr := ""
	ifAddr, err := util.GetInterfaceAddress(&req.Interface)
	if err == nil {
		toAddr = ifAddr
	}

	return fmt.Sprintf("(%s -> %s)", fromAddr, toAddr)
}

func (self *ControlPoint) DeviceNotifyReceived(req *ssdp.Request) {
	os.Stdout.WriteString(fmt.Sprintf("%s\n", req.String()))
}

func (self *ControlPoint) DeviceSearchReceived(req *ssdp.Request) {
	st, _ := req.GetST()
	msg := fmt.Sprintf("search : %s %s", st, GetFromToMessageFromSSDPRequest(req))
	printMessage(msg)
	os.Stdout.WriteString(fmt.Sprintf("%s\n", req.String()))
}

func (self *ControlPoint) DeviceResponseReceived(res *ssdp.Response) {
	os.Stdout.WriteString(fmt.Sprintf("%s\n", res.String()))
}

func (self *ControlPoint) DoAction(key int) bool {
	return self.ControlPointActionManager.DoAction(self, key)
}
