// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"fmt"
	gohttp "net/http"
	"testing"
)

func TestNewDevice(t *testing.T) {
	dev := NewDevice()

	err := dev.Start()
	if err != nil {
		t.Error(err)
	}

	res, err := gohttp.Get(fmt.Sprintf("http://localhost:%d/", dev.Port))
	if err != nil {
		t.Error(err)
	}

	if (dev.Port < DEVICE_DEFAULT_PORT_BASE) || (DEVICE_DEFAULT_PORT_MAX < dev.Port) {
		t.Errorf("got invalid port = [%d] : expected : [%d]~[%d]", dev.Port, DEVICE_DEFAULT_PORT_BASE, DEVICE_DEFAULT_PORT_MAX)
	}

	if res.StatusCode != gohttp.StatusInternalServerError {
		t.Errorf("got invalid test code = [%d] : expected : [%d]", res.StatusCode, gohttp.StatusInternalServerError)
	}

	err = dev.Stop()
	if err != nil {
		t.Error(err)
	}
}

const testMediaServerDeviceDesc = xml.Header +
	"<device>" +
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
	"</device>"

func TestDeviceLoadDescription(t *testing.T) {
	dev := NewDevice()

	err := dev.LoadDescriptionString(testMediaServerDeviceDesc)
	if err != nil {
		t.Error(err)
	}

	for n, service := range dev.ServiceList.Services {
		var expectedServiceType string
		var expectedServiceId string
		switch n {
		case 0:
			expectedServiceType = "urn:schemas-upnp-org:service:AVTransport:1"
			expectedServiceId = "AVTransport"
		case 1:
			expectedServiceType = "urn:schemas-upnp-org:service:ContentDirectory:1"
			expectedServiceId = "ContentDirectory"
		case 2:
			expectedServiceType = "urn:schemas-upnp-org:service:ConnectionManager:1"
			expectedServiceId = "ConnectionManager"
		}
		if service.ServiceType != expectedServiceType {
			t.Errorf("serviceType = %s, expected %s", service.ServiceType, expectedServiceType)
		}
		if service.ServiceId != expectedServiceId {
			t.Errorf("serviceId = %s, expected %s", service.ServiceId, expectedServiceId)
		}
	}
}
