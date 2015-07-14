// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
	"testing"
)

const (
	errorErrorResponseInvalidErrCode = "invalid error code (%d) : expected (%d)"
	errorErrorResponseInvalidErrMsg  = "invalid error msg (%s) : expected (%s)"
)

func TestNewErrorResponseNoArgument(t *testing.T) {
	const testSoapErrorResponse = xml.Header + "\n" +
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">" +
		"  <s:Body>" +
		"    <s:Fault>" +
		"      <faultcode>s:Client</faultcode>" +
		"      <faultstring>UPnPError</faultstring>" +
		"      <detail>" +
		"        <UPnPError xmlns=\"urn:schemas-upnp-org:control-1-0\">" +
		"          <errorCode>401</errorCode>" +
		"          <errorDescription>Invalid Action</errorDescription>" +
		"        </UPnPError>" +
		"      </detail>" +
		"    </s:Fault>" +
		"  </s:Body>" +
		"</s:Envelope>"

	res, err := NewErrorResponseFromSOAPBytes([]byte(testSoapErrorResponse))
	if err != nil {
		t.Error(err)
	}

	upnpErr, err := res.GetUPnPError()
	if err != nil {
		t.Error(err)
	}

	expectCode := 401
	if upnpErr.ErrorCode != expectCode {
		t.Errorf(errorErrorResponseInvalidErrCode, upnpErr.ErrorCode, expectCode)
	}

	expectedMsg := "Invalid Action"
	if upnpErr.ErrorDescription != expectedMsg {
		t.Errorf(errorErrorResponseInvalidErrMsg, upnpErr.ErrorDescription, expectedMsg)
	}
}
