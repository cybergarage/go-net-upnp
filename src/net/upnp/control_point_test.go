// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	errorInvalidControlPoint        = "ControlPoint is invalid"
	errorControlPointDeviceNotFound = "ControlPoint can't the device (%s)"
)

func TestNewControlPoint(t *testing.T) {
	cp := NewControlPoint()

	if len(cp.GetRootDevices()) != 0 {
		t.Error(errors.New(errorInvalidControlPoint))
	}

	err := cp.Start()
	if err != nil {
		t.Error(err)
	}

	err = cp.SearchRootDevice()
	if err != nil {
		t.Error(err)
	}

	err = cp.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestControlPointSearchDevice(t *testing.T) {
	// start device

	dev, err := NewSampleDevice()
	if err != nil {
		t.Error(err)
	}

	err = dev.Start()
	if err != nil {
		t.Error(err)
	}
	defer dev.Stop()

	devUDN := dev.UDN

	// start control point

	cp := NewControlPoint()

	if len(cp.GetRootDevices()) != 0 {
		t.Error(errors.New(errorInvalidControlPoint))
	}

	err = cp.Start()
	if err != nil {
		t.Error(err)
	}

	// find device

	err = cp.SearchRootDevice()
	if err != nil {
		t.Error(err)
	}

	var foundDev *Device
	loopCnt := 10
	for n := 0; n < loopCnt; n++ {
		waitMillSec := time.Duration(cp.SearchMX * 1000 / loopCnt)
		time.Sleep(waitMillSec * time.Millisecond)
		var ok bool
		foundDev, ok = cp.FindDeviceByUDN(devUDN)
		if ok {
			break
		}
	}

	if foundDev == nil {
		t.Errorf(errorControlPointDeviceNotFound, devUDN)
	}

	// check service

	devService, _ := dev.GetSwitchPowerService()

	foundService, err := foundDev.GetServiceByType(devService.ServiceType)
	if err != nil {
		t.Error(err)
	}

	// post action (set)

	postValue := fmt.Sprintf("target%d", rand.Int())

	devSetAction, _ := dev.GetSwitchPowerSetTargetAction()

	foundSetAction, err := foundService.GetActionByName(devSetAction.Name)
	if err != nil {
		t.Error(err)
	}

	foundSetActionArg := foundSetAction.ArgumentList.Arguments[0]
	foundSetActionArg.Value = postValue

	// post action (get)

	devGetAction, _ := dev.GetSwitchPowerGetTargetAction()

	_, err = foundService.GetActionByName(devGetAction.Name)
	if err != nil {
		t.Error(err)
	}

	// stop control point

	err = cp.Stop()
	if err != nil {
		t.Error(err)
	}
}
