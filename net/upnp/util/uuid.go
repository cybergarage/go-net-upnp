// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	uuidLength = (8 + 1 + 4 + 1 + 4 + 1 + 4 + 1 + 12)
)

type uuid struct {
	TimeLow       uint32
	TimeMid       uint16
	TimeHiVer     uint16
	ClockSeqHiRes byte
	ClockSeqLow   byte
	Node          []byte
}

// CreateUUID returns a UUID
// Note : the functions is implemented aboutly to the following specfication.
// https://www.ietf.org/rfc/rfc4122.txt
func CreateUUID() string {

	var uuid uuid

	/* UUIDs use time in 100ns ticks since Oct 15, 1582. */
	uuidStartTime := time.Date(1582, 8, 25, 0, 0, 0, 0, time.Local)
	uuidTs := time.Since(uuidStartTime).Nanoseconds() / 100
	now := time.Now()
	clockSeq := now.Nanosecond() // a random sequence number

	uuid.TimeLow = uint32(uuidTs & 0xFFFFFFFF)
	uuid.TimeMid = uint16((uuidTs >> 32) & 0xFFFF)
	uuid.TimeHiVer = uint16((uuidTs >> 48) & 0x0FFF)
	uuid.TimeHiVer |= (1 << 12)
	uuid.ClockSeqLow = byte(clockSeq & 0xFF)
	uuid.ClockSeqHiRes = byte((clockSeq & 0x3F00) >> 8)
	uuid.ClockSeqHiRes |= 0x80

	uuid.Node = make([]byte, 6)

	// random initializer
	for n := 0; n < len(uuid.Node); n++ {
		uuid.Node[n] = byte(rand.Int() & 0xFF)
	}
	uuid.Node[0] |= 0x01

	ifis, err := GetAvailableInterfaces()
	if (err != nil) && (0 < len(ifis)) {
		ifi := ifis[0]
		for n := 0; (n < len(uuid.Node)) && (n < len(ifi.HardwareAddr)); n++ {
			uuid.Node[n] = ifi.HardwareAddr[n]
		}
	}

	uuidStr := fmt.Sprintf("%08x-%04x-%04x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		uuid.TimeLow,
		uuid.TimeMid,
		uuid.TimeHiVer,
		uuid.ClockSeqHiRes,
		uuid.ClockSeqLow,
		uint(uuid.Node[0]),
		uint(uuid.Node[1]),
		uint(uuid.Node[2]),
		uint(uuid.Node[3]),
		uint(uuid.Node[4]),
		uint(uuid.Node[5]),
	)

	return uuidStr
}
