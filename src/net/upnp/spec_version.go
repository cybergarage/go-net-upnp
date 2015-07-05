// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A SpecVersion represents a UPnP spec version.
type SpecVersion struct {
	XMLName xml.Name `xml:"specVersion"`
	Major   int      `xml:"major"`
	Minor   int      `xml:"minor"`
}

// NewSpecVersion returns a new SpecVersion.
func NewSpecVersion() *SpecVersion {
	spec := &SpecVersion{Major: SUPPORT_VERSION_MAJOR, Minor: SUPPORT_VERSION_MINOR}
	return spec
}
