// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	urlDelim = "/"
)

const (
	errorUrlNotAbsolute   = "url (%s) is not absolute"
	errorUrlUnknownScheme = "url scheme (%s) is unknown"
)

func GetAbsoluteURLFromBaseAndPath(base string, path string) (*url.URL, error) {
	url, err := url.Parse(path)

	if err != nil || !url.IsAbs() {
		base = strings.TrimSuffix(base, urlDelim)
		path = strings.TrimSuffix(path, urlDelim)

		urlStr := base + urlDelim + path
		url, err = url.Parse(urlStr)
		if err != nil {
			return nil, err
		}
	}

	if !url.IsAbs() {
		return nil, fmt.Errorf(errorUrlNotAbsolute, url.String())
	}

	if (url.Scheme != "http") && (url.Scheme != "https") {
		return nil, fmt.Errorf(errorUrlUnknownScheme, url.Scheme)
	}

	return url, nil
}
