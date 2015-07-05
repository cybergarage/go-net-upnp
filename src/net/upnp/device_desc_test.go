// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"fmt"
	"testing"
)

const testDeviceDesecription = xml.Header +
	"<root>" +
	"<device>" +
	"    <friendlyName>MediaServer</friendlyName>" +
	"    <serviceList>" +
	"        <service>" +
	"		<Optional/>" +
	"            <serviceType>urn:schemas-upnp-org:service:AVTransport:1</serviceType>" +
	"			<serviceId>AVTransport</serviceId>" +
	"        </service>" +
	"        <service>" +
	"            <serviceType>urn:schemas-upnp-org:service:ContentDirectory:1</serviceType>" +
	"			<serviceId>ContentDirectory</serviceId>" +
	"        </service>" +
	"        <service>" +
	"            <serviceType>urn:schemas-upnp-org:service:ConnectionManager:1</serviceType>" +
	"			<serviceId>ConnectionManager</serviceId>" +
	"        </service>" +
	"    </serviceList>" +
	"</device>" +
	"</root>"

func TestNewDeviceDescriptionRoot(t *testing.T) {
	NewDeviceDescriptionRoot()
}

func TestParseDeviceDescriptionRoot(t *testing.T) {
	_, err := NewDeviceDescriptionRootFromString(testDeviceDesecription)
	if err != nil {
		t.Error(err)
	}
}

func TestGenerateDeviceDescription(t *testing.T) {
	dev, err := NewDeviceFromDescription(testDeviceDesecription)
	if err != nil {
		t.Error(err)
	}

	devDesc, err := dev.DescriptionString()
	if err != nil {
		t.Error(err)
	}

	name := dev.FriendlyName
	if name != "MediaServer" {
		fmt.Printf("%s", devDesc)
	}
}
