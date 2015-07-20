// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package go-net-upnp provides UPnP control point and device frameworks to implement the control point and any devices.

go-net-upnp supports UPnP control functions. The control point can search UPnP devices in the local netowrk, get the device and service descriptions. and post actions in the service:

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

In addition to the control point functions, go-net-upnp supports UPnP device functions to implement any UPnP devices using Go.

To implement UPnP devices, prepare the UPnP device and service descriptions as the following:

	type SampleDevice struct {
		*upnp.Device
		...
	}

	func NewSampleDevice() (*SampleDevice, error) {
		dev, err := upnp.NewDeviceFromDescription(xxxxDeviceDescription)
		...
		service, err := dev.GetServiceByType("urn:schemas-upnp-org:service:xxxx:x")
		...
		err = service.LoadDescriptionBytes([]byte(xxxxServiceDescription))
		...
		sampleDev := &SampleDevice{
			Device: dev,
			...
		}
		return sampleDev, nil
	}

Next, implement the control actions in the service descriptions using upnp.ActionListener as the following:

	sampleDev, err := NewSampleDevice()
	...
	sampleDev.ActionListener = sampleDev
	...
	func (self *SampleDevice) ActionRequestReceived(action *upnp.Action) upnp.Error {
		switch action.Name {
		case SetTarget:
			xxxx, err := action.GetArgumentString(xxxx)
			...
			err := action.SetArgumentBool(...., ....)
			...
			return nil
		case xxxx
			...
		}
		return upnp.NewErrorFromCode(upnp.ErrorOptionalActionNotImplemented)
	}
*/
package upnp
