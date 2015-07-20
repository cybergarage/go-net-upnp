// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"runtime"
)

func GetUserAgent() string {
	return fmt.Sprintf("%s/%s UPnP/%s %s/%s",
		runtime.GOOS,
		runtime.GOARCH,
		"1.1",
		"go-net-upnp",
		"0.8")
}

func GetServer() string {
	return GetUserAgent()
}
