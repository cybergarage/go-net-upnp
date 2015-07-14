// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
	"fmt"
	"testing"

	"net/upnp"
)

const (
	errorActionRequestInvalidName   = "invalid action name = '%s': expected '%s'"
	errorActionRequestInvalidArgCnt = "invalid arguments count (%d) = : expected (%d)'"
	errorActionRequestInvalidArg    = "invalid param (%s) = '%s': expected (%s) = '%s'"
)

func TestNewActionRequestNoArgument(t *testing.T) {
	const testSoapActionRequest = xml.Header + "\n" +
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">" +
		"  <s:Body>" +
		"    <u:SetValue xmlns:u=\"urn:schemas-upnp-org:service:serviceType:v\">" +
		"    </u:SetValue>" +
		"  </s:Body>" +
		"</s:Envelope>"

	req, err := NewActionRequestFromSOAPString(testSoapActionRequest)
	if err != nil {
		t.Error(err)
	}

	checkActionRequestParams(t, req, "SetValue", 0, []string{}, []string{})
}

func TestNewActionRequestOneSpaceArgument(t *testing.T) {
	const testSoapActionRequest = xml.Header + "\n" +
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">" +
		"  <s:Body>" +
		"    <u:SetValue xmlns:u=\"urn:schemas-upnp-org:service:serviceType:v\">" +
		"      <Value>Hello World</Value>" +
		"    </u:SetValue>" +
		"  </s:Body>" +
		"</s:Envelope>"

	req, err := NewActionRequestFromSOAPString(testSoapActionRequest)
	if err != nil {
		t.Error(err)
	}

	expactedArgNames := []string{"Value"}
	expactedArgValues := []string{"Hello World"}

	checkActionRequestParams(t, req, "SetValue", 1, expactedArgNames, expactedArgValues)
}

func TestNewActionRequestOneArguments(t *testing.T) {
	const testSoapActionRequest = xml.Header + "\n" +
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">" +
		"  <s:Body>" +
		"    <u:SetValue xmlns:u=\"urn:schemas-upnp-org:service:serviceType:v\">" +
		"      <Value>100</Value>" +
		"    </u:SetValue>" +
		"  </s:Body>" +
		"</s:Envelope>"

	req, err := NewActionRequestFromSOAPString(testSoapActionRequest)
	if err != nil {
		t.Error(err)
	}

	expactedArgNames := []string{"Value"}
	expactedArgValues := []string{"100"}

	checkActionRequestParams(t, req, "SetValue", 1, expactedArgNames, expactedArgValues)
}

func TestNewActionRequestForArguments(t *testing.T) {
	const testSoapActionRequest = xml.Header + "\n" +
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">" +
		"  <s:Body>" +
		"    <u:SetValue xmlns:u=\"urn:schemas-upnp-org:service:serviceType:v\">" +
		"      <Value1>100</Value1>" +
		"      <Value2>200</Value2>" +
		"      <Value3>300</Value3>" +
		"      <Value4>400</Value4>" +
		"    </u:SetValue>" +
		"  </s:Body>" +
		"</s:Envelope>"

	req, err := NewActionRequestFromSOAPString(testSoapActionRequest)
	if err != nil {
		t.Error(err)
	}

	expactedArgNames := []string{"Value1", "Value2", "Value3", "Value4"}
	expactedArgValues := []string{"100", "200", "300", "400"}

	checkActionRequestParams(t, req, "SetValue", 4, expactedArgNames, expactedArgValues)
}

func TestMarshalActionRequestFromAction(t *testing.T) {
	const nArgs = 5

	// create a new argument

	action := upnp.NewAction()
	action.Name = "Hello"
	action.ArgumentList.Arguments = make([]upnp.Argument, nArgs)
	argNames := make([]string, nArgs)
	argValues := make([]string, nArgs)
	for n := 0; n < nArgs; n++ {
		arg := upnp.NewArgument()

		argNames[n] = fmt.Sprintf("Name%d", n)
		arg.Name = argNames[n]

		argValues[n] = fmt.Sprintf("Value%d", n)
		arg.Value = argValues[n]

		arg.SetDirection(upnp.InDirection)

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

	actionReq, err = NewActionRequestFromSOAPString(soapReq)
	if err != nil {
		t.Error(err)
	}

	checkActionRequestParams(t, actionReq, action.Name, nArgs, argNames, argValues)
}

func checkActionRequestParams(t *testing.T, req *ActionRequest, actionName string, argCnt int, argNames []string, argValues []string) {
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
