// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"strings"
	"testing"
)

const (
	errorInvalidUUID = "invalide uuid = '%s'"
)

func TestCreateUUID(t *testing.T) {
	uuid := CreateUUID()

	if strings.Index(uuid, " ") != -1 {
		t.Errorf(errorInvalidUUID, uuid)
	}

	if len(uuid) != UUIDLength {
		t.Errorf(errorInvalidUUID, uuid)
	}
}
