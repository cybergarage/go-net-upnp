// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"errors"
	"testing"
)

const (
	nullControlPointError    = "ControlPoint is null"
	invalidControlPointError = "ControlPoint is invalid"
)

func TestNewControlPoint(t *testing.T) {
	cp := NewControlPoint()
	if cp == nil {
		t.Error(errors.New(nullControlPointError))
	}

	if len(cp.GetRootDevices()) != 0 {
		t.Error(errors.New(invalidControlPointError))
	}

	err := cp.Start()
	if err != nil {
		t.Error(err)
	}

	err = cp.SearchRootDevice()
	if err != nil {
		t.Error(err)
	}

	err = cp.Stop()
	if err != nil {
		t.Error(err)
	}
}
