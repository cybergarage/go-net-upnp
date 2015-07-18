// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A AllowedValueRange represents a icon.
// NOTE : Minimum, Maximum and Step uses string type instead float64
//        because some devices sets a blank into the fields such as BUFFALO WZR-900DHP
type AllowedValueRange struct {
	XMLName xml.Name `xml:"allowedValueRange"`
	Minimum string   `xml:"minimum"`
	Maximum string   `xml:"maximum"`
	Step    string   `xml:"step"`
}

// NewAllowedValueRange returns a new AllowedValueRange.
func NewAllowedValueRange() *AllowedValueRange {
	valRange := &AllowedValueRange{}
	return valRange
}
