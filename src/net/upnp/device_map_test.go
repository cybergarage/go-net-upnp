// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"testing"
	"math/rand"
)

func TesNewDeviceMap(t *testing.T) {
	NewDeviceMap()
}

func TesDeviceMapCount(t *testing.T) {
	devMap := NewDeviceMap()

	const typeCnt = (rand.Int() % 10) + 10
	const udnCnt = (rand.Int() % 10) + 10
}
