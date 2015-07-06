// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"errors"
	"fmt"
	"testing"
)

const (
	errorDeviceDesecriptionInvalidField = "%s = '%s' : expected %s"
)

func newDeviceDesecriptionInvalidFieldError(name, value, expected string) error {
	return errors.New(fmt.Sprintf(errorDeviceDesecriptionInvalidField, name, value, expected))
}

const testDeviceDesecription = xml.Header +
	"<root>" +
	"<device>" +
	"	 <deviceType>MediaServer:1</deviceType>" +
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
	var tagNames, tagValues, expectValues []string

	// Check loading description

	tagNames = []string{
		"friendlyName",
		"deviceType",
		"serviceType[0]",
	}

	expectValues = []string{
		"MediaServer",
		"MediaServer:1",
		"urn:schemas-upnp-org:service:AVTransport:1",
		"urn:schemas-upnp-org:service:ContentDirectory:1",
		"urn:schemas-upnp-org:service:ConnectionManager:1",
		"AVTransport",
		"ContentDirectory",
		"ConnectionManager",
	}

	dev, err := NewDeviceFromDescription(testDeviceDesecription)
	if err != nil {
		t.Error(err)
	}

	tagValues = []string{
		dev.FriendlyName,
		dev.DeviceType,
		dev.ServiceList.Services[0].ServiceType,
		dev.ServiceList.Services[1].ServiceType,
		dev.ServiceList.Services[2].ServiceType,
		dev.ServiceList.Services[0].ServiceId,
		dev.ServiceList.Services[1].ServiceId,
		dev.ServiceList.Services[2].ServiceId,
	}

	for n, expectValue := range expectValues {
		if expectValue != tagValues[n] {
			t.Error(newDeviceDesecriptionInvalidFieldError(tagNames[n], tagValues[n], expectValue))
		}
	}

	// Check output description

	outputDevDesc, err := dev.DescriptionString()
	if err != nil {
		t.Error(err)
	}

	outputDev, err := NewDeviceFromDescription(outputDevDesc)
	if err != nil {
		t.Error(err)
	}

	tagValues = []string{
		outputDev.FriendlyName,
		outputDev.DeviceType,
		outputDev.ServiceList.Services[0].ServiceType,
		outputDev.ServiceList.Services[1].ServiceType,
		outputDev.ServiceList.Services[2].ServiceType,
		outputDev.ServiceList.Services[0].ServiceId,
		outputDev.ServiceList.Services[1].ServiceId,
		outputDev.ServiceList.Services[2].ServiceId,
	}

	for n, expectValue := range expectValues {
		if expectValue != tagValues[n] {
			t.Error(newDeviceDesecriptionInvalidFieldError(tagNames[n], tagValues[n], expectValue))
		}
	}
}
