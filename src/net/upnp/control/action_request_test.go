// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
	"testing"
)

const (
	errorActionRequestInvalidName   = "invalid action name = '%s': expected '%s'"
	errorActionRequestInvalidArgCnt = "invalid arguments count (%d) = : expected (%d)'"
	errorActionRequestInvalidArg    = "invalid param (%s) = '%s': expected (%s) = '%s'"
)

func TestNewActionRequestNoArguments(t *testing.T) {
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

	action, err := req.GetAction()
	if err != nil {
		t.Error(err)
	}

	expectValue := "SetValue"
	if action.Name != expectValue {
		t.Errorf(errorActionRequestInvalidName, action.Name, expectValue)
	}

	expectedArgCnt := 0
	if len(action.Arguments) != expectedArgCnt {
		t.Errorf(errorActionRequestInvalidArgCnt, len(action.Arguments), expectedArgCnt)
	}
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

	action, err := req.GetAction()
	if err != nil {
		t.Error(err)
	}

	expectValue := "SetValue"
	if action.Name != expectValue {
		t.Errorf(errorActionRequestInvalidName, action.Name, expectValue)
	}

	expectedArgCnt := 1
	if len(action.Arguments) != expectedArgCnt {
		t.Errorf(errorActionRequestInvalidArgCnt, len(action.Arguments), expectedArgCnt)
	}

	expactedArgNames := []string{"Value"}
	expactedArgValues := []string{"100"}

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

	action, err := req.GetAction()
	if err != nil {
		t.Error(err)
	}

	expectValue := "SetValue"
	if action.Name != expectValue {
		t.Errorf(errorActionRequestInvalidName, action.Name, expectValue)
	}

	expectedArgCnt := 4
	if len(action.Arguments) != expectedArgCnt {
		t.Errorf(errorActionRequestInvalidArgCnt, len(action.Arguments), expectedArgCnt)
	}

	expactedArgNames := []string{"Value1", "Value2", "Value3", "Value4"}
	expactedArgValues := []string{"100", "200", "300", "400"}

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
