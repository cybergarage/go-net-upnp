// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"strconv"
	"testing"
)

func TestNewRequest(t *testing.T) {
	NewRequest()
}
func TestSSDPSearchRequest(t *testing.T) {

	const SearchRequest = "" +
		"M-SEARCH * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"MX: 3\r\n" +
		"ST: upnp:rootdevice\r\n" +
		"\r\n"

	req, err := NewRequestFromString(SearchRequest)
	if err != nil {
		t.Error(err)
	}

	checkSSDPSearchRequest(t, req)
}

func TestSSDPUnformalSearchRequest(t *testing.T) {

	const SearchRequest = "" +
		"M-SEARCH * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"MX: 3\r\n" +
		"ST: upnp:rootdevice\r\n"

	req, err := NewRequestFromString(SearchRequest)
	if err != nil {
		t.Error(err)
	}

	checkSSDPSearchRequest(t, req)
}

func checkSSDPSearchRequest(t *testing.T, req *Request) {

	// Check Method

	if !req.IsSearchRequest() {
		t.Errorf(testErrorMsgBadMethod, req.FirstLines[0], M_SEARCH)
	}

	// Check Headers

	var headerValue, expectValue string
	var headerInt, expectInt int

	headerValue, _ = req.GetHost()
	expectValue = "239.255.255.250:1900"
	if headerValue != expectValue {
		t.Errorf(testErrorMsgBadHeader, HOST, headerValue, expectValue)
	}

	headerValue, _ = req.GetMAN()
	expectValue = "\"ssdp:discover\""
	if headerValue != expectValue {
		t.Errorf(testErrorMsgBadHeader, MAN, headerValue, expectValue)
	}
	if !req.IsDiscover() {
		t.Errorf(testErrorMsgBadHeader, MAN, headerValue, expectValue)
	}

	headerInt, _ = req.GetMX()
	expectInt = 3
	if headerInt != expectInt {
		t.Errorf(testErrorMsgBadHeader, MX, strconv.Itoa(headerInt), strconv.Itoa(expectInt))
	}

	headerValue, _ = req.GetST()
	expectValue = "upnp:rootdevice"
	if headerValue != expectValue {
		t.Errorf(testErrorMsgBadHeader, ST, headerValue, expectValue)
	}
	if !req.IsRootDevice() {
		t.Errorf(testErrorMsgBadHeader, ST, headerValue, expectValue)
	}
}
