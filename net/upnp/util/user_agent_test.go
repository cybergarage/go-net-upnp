// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"testing"
)

func TestUserAgent(t *testing.T) {
	agent := GetUserAgent()
	if len(agent) <= 0 {
		t.Errorf("user agent is null")
	}
}
