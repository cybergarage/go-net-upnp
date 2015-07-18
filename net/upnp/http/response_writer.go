// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	gohttp "net/http"
)

// A ResponseWriter represents a ResponseWriter.
type ResponseWriter interface {
	gohttp.ResponseWriter
}
