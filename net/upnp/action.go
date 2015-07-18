// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"./control"
	"./http"
	"./log"
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

func (self *Action) reviseParentObject() error {
	for n := 0; n < len(self.ArgumentList.Arguments); n++ {
		arg := &self.ArgumentList.Arguments[n]
		arg.ParentAction = self
	}

	return nil
}

// GetArguments returns all arguments
func (self *Action) GetArguments() []*Argument {
	argCnt := len(self.ArgumentList.Arguments)
	args := make([]*Argument, argCnt)
	for n := 0; n < argCnt; n++ {
		args[n] = &self.ArgumentList.Arguments[n]
	}
	return args
}

// GetInputArguments returns all input arguments
func (self *Action) GetInputArguments() []*Argument {
	args := make([]*Argument, 0)
	for n := 0; n < len(self.ArgumentList.Arguments); n++ {
		arg := &self.ArgumentList.Arguments[n]
		if !arg.IsInDirection() {
			continue
		}
		args = append(args, arg)
	}
	return args
}

// GetOutputArguments returns all output arguments
func (self *Action) GetOutputArguments() []*Argument {
	args := make([]*Argument, 0)
	for n := 0; n < len(self.ArgumentList.Arguments); n++ {
		arg := &self.ArgumentList.Arguments[n]
		if !arg.IsOutDirection() {
			continue
		}
		args = append(args, arg)
	}
	return args
}

// GetArgumentByName returns an argument by the specified name
func (self *Action) GetArgumentByName(name string) (*Argument, error) {
	for n := 0; n < len(self.ArgumentList.Arguments); n++ {
		arg := &self.ArgumentList.Arguments[n]
		if arg.Name == name {
			return arg, nil
		}
	}
	return nil, errors.New(fmt.Sprintf(errorActionArgumentNotFound, name))
}

// SetArgumentString sets a value into the specified argument
func (self *Action) SetArgumentString(name string, value string) error {
	arg, err := self.GetArgumentByName(name)
	if err != nil {
		return err
	}
	return arg.SetString(value)
}

// GetArgumentString return a value into the specified argument
func (self *Action) GetArgumentString(name string) (string, error) {
	arg, err := self.GetArgumentByName(name)
	if err != nil {
		return "", err
	}
	return arg.GetString()
}

// SetArgumentInt sets a value into the specified argument
func (self *Action) SetArgumentInt(name string, value int) error {
	return self.SetArgumentString(name, strconv.Itoa(value))
}

// GetArgumentInt return a value into the specified argument
func (self *Action) GetArgumentInt(name string) (int, error) {
	value, err := self.GetArgumentString(name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(value)
}

// setArgumentsByActionControl sets control arguments into the specified argument
func (self *Action) setArgumentsByActionControl(actionCtrl *control.ActionControl) error {
	ctrlAction, err := actionCtrl.GetAction()
	if err != nil {
		return err
	}

	if ctrlAction.Name != self.Name {
		return errors.New(fmt.Sprintf(errorActionNameIsInvalid, ctrlAction.Name, self.Name))
	}

	for n := 0; n < len(ctrlAction.Arguments); n++ {
		ctrlArg := ctrlAction.Arguments[n]
		selfArg, err := self.GetArgumentByName(ctrlArg.Name)
		if err != nil {
			continue
		}
		selfArg.Value = ctrlArg.Value
	}

	return nil
}

// SetArgumentsByActionRequest sets request arguments into the specified argument
func (self *Action) SetArgumentsByActionRequest(actionReq *control.ActionRequest) error {
	return self.setArgumentsByActionControl(actionReq.ActionControl)
}

// SetArgumentsByActionResponse sets response arguments into the specified argument
func (self *Action) SetArgumentsByActionResponse(actionRes *control.ActionResponse) error {
	return self.setArgumentsByActionControl(actionRes.ActionControl)
}

// Post sends the specified arguments into the deveice.
func (self *Action) Post() error {
	// post response

	req, err := NewActionRequestFromAction(self)
	if err != nil {
		return err
	}

	soapReqStr, err := req.SOAPContentString()
	if err != nil {
		return err
	}

	service := self.ParentService
	if service == nil {
		return fmt.Errorf(errorActionHasNoParentService, self.Name)
	}

	controlAbsURL, err := service.GetAbsoluteControlURL()
	if err != nil {
		return err
	}

	soapAction := service.ServiceType + http.SoapActionDelim + self.Name
	httpReq, err := http.NewSOAPRequest(controlAbsURL, soapAction, strings.NewReader(soapReqStr))
	if err != nil {
		return err
	}

	log.Trace(fmt.Sprintf("action req = \n%s", soapReqStr))

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
	soapResBytes, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}

	log.Trace(fmt.Sprintf("action res [%d] = \n%s", statusCode, string(soapResBytes)))

	// parse response

	if statusCode == http.StatusOK {
		actionRes, err := control.NewActionResponseFromSOAPBytes(soapResBytes)
		if err != nil {
			return err
		}
		err = self.SetArgumentsByActionResponse(actionRes)
		if err != nil {
			return err
		}
	} else {
		upnpErrRes, err := control.NewErrorResponseFromSOAPBytes(soapResBytes)
		if err != nil {
			return err
		}
		return upnpErrRes.Envelope.Body.Fault.Detail.UPnPError
	}

	return nil
}
