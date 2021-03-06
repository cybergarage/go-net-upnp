// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"os"
	"testing"

	"./log"
)

func TestMain(m *testing.M) {
	logger := log.NewStdoutLogger(log.LoggerLevelTrace)
	log.SetSharedLogger(logger)

	os.Exit(m.Run())
}
