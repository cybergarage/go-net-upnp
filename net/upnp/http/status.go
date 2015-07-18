// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	gohttp "net/http"
)

const (
	StatusOK                  = gohttp.StatusOK
	StatusBadRequest          = gohttp.StatusBadRequest
	StatusNotFound            = gohttp.StatusNotFound
	StatusPreconditionFailed  = gohttp.StatusPreconditionFailed
	StatusInternalServerError = gohttp.StatusInternalServerError
)

func StatusCodeToString(code int) string {
	return gohttp.StatusText(code)
}
