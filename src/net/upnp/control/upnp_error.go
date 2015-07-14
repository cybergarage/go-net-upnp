// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
	"fmt"
)

const (
	upnpErrorFormat = "[%d] %s"
)

type UPnPError struct {
	XMLName          xml.Name `xml:"urn:schemas-upnp-org:control-1-0 UPnPError"`
	ErrorCode        int      `xml:"errorCode"`
	ErrorDescription string   `xml:"errorDescription"`
}

func (self UPnPError) Error() string {
	return fmt.Sprintf(upnpErrorFormat, self.ErrorCode, self.ErrorDescription)
}
