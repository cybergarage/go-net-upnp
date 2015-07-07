// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"runtime"
)

func GetUserAgent() string {
	return fmt.Sprintf("%s/%s UPnP/1.1 go-net-upnp/%s",
		runtime.GOOS,
		runtime.GOARCH,
		"go-net-upnp",
		"1.1",
		"0.1-4-g23a1766")
}
