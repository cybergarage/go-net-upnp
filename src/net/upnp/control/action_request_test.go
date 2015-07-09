// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
	"testing"
)

const (
	testSoapActionRequest = xml.Header + "\n" +
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">" +
		"  <s:Body>" +
		"    <u:SetValue xmlns:u=\"urn:schemas-upnp-org:service:serviceType:v\">" +
		"      <Value>123456</Value>" +
		"    </u:SetValue>" +
		"  </s:Body>" +
		"</s:Envelope>"
)

func TestNewActionRequest(t *testing.T) {
	_, err := NewActionRequestFromSOAPString(testSoapActionRequest)
	if err != nil {
		t.Error(err)
	}
}
