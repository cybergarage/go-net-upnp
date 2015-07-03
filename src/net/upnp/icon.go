// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A Icon represents a UPnP icon.
type Icon struct {
	XMLName  xml.Name `xml:"icon"`
	Mimetype string   `xml:"mimetype"`
	Width    string   `xml:"width"`
	Height   string   `xml:"height"`
	Depth    string   `xml:"depth"`
	Url      string   `xml:"url"`
}

// A IconList represents a UPnP iconList.
type IconList struct {
	XMLName xml.Name `xml:"iconList"`
	Icons   []Icon   `xml:"icon"`
}

// NewIcon returns a new Icon.
func NewIcon() *Icon {
	icon := &Icon{}
	return icon
}
