// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
upntctrl browses UPnP devices in the local network, and post the actions.

        NAME
        upntctrl

        SYNOPSIS
        upnpdump [OPTIONS]

        DESCRIPTION
        upnpdump is a utility to dump SSDP messages.


        OPTIONS
        -v [0 | 1] : Enable verbose output.

        EXIT STATUS
          Return EXIT_SUCCESS or EXIT_FAILURE

        EXAMPLES
          The following is how to enable the verbose output
            upntctrl -v 1
*/
package lightdev

import (
	"github.com/cybergarage/go-net-upnp/net/upnp"
)

const (
	SetTarget      = "SetTarget"
	NewTargetValue = "newTargetValue"
	GetTarget      = "GetTarget"
	RetTargetValue = "RetTargetValue"
	GetStatus      = "GetStatus"
	ResultStatus   = "ResultStatus"

	DefaultTarget = "Living"
	DefaultStatus = true
)

type LightDevice struct {
	*upnp.Device
	Target string
	Status bool
}

func NewLightDevice() (*LightDevice, error) {
	dev, err := upnp.NewDeviceFromDescription(binaryLightDeviceDescription)
	if err != nil {
		return nil, err
	}

	service, err := dev.GetServiceByType("urn:schemas-upnp-org:service:SwitchPower:1")
	if err != nil {
		return nil, err
	}

	err = service.LoadDescriptionBytes([]byte(switchPowerServiceDescription))
	if err != nil {
		return nil, err
	}

	lightDev := &LightDevice{
		Device: dev,
		Target: DefaultTarget,
		Status: DefaultStatus,
	}
	lightDev.ActionListener = lightDev

	return lightDev, nil
}

func (self *LightDevice) ActionRequestReceived(action *upnp.Action) upnp.Error {
	switch action.Name {
	case SetTarget:
		target, err := action.GetArgumentString(NewTargetValue)
		if err == nil {
			self.Target = target
		} else {
			return upnp.NewErrorFromCode(upnp.ErrorInvalidArgs)
		}
		return nil
	case GetTarget:
		err := action.SetArgumentString(RetTargetValue, self.Target)
		if err != nil {
			return upnp.NewErrorFromCode(upnp.ErrorInvalidArgs)
		}
		return nil
	case GetStatus:
		err := action.SetArgumentBool(ResultStatus, self.Status)
		if err != nil {
			return upnp.NewErrorFromCode(upnp.ErrorInvalidArgs)
		}
		return nil
	}

	return upnp.NewErrorFromCode(upnp.ErrorOptionalActionNotImplemented)
}
