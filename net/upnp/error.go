// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

// A Error represents a interface for UPnP error.
type Error interface {
	error
	GetCode() int
	GetDescription() string
}
