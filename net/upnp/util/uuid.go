// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"github.com/google/uuid"
)

// CreateUUID returns a UUID.
func CreateUUID() string {
	return uuid.New().String()
}
