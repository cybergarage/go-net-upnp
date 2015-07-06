// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"net/upnp"
)

type LightDevice struct {
	*upnp.Device
}

func NewLightDevice() (*LightDevice, error) {
	dev, err := upnp.NewDeviceFromDescription(binaryLightDeviceDescription)
	if err != nil {
		return nil, err
	}

	lightDev := &LightDevice{Device: dev}

	return lightDev, nil
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
