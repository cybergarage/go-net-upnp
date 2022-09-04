// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"os"
	"testing"

	"github.com/cybergarage/go-net-upnp/net/upnp/log"
)

func TestMain(m *testing.M) {
	logger := log.NewStdoutLogger(log.LoggerLevelTrace)
	log.SetSharedLogger(logger)

	os.Exit(m.Run())
}
