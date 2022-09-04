// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"testing"

	"github.com/cybergarage/go-net-upnp/net/upnp/http"
)

func TestNewResponse(t *testing.T) {
	NewResponse()
}

func TestSearchResponse(t *testing.T) {
	const SearchResponse = "" +
		"HTTP/1.1 200 OK\r\n" +
		"CACHE-CONTROL: 400\r\n" +
		"DATE: \r\n" +
		"EXT: \r\n" +
		"LOCATION: \r\n" +
		"SERVER: OS/version UPnP/1.1 product/version \r\n" +
		"ST: \r\n" +
		"USN: \r\n" +
		"BOOTID.UPNP.ORG: \r\n" +
		"CONFIGID.UPNP.ORG: \r\n" +
		"SEARCHPORT.UPNP.ORG: \r\n" +
		"\r\n"

	res, err := NewResponseFromString(SearchResponse)
	if err != nil {
		t.Error(err)
	}

	// Check StatusCode

	code := res.GetStatusCode()
	if code != http.StatusOK {
		t.Errorf(testErrorMsgBadStatusCode, code, 200)
	}

	// Check Headers

	var headerValue, expectValue string
	//var headerInt, expectInt int

	headerValue, err = res.GetEXT()
	if err != nil {
		t.Error(err)
	}
	expectValue = ""
	if headerValue != expectValue {
		t.Errorf(testErrorMsgBadHeader, EXT, headerValue, expectValue)
	}

}
