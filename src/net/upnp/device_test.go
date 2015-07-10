// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"fmt"
	"testing"

	gohttp "net/http"
)

const (
	errorTestDeviceInvalidURL = "invalid %s = '%s', expected : '%s'"
)

type sampleDevice struct {
	*Device
}

func NewSampleDevice() (*sampleDevice, error) {
	dev, err := NewDeviceFromDescription(binaryLightDeviceDescription)
	if err != nil {
		return nil, err
	}

	service, err := dev.GetServiceByType("urn:schemas-upnp-org:service:SwitchPower:1")
	if err != nil {
		return nil, err
	}

	err = service.LoadDescriptionString(switchPowerServiceDescription)
	if err != nil {
		return nil, err
	}

	sampleDev := &sampleDevice{Device: dev}

	return sampleDev, nil
}

func TestNullDevice(t *testing.T) {
	dev := NewDevice()

	err := dev.Start()
	if err != nil {
		t.Error(err)
	}

	res, err := gohttp.Get(fmt.Sprintf("http://localhost:%d/", dev.Port))
	if err != nil {
		t.Error(err)
	}

	if (dev.Port < DeviceDefaultPortBase) || (DeviceDefaultPortMax < dev.Port) {
		t.Errorf("got invalid port = [%d] : expected : [%d]~[%d]", dev.Port, DeviceDefaultPortBase, DeviceDefaultPortMax)
	}

	if res.StatusCode != gohttp.StatusInternalServerError {
		t.Errorf("got invalid test code = [%d] : expected : [%d]", res.StatusCode, gohttp.StatusInternalServerError)
	}

	err = dev.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestSampleDevice(t *testing.T) {
	dev, err := NewSampleDevice()

	if err != nil {
		t.Error(err)
	}

	// check service

	service, err := dev.GetServiceByType("urn:schemas-upnp-org:service:SwitchPower:1")
	if err != nil {
		t.Error(err)
	}

	service, err = dev.GetServiceById("urn:upnp-org:serviceId:SwitchPower.1")
	if err != nil {
		t.Error(err)
	}

	// check actions

	actionNames := []string{"SetTarget", "GetTarget", "GetStatus"}
	for _, name := range actionNames {
		_, err := service.GetActionByName(name)
		if err != nil {
			t.Error(err)
		}
	}

	// check argumengs

	action, err := service.GetActionByName("SetTarget")
	if err == nil {
		argNames := []string{"newTargetValue"}
		for _, name := range argNames {
			_, err := action.GetArgumentByName(name)
			if err != nil {
				t.Error(err)
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

	checkServiceURLs := func(dev *sampleDevice, serviceType string, urls []string) {
		service, err := dev.GetServiceByType(serviceType)
		if err != nil {
			t.Error(err)
		}

		expectURL := urls[0]
		if len(service.SCPDURL) <= 0 || service.SCPDURL != expectURL {
			t.Errorf(errorTestDeviceInvalidURL, "SCPDURL", service.SCPDURL, expectURL)
		}

		expectURL = urls[1]
		if len(service.ControlURL) <= 0 || service.ControlURL != expectURL {
			t.Errorf(errorTestDeviceInvalidURL, "ControlURL", service.ControlURL, expectURL)
		}

		expectURL = urls[2]
		if len(service.EventSubURL) <= 0 || service.EventSubURL != expectURL {
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
