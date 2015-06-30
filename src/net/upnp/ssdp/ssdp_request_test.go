// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"testing"
)

const (
	errorSSDPRequestHeader = "%s is %s : expected %s"
)

func TestNewSSDPRequest(t *testing.T) {
	NewSSDPRequest()
}

func TestSSDPSearchRequest(t *testing.T) {

	const SearchRequest = "M-SEARCH * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"MX: 3\r\n" +
		"ST: upnp:rootdevices\r\n" +
		"\r\n"

	ssdpReq, err := NewSSDPRequestFromString(SearchRequest)
	if err != nil {
		t.Error(err)
	}

	value, _ := ssdpReq.GetHost()
	expectValue := "239.255.255.250:1900"
	if value != expectValue {
		t.Errorf(errorSSDPRequestHeader, HOST, value, expectValue)
	}

	value, _ = ssdpReq.GetMAN()
	expectValue = "\"ssdp:discover\""
	if value != expectValue {
		t.Errorf(errorSSDPRequestHeader, MAN, value, expectValue)
	}

	value, _ = ssdpReq.GetMX()
	expectValue = "3"
	if value != expectValue {
		t.Errorf(errorSSDPRequestHeader, MX, value, expectValue)
	}

	value, _ = ssdpReq.GetST()
	expectValue = "upnp:rootdevices"
	if value != expectValue {
		t.Errorf(errorSSDPRequestHeader, ST, value, expectValue)
	}
}
