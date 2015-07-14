// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import ()

// A Argument represents arguments in as SOAP action.
type Argument struct {
	Name  string
	Value string
}

// NewArgument returns a new argument.
func NewArgument() *Argument {
	arg := &Argument{}
	return arg
}
