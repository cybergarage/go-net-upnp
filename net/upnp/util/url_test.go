// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"regexp"
	"testing"
)

const (
	errorUrlNotInvalid = "url '%s' is not invalid"
	errorInvalidUrl    = "invalid url '%s' : expected '%s'"
)

func badBases() []string {
	return []string{
		"",
		"192.168.100.1",
		"yahoo.co.jp",
		"bad://192.168.100.1",
		"bad://yahoo.co.jp",
	}
}

func okBases() []string {
	return []string{
		"http://192.168.100.1/",
		"http://192.168.100.1",
		"http://yahoo.co.jp",
		"http://yahoo.co.jp/",
		"http://192.168.100.1:80/",
		"http://192.168.100.1:80",
		"http://yahoo.co.jp:80",
		"http://yahoo.co.jp/:80",
		"https://192.168.100.1/",
		"https://192.168.100.1",
		"https://yahoo.co.jp",
		"https://yahoo.co.jp/",
		"https://192.168.100.1:80/",
		"https://192.168.100.1:80",
		"https://yahoo.co.jp:80",
		"https://yahoo.co.jp/:80",
	}
}

func badPaths() []string {
	return []string{
		"{}",
		"[]",
		"\\",
	}
}

func okPaths() []string {
	return []string{
		"",
		"/",
		"index.html",
		"/index.html",
		"foo/index.html",
		"/foo/index.html",
	}
}

func absPaths() []string {
	return []string{
		"http://192.168.100.1",
		"http://192.168.100.1/",
		"http://192.168.100.1/index.html",
		"http://192.168.100.1/foo/index.html",
		"http://cybergarage.org",
		"http://cybergarage.org/",
		"http://yahoo.co.jp/index.html",
		"http://yahoo.co.jp/foo/index.html",
	}
}

func TestInvalidBaseAndPathURL(t *testing.T) {
	for _, base := range badBases() {
		for _, path := range badPaths() {
			url, err := GetAbsoluteURLFromBaseAndPath(base, path)
			if err == nil {
				t.Errorf(errorUrlNotInvalid, url.String())
			}
		}
	}
}

func TestInvalidBaseAndValidPathURL(t *testing.T) {
	for _, base := range badBases() {
		for _, path := range okPaths() {
			url, err := GetAbsoluteURLFromBaseAndPath(base, path)
			if err == nil {
				t.Errorf(errorUrlNotInvalid, url.String())
			}
		}
	}
}

func TestValidBaseAndPathURL(t *testing.T) {
	re := regexp.MustCompile(`([^\:])[\/]{2,}`)
	for _, base := range okBases() {
		for _, path := range okPaths() {
			absUrl, err := GetAbsoluteURLFromBaseAndPath(base, path)
			if err != nil {
				t.Error(err)
			}
			expected := re.ReplaceAllString(base + "/" + path, `$1/`)
			if expected != absUrl.String() {
				t.Errorf(errorInvalidUrl, absUrl.String(), expected)
			}
		}
	}
}

func TestAbaPathURL(t *testing.T) {
	for _, base := range okBases() {
		for _, path := range absPaths() {
			url, err := GetAbsoluteURLFromBaseAndPath(base, path)
			if err != nil {
				t.Error(err)
			}
			if url.String() != path {
				t.Errorf(errorInvalidUrl, url.String(), path)
			}
		}
	}
}
