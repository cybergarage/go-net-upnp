// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

const (
	PORT              = 1900
	ADDRESS           = "239.255.255.250"
	MULTICAST_ADDRESS = "239.255.255.250:1900"

	MULTICAST_DEFAULT_TTL = 4

	IPV6_IF_LOCAL_ADDRESS       = "FF01::C"
	IPV6_LINK_LOCAL_ADDRESS     = "FF02::C"
	IPV6_SUBNET_ADDRESS         = "FF03::C"
	IPV6_ADMINISTRATIVE_ADDRESS = "FF04::C"
	IPV6_SITE_LOCAL_ADDRESS     = "FF05::C"
	IPV6_GLOBAL_ADDRESS         = "FF0E::C"

	DEFAULT_MSEARCH_MX     = 3
	DEFAULT_ANNOUNCE_COUNT = 3

	MAX_PACKET_SIZE     = 8192
	HEADER_LINE_MAXSIZE = 128
	CRLF                = "\r\n"
	SP                  = " "

	HTTP_VERSION = "1.1"
	HTTP_PATH    = "*"
	NOTIFY       = "NOTIFY"
	M_SEARCH     = "M-SEARCH"

	HOST            = "HOST"
	DATE            = "DATE"
	USER_AGENT      = "USER-AGENT"
	LOCATION        = "LOCATION"
	SERVER          = "SERVER"
	ST              = "ST"
	MX              = "MX"
	MAN             = "MAN"
	NT              = "NT"
	NTS             = "NTS"
	NTS_PROPCHANGE  = "upnp:propchange"
	USN             = "USN"
	EXT             = "EXT"
	SID             = "SID"
	SEQ             = "SEQ"
	CALLBACK        = "CALLBACK"
	CACHE_CONTROL   = "CACHE-CONTROL"
	TIMEOUT         = "TIMEOUT"
	BOOTID_UPNP_ORG = "BOOTID.UPNP.ORG"

	ROOT_DEVICE = "upnp:rootdevice"
	ALL         = "ssdp:all"
	DISCOVER    = "\"ssdp:discover\""
	NTS_ALIVE   = "ssdp:alive"
	NTS_BYEBYE  = "ssdp:byebye"
	NTS_UPDATE  = "ssdp:update"
	MAX_AGE     = "max-age"
)
