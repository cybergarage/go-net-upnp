// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

import (
	"encoding/xml"
)

// A AllowedValue represents a icon.
type AllowedValue struct {
	XMLName xml.Name `xml:"allowedValue"`
	Value   string   `xml:",innerxml"`
}

// NewAllowedValue returns a new AllowedValue.
func NewAllowedValue() *AllowedValue {
	value := &AllowedValue{}
	return value
}
