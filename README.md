# go-net-upnp

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-net-upnp)
[![Go](https://github.com/cybergarage/go-net-upnp/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/go-net-upnp/actions/workflows/make.yml)
 [![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-net-upnp.svg)](https://pkg.go.dev/github.com/cybergarage/go-net-upnp)
 [![Go Report Card](https://img.shields.io/badge/go%20report-A%2B-brightgreen)](https://goreportcard.com/report/github.com/cybergarage/go-net-upnp) 
[![codecov](https://codecov.io/gh/cybergarage/go-net-upnp/graph/badge.svg?token=SS6L6DRNHF)](https://codecov.io/gh/cybergarage/go-net-upnp)

go-net-upnp is a new open source framework for Go and UPnP™ developers.

UPnP™ is a standard protocol for IoT, it consist of other standard protocols, such as GENA, SSDP, SOAP, HTTPU and HTTP. Therefore UPnP developers have to understand and implement these protocols to create UPnP™ applications.

go-net-upnp manages these protocols automatically to support to create any UPnP devices and control points quickly.

## Installation

The project is released on [GitHub](https://github.com/cybergarage/go-net-upnp). To use go-net-upnp in your projct, run `go get` as the following:

```
go get -u github.com/cybergarage/go-net-upnp
```
## Overview

go-net-upnp provides UPnP control point and device frameworks to implement the control point and any devices.

### Control Point Implementation

go-net-upnp supports UPnP control functions. The control point can search UPnP devices in the local network, get the device and service descriptions. and post actions in the service:

```
	import (
		"github.com/cybergarage/go-net-upnp/net/upnp"
	)
	
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
```

The control point can post actions in the service, and get the action response:

```
	service, err := dev.GetServiceByType("xxxx")
	...
	action, err := service.GetActionByName("xxxx")
	...
	action.SetArgumentString("xxxx", "xxxx")
	err = action.Post()
	...
	resArg = action.GetArgumentString("xxxx")
```

## Device Implementation

In addition to the control point functions, go-net-upnp supports UPnP device functions to implement any UPnP devices using Go.

To implement UPnP devices, prepare the UPnP device and service descriptions as the following:

```
	import (
		"github.com/cybergarage/go-net-upnp/net/upnp"
	)
	
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
```

Next, implement the control actions in the service descriptions using upnp.ActionListener as the following:

```
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
```

## Next Steps

To know how to implement UPnP control point or devices in more deital using go-net-upnp, please check the sample implementations in the [example](https://github.com/cybergarage/go-net-upnp/tree/master/examples) directory and the [godoc](https://pkg.go.dev/github.com/cybergarage/go-net-upnp) documentation :-)

## Release plan

In the first release version, 0.8, `go-net-upnp` supports major UPnP control point and device functions. However I will implements other functions by the final release version, 1.0, as the following.

- 0.9 : Support the event subscription function of UPnP and deprecated functions from UPnP v1.1 such as query function.
- 1.0 : Support UPnP v2.0 specifications more correctly.



