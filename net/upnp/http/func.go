// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	gohttp "net/http"
)

func Get(url string) (resp *Response, err error) {
	res, err := gohttp.Get(url)
	if err != nil {
		return nil, err
	}
	return NewResponseFromResponse(res), nil
}
