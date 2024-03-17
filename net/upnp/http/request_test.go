// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"bytes"
	"io"
	gohttp "net/http"

	"testing"
)

func TestNewHTTPRequest(t *testing.T) {
	testURL := "http://example.com"
	testMethod := gohttp.MethodGet
	testBody := io.NopCloser(bytes.NewBufferString(`OK`))

	req, err := NewRequest(testMethod, testURL, testBody)
	if err != nil {
		t.Fatal(err)
	}

	if testURL != req.URL.String() {
		t.Fatalf(assertMessageStringsNotEqual, testURL, req.URL.String())
	}
	if testMethod != req.Method {
		t.Fatalf(assertMessageStringsNotEqual, testMethod, req.Method)
	}
	if testBody != req.Body {
		t.Fatal("Body does not equal expected body")
	}
}

func TestNewRequestFromRequest(t *testing.T) {
	req := Request{}
	if req.Request != nil {
		t.Fatal("Request has unexpected embedded \"net/http\".Request")
	}
	goRequest := gohttp.Request{}
	actualRequest := NewRequestFromRequest(&goRequest)
	if actualRequest.Request == nil {
		t.Fatal("Request has not embedded \"net/http\".Request")
	}
}
