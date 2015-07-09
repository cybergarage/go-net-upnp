// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
	"fmt"
	"testing"
)

const (
	errorActionRequestInvalidName  = "invalid action name = '%s': expected '%s'"
	errorActionRequestInvalidParam = "invalid param (%s) = '%s': expected (%s) = '%s'"
)

const (
	testSoapActionRequest = xml.Header + "\n" +
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
)

func TestNewActionRequest(t *testing.T) {
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

	expactedParamNames := []string{"Value1", "Value2", "Value3", "Value4"}
	expactedParamValues := []string{"100", "200", "300", "400"}

	for n := 0; n < len(expactedParamNames); n++ {
		fmt.Printf("expactedParamNames[%d] %p %d\n", n, action.Arguments, len(action.Arguments))
		arg := action.Arguments[n]
		if arg.Name != expactedParamNames[n] {
			t.Errorf(errorActionRequestInvalidParam, arg.Name, arg.Value, expactedParamNames[n], expactedParamValues[n])
		}
		if arg.Value != expactedParamValues[n] {
			t.Errorf(errorActionRequestInvalidParam, arg.Name, arg.Value, expactedParamNames[n], expactedParamValues[n])
		}
	}
}
