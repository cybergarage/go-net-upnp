// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"fmt"
	"testing"

	"github.com/cybergarage/go-net-upnp/net/upnp/control"
)

const (
	errorActionRequestInvalidName    = "invalid action name = '%s': expected '%s'"
	errorActionRequestInvalidArgCnt  = "invalid arguments count (%d) = : expected (%d)'"
	errorActionRequestInvalidArg     = "invalid param (%s) = '%s': expected (%s) = '%s'"
	errorActionResponseInvalidName   = "invalid action name = '%s': expected '%s'"
	errorActionResponseInvalidArgCnt = "invalid arguments count (%d) = : expected (%d)'"
	errorActionResponseInvalidArg    = "invalid param (%s) = '%s': expected (%s) = '%s'"
)

func TestMarshalActionRequestFromAction(t *testing.T) {
	const nArgs = 5

	// create a new argument

	action := NewAction()
	action.Name = "Hello"
	action.ArgumentList.Arguments = make([]Argument, nArgs)
	argNames := make([]string, nArgs)
	argValues := make([]string, nArgs)
	for n := 0; n < nArgs; n++ {
		arg := NewArgument()

		argNames[n] = fmt.Sprintf("Name%d", n)
		arg.Name = argNames[n]

		argValues[n] = fmt.Sprintf("Value%d", n)
		arg.Value = argValues[n]

		arg.SetDirection(InDirection)

		action.ArgumentList.Arguments[n] = *arg
	}

	// marshal the argument request

	actionReq, err := NewActionRequestFromAction(action)
	if err != nil {
		t.Error(err)
	}

	soapReq, err := actionReq.SOAPContentString()
	if err != nil {
		t.Error(err)
	}

	// unmarshal the action request

	actionReq, err = control.NewActionRequestFromSOAPBytes([]byte(soapReq))
	if err != nil {
		t.Error(err)
	}

	checkActionRequestParams(t, actionReq, action.Name, nArgs, argNames, argValues)
}

func checkActionRequestParams(t *testing.T, req *control.ActionRequest, actionName string, argCnt int, argNames []string, argValues []string) {
	action, err := req.GetAction()
	if err != nil {
		t.Error(err)
	}

	expectValue := actionName
	if action.Name != expectValue {
		t.Errorf(errorActionRequestInvalidName, action.Name, expectValue)
	}

	expectedArgCnt := argCnt
	if len(action.Arguments) != expectedArgCnt {
		t.Errorf(errorActionRequestInvalidArgCnt, len(action.Arguments), expectedArgCnt)
	}

	expactedArgNames := argNames
	expactedArgValues := argValues

	for n := 0; n < len(expactedArgNames); n++ {
		arg := action.Arguments[n]
		if arg.Name != expactedArgNames[n] {
			t.Errorf(errorActionRequestInvalidArg, arg.Name, arg.Value, expactedArgNames[n], expactedArgValues[n])
		}
		if arg.Value != expactedArgValues[n] {
			t.Errorf(errorActionRequestInvalidArg, arg.Name, arg.Value, expactedArgNames[n], expactedArgValues[n])
		}
	}
}

func TestMarshalActionResponseFromAction(t *testing.T) {
	const nArgs = 5

	// create a new argument

	action := NewAction()
	action.Name = "Hello"
	action.ArgumentList.Arguments = make([]Argument, nArgs)
	argNames := make([]string, nArgs)
	argValues := make([]string, nArgs)
	for n := 0; n < nArgs; n++ {
		arg := NewArgument()

		argNames[n] = fmt.Sprintf("Name%d", n)
		arg.Name = argNames[n]

		argValues[n] = fmt.Sprintf("Value%d", n)
		arg.Value = argValues[n]

		arg.SetDirection(OutDirection)

		action.ArgumentList.Arguments[n] = *arg
	}

	// marshal the argument request

	actionRes, err := NewActionResponseFromAction(action)
	if err != nil {
		t.Error(err)
	}

	soapRes, err := actionRes.SOAPContentString()
	if err != nil {
		t.Error(err)
	}

	// unmarshal the action request

	actionRes, err = control.NewActionResponseFromSOAPBytes([]byte(soapRes))
	if err != nil {
		t.Error(err)
	}

	checkActionResponseParams(t, actionRes, action.Name, nArgs, argNames, argValues)
}

func checkActionResponseParams(t *testing.T, res *control.ActionResponse, actionName string, argCnt int, argNames []string, argValues []string) {
	action, err := res.GetAction()
	if err != nil {
		t.Error(err)
	}

	expectValue := actionName
	if action.Name != expectValue {
		t.Errorf(errorActionResponseInvalidName, action.Name, expectValue)
	}

	expectedArgCnt := argCnt
	if len(action.Arguments) != expectedArgCnt {
		t.Errorf(errorActionResponseInvalidArgCnt, len(action.Arguments), expectedArgCnt)
	}

	expactedArgNames := argNames
	expactedArgValues := argValues

	for n := 0; n < len(expactedArgNames); n++ {
		arg := action.Arguments[n]
		if arg.Name != expactedArgNames[n] {
			t.Errorf(errorActionResponseInvalidArg, arg.Name, arg.Value, expactedArgNames[n], expactedArgValues[n])
		}
		if arg.Value != expactedArgValues[n] {
			t.Errorf(errorActionResponseInvalidArg, arg.Name, arg.Value, expactedArgNames[n], expactedArgValues[n])
		}
	}
}
