// Copyright 2015 Satoshi Konno. All rights reserved.
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
	errorActionResponseInvalidName   = "invalid action name = '%s': expected '%s'"
	errorActionResponseInvalidArgCnt = "invalid arguments count (%d) = : expected (%d)'"
	errorActionResponseInvalidArg    = "invalid param (%s) = '%s': expected (%s) = '%s'"
)

func TestNewActionResponseNoArgument(t *testing.T) {
	const testSoapActionResponse = xml.Header + "\n" +
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">" +
		"  <s:Body>" +
		"    <u:SetValueResponse xmlns:u=\"urn:schemas-upnp-org:service:serviceType:v\">" +
		"    </u:SetValueResponse>" +
		"  </s:Body>" +
		"</s:Envelope>"

	res, err := NewActionResponseFromSOAPString(testSoapActionResponse)
	if err != nil {
		t.Error(err)
	}

	checkActionResponseParams(t, res, "SetValue", 0, []string{}, []string{})
}

func TestNewActionResponseOneSpaceArgument(t *testing.T) {
	const testSoapActionResponse = xml.Header + "\n" +
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">" +
		"  <s:Body>" +
		"    <u:SetValueResponse xmlns:u=\"urn:schemas-upnp-org:service:serviceType:v\">" +
		"      <Value>Hello World</Value>" +
		"    </u:SetValueResponse>" +
		"  </s:Body>" +
		"</s:Envelope>"

	res, err := NewActionResponseFromSOAPString(testSoapActionResponse)
	if err != nil {
		t.Error(err)
	}

	expactedArgNames := []string{"Value"}
	expactedArgValues := []string{"Hello World"}

	checkActionResponseParams(t, res, "SetValue", 1, expactedArgNames, expactedArgValues)
}

func TestNewActionResponseOneArguments(t *testing.T) {
	const testSoapActionResponse = xml.Header + "\n" +
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">" +
		"  <s:Body>" +
		"    <u:SetValueResponse xmlns:u=\"urn:schemas-upnp-org:service:serviceType:v\">" +
		"      <Value>100</Value>" +
		"    </u:SetValueResponse>" +
		"  </s:Body>" +
		"</s:Envelope>"

	res, err := NewActionResponseFromSOAPString(testSoapActionResponse)
	if err != nil {
		t.Error(err)
	}

	expactedArgNames := []string{"Value"}
	expactedArgValues := []string{"100"}

	checkActionResponseParams(t, res, "SetValue", 1, expactedArgNames, expactedArgValues)
}

func TestNewActionResponseForArguments(t *testing.T) {
	const testSoapActionResponse = xml.Header + "\n" +
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">" +
		"  <s:Body>" +
		"    <u:SetValueResponse xmlns:u=\"urn:schemas-upnp-org:service:serviceType:v\">" +
		"      <Value1>100</Value1>" +
		"      <Value2>200</Value2>" +
		"      <Value3>300</Value3>" +
		"      <Value4>400</Value4>" +
		"    </u:SetValueResponse>" +
		"  </s:Body>" +
		"</s:Envelope>"

	res, err := NewActionResponseFromSOAPString(testSoapActionResponse)
	if err != nil {
		t.Error(err)
	}

	expactedArgNames := []string{"Value1", "Value2", "Value3", "Value4"}
	expactedArgValues := []string{"100", "200", "300", "400"}

	checkActionResponseParams(t, res, "SetValue", 4, expactedArgNames, expactedArgValues)
}

func TestMarshalActionResponseFromAction(t *testing.T) {
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

		arg.SetDirection(upnp.OutDirection)

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

	actionRes, err = NewActionResponseFromSOAPString(soapRes)
	if err != nil {
		t.Error(err)
	}

	checkActionResponseParams(t, actionRes, action.Name, nArgs, argNames, argValues)
}

func checkActionResponseParams(t *testing.T, res *ActionResponse, actionName string, argCnt int, argNames []string, argValues []string) {
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
