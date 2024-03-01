// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"testing"

	"github.com/cybergarage/go-net-upnp/net/upnp/control"
)

const (
	errorTestDeviceInvalidURL           = "invalid url %s = '%s', expected : '%s'"
	errorTestDeviceInvalidStatusCode    = "invalid status code (%s) = [%d] : expected : [%d]"
	errorTestDeviceInvalidPortRange     = "invalid port range = [%d] : expected : [%d]~[%d]"
	errorTestDeviceInvalidParentObject  = "invalid parent object %p = '%p', expected : '%p'"
	errorTestDeviceInvalidArgumentValue = "invalid argument value %s = '%s', expected : '%s'"
	errorTestDeviceInvalidArgumentDir   = "invalid argument direction %s = %d, expected : %d"
)

const (
	SetTarget      = "SetTarget"
	GetTarget      = "GetTarget"
	GetStatus      = "GetStatus"
	NewTargetValue = "newTargetValue"
	RetTargetValue = "RetTargetValue"
)

type TestDevice struct {
	*Device
	Target string
}

func NewTestDevice() (*TestDevice, error) {
	dev, err := NewDeviceFromDescription(binaryLightDeviceDescription)
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

	testDev := &TestDevice{Device: dev}
	testDev.ActionListener = testDev

	return testDev, nil
}

func (dev *TestDevice) GetSwitchPowerService() (*Service, error) {
	return dev.GetServiceByType("urn:schemas-upnp-org:service:SwitchPower:1")
}

func (dev *TestDevice) GetSwitchPowerSetTargetAction() (*Action, error) {
	service, err := dev.GetSwitchPowerService()
	if err != nil {
		return nil, err
	}
	return service.GetActionByName(SetTarget)
}

func (dev *TestDevice) GetSwitchPowerGetTargetAction() (*Action, error) {
	service, err := dev.GetSwitchPowerService()
	if err != nil {
		return nil, err
	}
	return service.GetActionByName(GetTarget)
}

func (dev *TestDevice) GetOptionalAction() (*Action, error) {
	service, err := dev.GetSwitchPowerService()
	if err != nil {
		return nil, err
	}
	return service.GetActionByName(GetStatus)
}

func (dev *TestDevice) ActionRequestReceived(action *Action) Error {
	switch action.Name {
	case SetTarget:
		arg, err := action.GetArgumentByName(NewTargetValue)
		if err == nil {
			dev.Target = arg.Value
		}
		return nil
	case GetTarget:
		arg, err := action.GetArgumentByName(RetTargetValue)
		if err == nil {
			arg.Value = dev.Target
		}
		return nil
	}

	return control.NewUPnPErrorFromCode(control.ErrorOptionalActionNotImplemented)
}

