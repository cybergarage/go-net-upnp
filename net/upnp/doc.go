// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package upnp privides UPnP control point and device implementations.

The control point can search UPnP devices in the local netowrk, get the device and service descriptions. and post actions in the service:

	cp := upnp.NewControlPoint()
	err := cp.Start()
	...
	defer cp.Stop()
	...
	err = cp.SearchRootDevice()
	...
	for n, dev := range cp.GetRootDevices() {
		...
	}

The control point can post actions in the service, and get the action response:

	service, err := dev.GetServiceByType("xxxx")
	...
	action, err := service.GetActionByName("xxxx")
	...
	action.SetArgumentString("xxxx", "xxxx")
	err = action.Post()
	...
	resArg = action.GetArgumentString("xxxx")
*/
package upnp