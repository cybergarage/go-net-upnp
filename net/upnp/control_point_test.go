// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cybergarage/go-net-upnp/net/upnp/control"
)

const (
	errorInvalidControlPoint        = "control point is invalid"
	errorControlPointDeviceNotFound = "control point can't find the device (%s, %s)"
	errorPostActionResultFailed     = "post action (%s) failed '%s' : expected '%s'"
	errorPostActionSuccess          = "post action (%s) successed : expected failed"
	errorPostActionInvalidErrorType = "error object is invalid : %#v"
	errorPostActionInvalidErrorCode = "post action (%s) error code = %d : expected %d"
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

	dev, err := NewTestDevice()
	if err != nil {
		t.Error(err)
	}

	err = dev.Start()
	if err != nil {
		t.Error(err)
	}
	defer dev.Stop()

	devType := dev.DeviceType
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
	for range loopCnt {
		wait := time.Duration(cp.SearchMX*1000/loopCnt) * time.Millisecond
		time.Sleep(wait)
		var ok bool
		foundDev, ok = cp.FindDeviceByTypeAndUDN(devType, devUDN)
		if ok {
			break
		}
	}

	if foundDev == nil {
		t.Skipf(errorControlPointDeviceNotFound, devType, devUDN)
		return
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

	foundSetActionArg := &foundSetAction.ArgumentList.Arguments[0]
	foundSetActionArg.Value = postValue

	err = foundSetAction.Post()
	if err != nil {
		t.Error(err)
	}

	// post action (get)

	devGetAction, _ := dev.GetSwitchPowerGetTargetAction()

	foundGetAction, err := foundService.GetActionByName(devGetAction.Name)
	if err != nil {
		t.Error(err)
	}

	err = foundGetAction.Post()
	if err != nil {
		t.Error(err)
	}

	foundGetActionArg := foundGetAction.ArgumentList.Arguments[0]
	if foundGetActionArg.Value != postValue {
		t.Errorf(errorPostActionResultFailed, foundGetActionArg.Name, foundGetActionArg.Value, postValue)
	}

	// post optionl action which is not implemented yet

	devOptAction, _ := dev.GetOptionalAction()

	foundOptAction, err := foundService.GetActionByName(devOptAction.Name)
	if err != nil {
		t.Error(err)
	}

	err = foundOptAction.Post()
	if err == nil {
		t.Errorf(errorPostActionSuccess, foundOptAction.Name)
	}

	var upnpErr Error
	if errors.As(err, &upnpErr) == false {
		t.Errorf(errorPostActionInvalidErrorType, err)
	}

	expectErrorCode := control.ErrorOptionalActionNotImplemented
	if upnpErr.GetCode() != expectErrorCode {
		t.Errorf(errorPostActionInvalidErrorCode, foundOptAction.Name, upnpErr.GetCode(), expectErrorCode)
	}

	// stop control point

	err = cp.Stop()
	if err != nil {
		t.Error(err)
	}
}
