// Copyright 2015 The go-net-upnp Authors. All rights reserved.
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
	errorURLNotAbsolute   = "url (%s) is not absolute"
	errorURLUnknownScheme = "url scheme (%s) is unknown"
)

func GetAbsoluteURLFromBaseAndPath(base string, path string) (*url.URL, error) {
	url, err := url.Parse(path)

	if err != nil || !url.IsAbs() {
		base = strings.TrimSuffix(base, urlDelim)
		path = strings.TrimPrefix(path, urlDelim)
		path = strings.TrimSuffix(path, urlDelim)

		urlStr := base + urlDelim + path
		url, err = url.Parse(urlStr)
		if err != nil {
			return nil, err
		}
	}

	if !url.IsAbs() {
		return nil, fmt.Errorf(errorURLNotAbsolute, url.String())
	}

	if (url.Scheme != "http") && (url.Scheme != "https") {
		return nil, fmt.Errorf(errorURLUnknownScheme, url.Scheme)
	}

	return url, nil
}
