// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestNewAction(t *testing.T) {
	NewAction()
}

func TestMarshalAction(t *testing.T) {
	const nArgs = 5

	action := NewAction()
	action.Name = "Hello"
	action.ArgumentList.Arguments = make([]Argument, nArgs)
	for n := 0; n < nArgs; n++ {
		arg := NewArgument()
		arg.Name = fmt.Sprintf("name%d", n)
		arg.Value = fmt.Sprintf("value%d", n)
		action.ArgumentList.Arguments[n] = *arg
	}

	_, err := xml.MarshalIndent(action, "", "  ")
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(string(buf))
}
