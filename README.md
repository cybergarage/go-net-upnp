# go-net-upnp

go-net-upnp is a open source framework for Go and UPnP™ developers.

UPnP™ is a standard protocol for IoT, the protocols consist of other standard protocols, such as GENA, SSDP, SOAP, HTTPU and HTTP. Therefore you have to understand and implement these protocols to create UPnP™ applications.

go-net-upnp hansles these protocols automatically to support to create UPnP devices and control points quickly.

## Installation

The project is released on [GitHub](https://github.com/cybergarage/go-net-upnp). To use go-net-upnp in your projct, run `go get` as the following:

```
go get -u github.com/cybergarage/go-net-upnp/net/upnp
```
## Overview

go-net-upnp provides UPnP control point and device frameworks to implement the control point and any devices.

### Control Point Implementation

go-net-upnp supports UPnP control functions. The control point can search UPnP devices in the local netowrk, get the device and service descriptions. and post actions in the service:

```
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

To know how to implement UPnP control point or devices in more deital using go-net-upnp, please check the sample implementations in the `example` directory and the `godoc` documentation :-)