func TestTestDeviceDescription(t *testing.T) {
	dev, err := NewTestDevice()

	if err != nil {
		t.Error(err)
	}

	// check service

	service, err := dev.GetServiceByType("urn:schemas-upnp-org:service:SwitchPower:1")
	if err != nil {
		t.Error(err)
	}

	if service.ParentDevice != dev.Device {
		t.Errorf(errorTestDeviceInvalidParentObject, service, service.ParentDevice, dev.Device)
	}

	service, err = dev.GetServiceByID("urn:upnp-org:serviceId:SwitchPower.1")
	if err != nil {
		t.Error(err)
	}

	if service.ParentDevice != dev.Device {
		t.Errorf(errorTestDeviceInvalidParentObject, service, service.ParentDevice, dev.Device)
	}

	// check actions

	actionNames := []string{"SetTarget", "GetTarget", "GetStatus"}
	for _, name := range actionNames {
		action, err := service.GetActionByName(name)
		if err != nil {
			t.Error(err)
		}
		if action.ParentService != service {
			t.Errorf(errorTestDeviceInvalidParentObject, action, action.ParentService, service)
		}
	}

	// check argumengs (SetTarget)

	action, err := service.GetActionByName("SetTarget")
	if err == nil {
		argNames := []string{"newTargetValue"}
		argDirs := []int{InDirection}
		for n, name := range argNames {
			arg, err := action.GetArgumentByName(name)
			if err != nil {
				t.Error(err)
			}

			argDir := arg.GetDirection()
			if argDir != argDirs[n] {
				t.Errorf(errorTestDeviceInvalidArgumentDir, name, argDir, argDirs[n])
			}

			// check parent service

			if arg.ParentAction != action {
				t.Errorf(errorTestDeviceInvalidParentObject, arg, arg.ParentAction, action)
			}

			// check setter and getter

			value := fmt.Sprintf("%d", rand.Int())
			err = arg.SetString(value)
			if err != nil {
				t.Error(err)
			}
			argValue, err := arg.GetString()
			if err != nil {
				t.Error(err)
			}
			if value != argValue {
				t.Errorf(errorTestDeviceInvalidArgumentValue, name, argValue, value)
			}
		}
	} else {
		t.Error(err)
	}

	// start device

	err = dev.Start()
	if err != nil {
		t.Error(err)
	}

	// check service

	checkServiceURLs := func(dev *TestDevice, serviceType string, urls []string) {
		service, err := dev.GetServiceByType(serviceType)
		if err != nil {
			t.Error(err)
		}

		expectURL := urls[0]
		if len(service.SCPDURL) == 0 || service.SCPDURL != expectURL {
			t.Errorf(errorTestDeviceInvalidURL, "SCPDURL", service.SCPDURL, expectURL)
		}

		expectURL = urls[1]
		if len(service.ControlURL) == 0 || service.ControlURL != expectURL {
			t.Errorf(errorTestDeviceInvalidURL, "ControlURL", service.ControlURL, expectURL)
		}

		expectURL = urls[2]
		if len(service.EventSubURL) == 0 || service.EventSubURL != expectURL {
			t.Errorf(errorTestDeviceInvalidURL, "EventSubURL", service.EventSubURL, expectURL)
		}
	}

	urls := []string{
		"/service/scpd/SwitchPower.xml",
		"/service/control/SwitchPower",
		"/service/event/SwitchPower"}
	checkServiceURLs(dev, "urn:schemas-upnp-org:service:SwitchPower:1", urls)

	// stop device
	err = dev.Stop()
	if err != nil {
		t.Error(err)
	}
}

const binaryLightDeviceDescription = xml.Header +
	"<root>" +
	"  <device>" +
	"    <deviceType>urn:schemas-upnp-org:device:BinaryLight:1</deviceType>" +
	"    <serviceList>" +
	"      <service>" +
	"        <serviceType>urn:schemas-upnp-org:service:SwitchPower:1</serviceType>" +
	"        <serviceId>urn:upnp-org:serviceId:SwitchPower.1</serviceId>" +
	"      </service>" +
	"    </serviceList>" +
	"  </device>" +
	"</root>"

const switchPowerServiceDescription = xml.Header +
	"<scpd>" +
	"  <serviceStateTable>" +
	"    <stateVariable>" +
	"      <name>Target</name>" +
	"      <sendEventsAttribute>no</sendEventsAttribute> " +
	"      <dataType>boolean</dataType>" +
	"      <defaultValue>0</defaultValue>" +
	"    </stateVariable>" +
	"    <stateVariable>" +
	"      <name>Status</name>" +
	"      <dataType>boolean</dataType>" +
	"      <defaultValue>0</defaultValue>" +
	"    </stateVariable>" +
	"  </serviceStateTable>" +
	"  <actionList>" +
	"    <action>" +
	"    <name>SetTarget</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>newTargetValue</name>" +
	"          <direction>in</direction>" +
	"          <relatedStateVariable>Target</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action>" +
	"    <name>GetTarget</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>RetTargetValue</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>Target</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"    <action>" +
	"    <name>GetStatus</name>" +
	"      <argumentList>" +
	"        <argument>" +
	"          <name>ResultStatus</name>" +
	"          <direction>out</direction>" +
	"          <relatedStateVariable>Status</relatedStateVariable>" +
	"        </argument>" +
	"      </argumentList>" +
	"    </action>" +
	"  </actionList>" +
	"</scpd>"
