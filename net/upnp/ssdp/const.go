// Copyright 2015 The go-net-upnp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

const (
	Port             = 1900
	ADDRESS          = "239.255.255.250"
	MulticastAddress = "239.255.255.250:1900"

	MulticastDefaultTTL = 4

	IPv6IfLocalAddress        = "FF01::C"
	IPv6LinkLocalAddress      = "FF02::C"
	IPv6SubnetAddress         = "FF03::C"
	IPv6AdministrativeAddress = "FF04::C"
	IPv6SiteLocalAddress      = "FF05::C"
	IPv6GlobalAddress         = "FF0E::C"

	DefaultMSearchMX     = 3
	DefaultAnnounceCount = 3

	MaxPacketSize     = 8192
	HeaderLineMaxSize = 128
	CRLF              = "\r\n"
	SP                = " "

	HTTPVersion = "1.1"
	HTTPPath    = "*"
	Notify      = "NOTIFY"
	MSearch     = "M-SEARCH"

	Host           = "HOST"
	Date           = "DATE"
	UserAgent      = "USER-AGENT"
	Location       = "LOCATION"
	Server         = "SERVER"
	ST             = "ST"
	MX             = "MX"
	MAN            = "MAN"
	NT             = "NT"
	NTS            = "NTS"
	NTSPropChange  = "upnp:propchange"
	USN            = "USN"
	EXT            = "EXT"
	SID            = "SID"
	SEQ            = "SEQ"
	Callback       = "CALLBACK"
	CacheControl   = "CACHE-CONTROL"
	DefaultTimeout = "DefaultTimeout"
	BOOTIDUPnPOrg  = "BOOTID.UPNP.ORG"

	RootDevice = "upnp:rootdevice"
	All        = "ssdp:all"
	Discover   = "\"ssdp:discover\""
	NTSAlive   = "ssdp:alive"
	NTSByeBye  = "ssdp:byebye"
	NTSUpdate  = "ssdp:update"
	MaxAge     = "max-age"
)
