// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A AllowedValueRange represents a icon.
type AllowedValueRange struct {
	XMLName xml.Name `xml:"allowedValueRange"`
	Minimum float64  `xml:"minimum"`
	Maximum float64  `xml:"maximum"`
	Step    float64  `xml:"step"`
}

// NewAllowedValueRange returns a new AllowedValueRange.
func NewAllowedValueRange() *AllowedValueRange {
	valRange := &AllowedValueRange{}
	return valRange
}
