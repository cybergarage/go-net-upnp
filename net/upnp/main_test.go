// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"os"
	"testing"

	"github.com/cybergarage/go-logger/log"
)

func TestMain(m *testing.M) {
	log.SetDefault(log.NewStdoutLogger(log.LevelTrace))
	os.Exit(m.Run())
}
