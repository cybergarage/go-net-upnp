// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"bytes"
	"io"
	gohttp "net/http"
	"testing"

	"github.com/cybergarage/go-net-upnp/net/upnp/util"
)

func TestNewClient(t *testing.T) {
	NewClient()
}

type RoundTripFunc func(req *gohttp.Request) *gohttp.Response

func (f RoundTripFunc) RoundTrip(req *gohttp.Request) (*gohttp.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *Client {
	return &Client{
		&gohttp.Client{
			Transport: RoundTripFunc(fn),
		},
	}
}

func TestDoWithoutUserAgentAddsDefaultUserAgent(t *testing.T) {
	testURL := "http://example.com"
	defaultUA := util.GetUserAgent()
	testMethod := gohttp.MethodGet
	testBody := io.NopCloser(bytes.NewBufferString(`OK`))

	testRequest, err := NewRequest(testMethod, testURL, testBody)
	if err != nil {
		t.Fatal(err)
	}

	client := NewTestClient(func(req *gohttp.Request) *gohttp.Response {
		if testURL != req.URL.String() {
			t.Fatalf(assertMessageStringsNotEqual, testURL, req.URL.String())
		}
		if defaultUA != req.UserAgent() {
			t.Fatalf(assertMessageStringsNotEqual, defaultUA, req.UserAgent())
		}
		if testMethod != req.Method {
			t.Fatalf(assertMessageStringsNotEqual, testMethod, req.Method)
		}
		if testBody != req.Body {
			t.Fatal("Body does not equal expected body")
		}
		return &gohttp.Response{
			StatusCode: gohttp.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(gohttp.Header),
		}
	})

	client.Do(testRequest)
}

func TestDoWithUserAgentDoesntOverwrite(t *testing.T) {
	testURL := "http://example.com"
	testUA := "foo/bar"
	testMethod := gohttp.MethodGet
	testBody := io.NopCloser(bytes.NewBufferString(`OK`))

	testRequest, err := NewRequest(testMethod, testURL, testBody)
	if err != nil {
		t.Fatal(err)
	}
	testRequest.Header.Set(UserAgent, testUA)

	client := NewTestClient(func(req *gohttp.Request) *gohttp.Response {
		if testURL != req.URL.String() {
			t.Fatalf(assertMessageStringsNotEqual, testURL, req.URL.String())
		}
		if testUA != req.UserAgent() {
			t.Fatalf(assertMessageStringsNotEqual, testUA, req.UserAgent())
		}
		if testMethod != req.Method {
			t.Fatalf(assertMessageStringsNotEqual, testMethod, req.Method)
		}
		if testBody != req.Body {
			t.Fatal("Body does not equal expected body")
		}
		return &gohttp.Response{
			StatusCode: gohttp.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(gohttp.Header),
		}
	})

	client.Do(testRequest)
}
