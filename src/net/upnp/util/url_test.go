// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"testing"
)

const (
	errorUrlNotInvalid = "url (%s) is not invalid"
)

func TestInvalidBaseAndPathURL(t *testing.T) {
	bases := []string{"", "192.168.100.1"}
	paths := []string{""}

	for _, base := range bases {
		for _, path := range paths {
			url, err := GetAbsoluteURLFromBaseAndPath(base, path)
			if err == nil {
				t.Errorf(errorUrlNotInvalid, url.String())
			}
		}
	}
}

func TestValidBaseAndPathURL(t *testing.T) {
	bases := []string{"http://192.168.100.1/", "http://192.168.100.1"}
	paths := []string{"/index.html", "index.html"}

	for _, base := range bases {
		for _, path := range paths {
			_, err := GetAbsoluteURLFromBaseAndPath(base, path)
			if err != nil {
				t.Error(err)
			}
		}
	}
}
