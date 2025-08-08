// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-net-upnp/net/upnp/control"
	"github.com/cybergarage/go-net-upnp/net/upnp/http"
)

const (
	errorActionArgumentNotFound   = "argument (%s) is not found"
	errorActionNameIsInvalid      = "action name of action response (%s) is not equal this action name (%s)"
	errorActionHasNoParentService = "action (%s) has no parent service"
)

// A Action represents a UPnP action.
type Action struct {
	XMLName       xml.Name     `xml:"action"`
	Name          string       `xml:"name"`
	ArgumentList  ArgumentList `xml:"argumentList"`
	ParentService *Service     `xml:"-"`
}

// A ActionList represents a UPnP action list.
type ActionList struct {
	XMLName xml.Name `xml:"actionList"`
	Actions []Action `xml:"action"`
}

// NewAction returns a new Action.
func NewAction() *Action {
	action := &Action{}
	return action
}

func (action *Action) reviseParentObject() error {
	for n := range len(action.ArgumentList.Arguments) {
		arg := &action.ArgumentList.Arguments[n]
		arg.ParentAction = action
	}

	return nil
}

// GetArguments returns all arguments.
func (action *Action) GetArguments() []*Argument {
	argCnt := len(action.ArgumentList.Arguments)
	args := make([]*Argument, argCnt)
	for n := range argCnt {
		args[n] = &action.ArgumentList.Arguments[n]
	}
	return args
}

// GetInputArguments returns all input arguments.
func (action *Action) GetInputArguments() []*Argument {
	args := make([]*Argument, 0)
	for n := range len(action.ArgumentList.Arguments) {
		arg := &action.ArgumentList.Arguments[n]
		if !arg.IsInDirection() {
			continue
		}
		args = append(args, arg)
	}
	return args
}

// GetOutputArguments returns all output arguments.
func (action *Action) GetOutputArguments() []*Argument {
	args := make([]*Argument, 0)
	for n := range len(action.ArgumentList.Arguments) {
		arg := &action.ArgumentList.Arguments[n]
		if !arg.IsOutDirection() {
			continue
		}
		args = append(args, arg)
	}
	return args
}

// GetArgumentByName returns an argument by the specified name.
func (action *Action) GetArgumentByName(name string) (*Argument, error) {
	for n := range len(action.ArgumentList.Arguments) {
		arg := &action.ArgumentList.Arguments[n]
		if arg.Name == name {
			return arg, nil
		}
	}
	return nil, fmt.Errorf(errorActionArgumentNotFound, name)
}

// SetArgumentString sets a string value into the specified argument.
func (action *Action) SetArgumentString(name string, value string) error {
	arg, err := action.GetArgumentByName(name)
	if err != nil {
		return err
	}
	return arg.SetString(value)
}

// GetArgumentString return a string value into the specified argument.
func (action *Action) GetArgumentString(name string) (string, error) {
	arg, err := action.GetArgumentByName(name)
	if err != nil {
		return "", err
	}
	return arg.GetString()
}

// SetArgumentInt sets a integer value into the specified argument.
func (action *Action) SetArgumentInt(name string, value int) error {
	arg, err := action.GetArgumentByName(name)
	if err != nil {
		return err
	}
	return arg.SetInt(value)
}

// GetArgumentInt return a integer value into the specified argument.
func (action *Action) GetArgumentInt(name string) (int, error) {
	arg, err := action.GetArgumentByName(name)
	if err != nil {
		return 0, err
	}
	return arg.GetInt()
}

// SetArgumentFloat sets a integer value into the specified argument.
func (action *Action) SetArgumentFloat(name string, value float64) error {
	arg, err := action.GetArgumentByName(name)
	if err != nil {
		return err
	}
	return arg.SetFloat(value)
}

// GetArgumentFloat return a integer value into the specified argument.
func (action *Action) GetArgumentFloat(name string) (float64, error) {
	arg, err := action.GetArgumentByName(name)
	if err != nil {
		return 0, err
	}
	return arg.GetFloat()
}

// SetArgumentBool sets a boolean value into the specified argument.
func (action *Action) SetArgumentBool(name string, value bool) error {
	arg, err := action.GetArgumentByName(name)
	if err != nil {
		return err
	}
	return arg.SetBool(value)
}

// GetArgumentBool return a boolean value into the specified argument.
func (action *Action) GetArgumentBool(name string) (bool, error) {
	arg, err := action.GetArgumentByName(name)
	if err != nil {
		return false, err
	}
	return arg.GetBool()
}

// setArgumentsByActionControl sets control arguments into the specified argument.
func (action *Action) setArgumentsByActionControl(actionCtrl *control.ActionControl) error {
	ctrlAction, err := actionCtrl.GetAction()
	if err != nil {
		return err
	}

	if ctrlAction.Name != action.Name {
		return fmt.Errorf(errorActionNameIsInvalid, ctrlAction.Name, action.Name)
	}

	for n := range len(ctrlAction.Arguments) {
		ctrlArg := ctrlAction.Arguments[n]
		actionArg, err := action.GetArgumentByName(ctrlArg.Name)
		if err != nil {
			continue
		}
		actionArg.Value = ctrlArg.Value
	}

	return nil
}

// SetArgumentsByActionRequest sets request arguments into the specified argument.
func (action *Action) SetArgumentsByActionRequest(actionReq *control.ActionRequest) error {
	return action.setArgumentsByActionControl(actionReq.ActionControl)
}

// SetArgumentsByActionResponse sets response arguments into the specified argument.
func (action *Action) SetArgumentsByActionResponse(actionRes *control.ActionResponse) error {
	return action.setArgumentsByActionControl(actionRes.ActionControl)
}

// Post sends the specified arguments into the deveice.
func (action *Action) Post() error {
	// post response

	req, err := NewActionRequestFromAction(action)
	if err != nil {
		return err
	}

	soapReqStr, err := req.SOAPContentString()
	if err != nil {
		return err
	}

	service := action.ParentService
	if service == nil {
		return fmt.Errorf(errorActionHasNoParentService, action.Name)
	}

	controlAbsURL, err := service.GetAbsoluteControlURL()
	if err != nil {
		return err
	}

	soapAction := service.ServiceType + http.SOAPActionDelim + action.Name
	httpReq, err := http.NewSOAPRequest(controlAbsURL, soapAction, strings.NewReader(soapReqStr))
	if err != nil {
		return err
	}

	log.Tracef("action req = \n%s", soapReqStr)

	httpClient, err := http.NewClient()
	if err != nil {
		return err
	}

	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		return err
	}

	// read response

	statusCode := httpRes.StatusCode
	defer httpRes.Body.Close()
	soapResBytes, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}

	log.Tracef("action res [%d] = \n%s", statusCode, string(soapResBytes))

	// parse response

	if statusCode == http.StatusOK {
		actionRes, err := control.NewActionResponseFromSOAPBytes(soapResBytes)
		if err != nil {
			return err
		}
		err = action.SetArgumentsByActionResponse(actionRes)
		if err != nil {
			return err
		}
	} else {
		upnpErrRes, err := control.NewErrorResponseFromSOAPBytes(soapResBytes)
		if err != nil {
			return err
		}
		return &upnpErrRes.Envelope.Body.Fault.Detail.UPnPError
	}

	return nil
}
