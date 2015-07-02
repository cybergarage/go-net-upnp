// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"testing"
)

func TestNewRequest(t *testing.T) {
	NewRequest()
}

func TestSSDPSearchRequest(t *testing.T) {

	const SearchRequest = "M-SEARCH * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"MX: 3\r\n" +
		"ST: upnp:rootdevices\r\n" +
		"\r\n"

	ssdpReq, err := NewRequestFromString(SearchRequest)
	if err != nil {
		t.Error(err)
	}

	if !ssdpReq.IsSearchRequest() {
		t.Errorf(testErrorMsgBadMethod, ssdpReq.FirstLines[0], M_SEARCH)
	}

	value, _ := ssdpReq.GetHost()
	expectValue := "239.255.255.250:1900"
	if value != expectValue {
		t.Errorf(testErrorMsgBadHeader, HOST, value, expectValue)
	}

	value, _ = ssdpReq.GetMAN()
	expectValue = "\"ssdp:discover\""
	if value != expectValue {
		t.Errorf(testErrorMsgBadHeader, MAN, value, expectValue)
	}

	value, _ = ssdpReq.GetMX()
	expectValue = "3"
	if value != expectValue {
		t.Errorf(testErrorMsgBadHeader, MX, value, expectValue)
	}

	value, _ = ssdpReq.GetST()
	expectValue = "upnp:rootdevices"
	if value != expectValue {
		t.Errorf(testErrorMsgBadHeader, ST, value, expectValue)
	}
}
