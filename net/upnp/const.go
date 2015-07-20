// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"github.com/cybergarage/go-net-upnp/net/upnp/ssdp"
)

const (
	ProductName = "go-net-upnp"

	SupportVersion      = "1.1"
	SupportVersionMajor = 1
	SupportVersionMinor = 1

	ControlPointDefaultPortBase  = 5004
	ControlPointDefaultPortRange = 1024
	ControlPointDefaultPortMax   = ControlPointDefaultPortBase + ControlPointDefaultPortRange
	ControlPointDefaultSearchMX  = ssdp.DEFAULT_MSEARCH_MX

	DeviceDefaultPortBase  = 6004
	DeviceDefaultPortRange = 1024
	DeviceDefaultPortMax   = DeviceDefaultPortBase + DeviceDefaultPortRange
	DeviceUUIDPrefix       = "uuid:"

	DeviceProtocol              = "http"
	DeviceDefaultDescriptionURL = "/description.xml"

	In  = "in"
	Out = "out"

	UnknownDirection = 0
	InDirection      = 1
	OutDirection     = 2

	xmlMarshallIndent = " "
	xmlNs             = "xmlns"
)
