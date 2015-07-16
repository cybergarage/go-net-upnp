// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	errorSelectDeviceCancel  = "device is not selected"
	errorSelectDeviceInvalid = "device no (%d) is invalid"
	deviceIsSelected         = "[%d] '%s' number selected"

	errorSelectServiceCancel  = "service is not selected"
	errorSelectServiceInvalid = "service number (%d) is invalid"
	serviceIsSelected         = "[%d] '%s' is selected"

	errorSelectActionCancel  = "action is not selected"
	errorSelectActionInvalid = "action number (%d) is invalid"
	actionIsSelected         = "[%d] '%s' is selected"

	headerDeviceList  = "==== Select a target device (%d) ===="
	headerServiceList = "==== Select a target service (%d) ===="
	headerActionList  = "==== Select a target action (%d) ===="

	selectDeviceMsg  = "Input device number"
	selectServiceMsg = "Input service number"
	selectActionMsg  = "Input action number"
	inputArgParamMsg = "[%d] %s : Input param"
	postResultMsg    = "[%d] %s = '%s'"
)

func printMsg(msg string) {
	os.Stdout.WriteString(fmt.Sprintf("%s\n", msg))
}

func printError(msg string) {
	os.Stderr.WriteString(fmt.Sprintf("%s\n", msg))
}

func inputMsg(msg string) string {
	os.Stdout.WriteString(fmt.Sprintf("%s : ", msg))
	os.Stdout.Sync()
	input, err := ReadKeyboardLine()
	if err != nil {
		return ""
	}
	return string(input)
}

func inputNo(msg string) (int, error) {
	kbInput := inputMsg(msg)
	i, err := strconv.Atoi(kbInput)
	if err != nil {
		return 0, err
	}
	return i, nil
}

////////////////////////////////////////
// h : Help
////////////////////////////////////////

const (
	H_KEY  = 'h'
	H_DESC = "print this (H)elp message"
)

func HelpAction(cp *ControlPoint) bool {
	for key, action := range cp.ControlPointActionManager.Commands {
		printMsg(fmt.Sprintf("'%c' : %s", key, action.Desc))
	}
	return true
}

////////////////////////////////////////
// q : Quit
////////////////////////////////////////

const (
	Q_KEY  = 'q'
	Q_DESC = "(!)uit"
)

func QuitAction(cp *ControlPoint) bool {
	return false
}

////////////////////////////////////////
// s : Search
////////////////////////////////////////

const (
	S_KEY  = 's'
	S_DESC = "(S)earch root devices"
)

func SearchAction(cp *ControlPoint) bool {
	err := cp.SearchRootDevice()
	if err != nil {
		return false
	}
	return true
}

////////////////////////////////////////
// a : post action
////////////////////////////////////////

const (
	A_KEY  = 'a'
	A_DESC = "post (A)ction"
)

func PostAction(cp *ControlPoint) bool {
	// select a target device

	printRootDevices(cp)

	devNo, err := inputNo(selectDeviceMsg)
	if err != nil {
		printError(errorSelectDeviceCancel)
		return false
	}

	rootDevs := cp.GetRootDevices()
	if (devNo < 0) || (len(rootDevs) < devNo) {
		printError(fmt.Sprintf(errorSelectDeviceInvalid, devNo))
		return false
	}

	dev := rootDevs[devNo]
	printMsg(fmt.Sprintf(deviceIsSelected, devNo, dev.FriendlyName))

	// select a target service

	services := dev.GetServices()
	serviceCnt := len(services)

	printMsg(fmt.Sprintf(headerServiceList, serviceCnt))
	for n, service := range services {
		printMsg(fmt.Sprintf("[%d] '%s'", n, service.ServiceType))
	}

	serviceNo, err := inputNo(selectServiceMsg)
	if err != nil {
		printError(errorSelectServiceCancel)
		return false
	}

	if (serviceNo < 0) || (serviceCnt < serviceNo) {
		printError(fmt.Sprintf(errorSelectServiceInvalid, serviceNo))
		return false
	}

	service := &dev.ServiceList.Services[serviceNo]
	printMsg(fmt.Sprintf(serviceIsSelected, serviceNo, service.ServiceType))

	// select a target action

	actions := service.GetActions()
	actionCnt := len(actions)

	printMsg(fmt.Sprintf(headerActionList, actionCnt))
	for n, action := range actions {
		printMsg(fmt.Sprintf("[%d] '%s'", n, action.Name))
	}

	actionNo, err := inputNo(selectActionMsg)
	if err != nil {
		printError(errorSelectActionCancel)
		return false
	}

	if (actionNo < 0) || (actionCnt < actionNo) {
		printError(fmt.Sprintf(errorSelectActionInvalid, actionNo))
		return false
	}

	action := &service.ActionList.Actions[actionNo]
	printMsg(fmt.Sprintf(actionIsSelected, actionNo, action.Name))

	// input argument params

	inputArgs := action.GetInputArguments()
	for n, arg := range inputArgs {
		argParam := inputMsg(fmt.Sprintf(inputArgParamMsg, n, arg.Name))
		arg.Value = argParam
	}

	// post action

	upnpErr := action.Post()
	if upnpErr == nil {
		outArgs := action.GetOutputArguments()
		for n, arg := range outArgs {
			printMsg(fmt.Sprintf(postResultMsg, n, arg.Name, arg.Value))
		}
	} else {
		printError(upnpErr.Error())
	}

	return true
}

////////////////////////////////////////
// p : Print
////////////////////////////////////////

const (
	P_KEY  = 'p'
	P_DESC = "(P)rint found devices"
)

func printRootDevices(cp *ControlPoint) {
	rootDevs := cp.GetRootDevices()
	printMsg(fmt.Sprintf(headerDeviceList, len(rootDevs)))
	for n, dev := range rootDevs {
		printMsg(fmt.Sprintf("[%d] '%s', '%s'", n, dev.FriendlyName, dev.DeviceType))
	}
}

func PrintAction(cp *ControlPoint) bool {
	printRootDevices(cp)
	return true
}
