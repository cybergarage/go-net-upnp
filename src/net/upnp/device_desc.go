// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A Device represents a clinet.
type DeviceRoot struct {
	XMLName xml.Name `xml:"root"`
	URLBase string   `xml:"URLBase"`
	Device  Device   `xml:"device"`
}

// A Device represents a clinet.
type SpecVersion struct {
	XMLName xml.Name `xml:"specVersion"`
	Major   string   `xml:"major"`
	Minor   string   `xml:"minor"`
}
