// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A ServiceDescription represents a UPnP service description.
type ServiceDescription struct {
	XMLName           xml.Name          `xml:"scpd"`
	ServiceStateTable ServiceStateTable `xml:"serviceStateTable"`
	ActionList        ActionList        `xml:"actionList"`
}

// A ServiceList represents a UPnP serviceList.
type ServiceList struct {
	XMLName  xml.Name  `xml:"serviceList"`
	Services []Service `xml:"service"`
}
