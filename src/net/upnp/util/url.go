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
	errorUrlNotAbsolute = "url (%s) is not absolute"
)

func GetAbsoluteURLFromBaseAndPath(base string, path string) (*url.URL, error) {
	base = strings.TrimSuffix(base, urlDelim)
	path = strings.TrimSuffix(path, urlDelim)

	urlStr := base + urlDelim + path
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if !url.IsAbs() {
		return nil, fmt.Errorf(errorUrlNotAbsolute, url.String())
	}

	return url, nil
}
