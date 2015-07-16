// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package control

import (
	"encoding/xml"
)

const (
	XmlMarshallIndent = " "
	XmlNs             = "xmlns"
	XmlNsDelim        = ":"

	SoapActionSpace      = "u"
	SoapActionNamePrefix = SoapActionSpace + XmlNsDelim
	SoapActionNameSpace  = XmlNs + XmlNsDelim + SoapActionSpace

	XMLHeader = xml.Header
)
