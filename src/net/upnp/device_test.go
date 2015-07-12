// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
)

const (
	errorTestDeviceInvalidURL          = "invalid url %s = '%s', expected : '%s'"
	errorTestDeviceInvalidStatusCode   = "invalid status code (%s) = [%d] : expected : [%d]"
	errorTestDeviceInvalidPortRange    = "invalid port range = [%d] : expected : [%d]~[%d]"
	errorTestDeviceInvalidParentObject = "invalid parent object %p = '%p', expected : '%p'"
	errorTestDeviceInvalidArgument     = "invalid argument %s = '%s', expected : '%s'"
)

func TestNullDevice(t *testing.T) {
	dev := NewDevice()

	err := dev.Start()
	if err != nil {
		t.Error(err)
	}

	if (dev.Port < DeviceDefaultPortBase) || (DeviceDefaultPortMax < dev.Port) {
		t.Errorf(errorTestDeviceInvalidPortRange, dev.Port, DeviceDefaultPortBase, DeviceDefaultPortMax)
	}

	url := fmt.Sprintf("http://localhost:%d/", dev.Port)
	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf(errorTestDeviceInvalidStatusCode, url, res.StatusCode, http.StatusInternalServerError)
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

	if service.ParentDevice != dev.Device {
		t.Errorf(errorTestDeviceInvalidParentObject, service, service.ParentDevice, dev.Device)
	}

	service, err = dev.GetServiceById("urn:upnp-org:serviceId:SwitchPower.1")
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

	// check argumengs

	action, err := service.GetActionByName("SetTarget")
	if err == nil {
		argNames := []string{"newTargetValue"}
		for _, name := range argNames {
			arg, err := action.GetArgumentByName(name)
			if err != nil {
				t.Error(err)
			}
			if arg.ParentAction != action {
				t.Errorf(errorTestDeviceInvalidParentObject, arg, arg.ParentAction, action)
			}
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
				t.Errorf(errorTestDeviceInvalidArgument, name, argValue, value)
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

	// check device
	if (dev.Port < DeviceDefaultPortBase) || (DeviceDefaultPortMax < dev.Port) {
		t.Errorf(errorTestDeviceInvalidPortRange, dev.Port, DeviceDefaultPortBase, DeviceDefaultPortMax)
	}

	devDescURL := fmt.Sprintf("http://localhost:%d%s", dev.Port, dev.DescriptionURL)
	res, err := http.Get(devDescURL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf(errorTestDeviceInvalidStatusCode, devDescURL, res.StatusCode, http.StatusOK)
	}

	// stop device
	err = dev.Stop()
	if err != nil {
		t.Error(err)
	}
}
