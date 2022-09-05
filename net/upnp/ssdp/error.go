// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

const (
	errorZeroPacket              = "packet length is zero"
	errorPacketFirstLineNotFound = "first line is not found\n%s"
	errorPacketHeadersNotFound   = "headers is not found (%d:%d)\n%s"
)
