// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"errors"
	"fmt"
	"testing"

	gohttp "net/http"
)

type sampleDevice struct {
	*Device
}

func NewSampleDevice() (*sampleDevice, error) {
	dev, err := NewDeviceFromDescription(binaryLightDeviceDescription)
	if err != nil {
		return nil, err
	}

	if len(dev.ServiceList.Services) != 1 {
		return nil, errors.New("service is not found !!")
	}

	switchSrv := dev.ServiceList.Services[0]

	err = switchSrv.LoadDescriptionString(switchPowerServiceDescription)
	if err != nil {
		return nil, err
	}

	sampleDev := &sampleDevice{Device: dev}

	return sampleDev, nil
}

func TestNewNullDevice(t *testing.T) {
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

	err = dev.Start()
	if err != nil {
		t.Error(err)
	}

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
