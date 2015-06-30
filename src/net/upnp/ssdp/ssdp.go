// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

const (
	SSDP_PORT              = 1900
	SSDP_ADDRESS           = "239.255.255.250"
	SSDP_MULTICAST_ADDRESS = "239.255.255.250:1900"

	SSDP_MULTICAST_DEFAULT_TTL = 4

	SSDP_IPV6_IF_LOCAL_ADDRESS       = "FF01::C"
	SSDP_IPV6_LINK_LOCAL_ADDRESS     = "FF02::C"
	SSDP_IPV6_SUBNET_ADDRESS         = "FF03::C"
	SSDP_IPV6_ADMINISTRATIVE_ADDRESS = "FF04::C"
	SSDP_IPV6_SITE_LOCAL_ADDRESS     = "FF05::C"
	SSDP_IPV6_GLOBAL_ADDRESS         = "FF0E::C"

	SSDP_DEFAULT_MSEARCH_MX     = 3
	SSDP_DEFAULT_ANNOUNCE_COUNT = 3

	SSDP_HEADER_LINE_MAXSIZE = 128
	CRLF                     = "\r\n"

	SSDP_ST              = "ST"
	SSDP_MX              = "MX"
	SSDP_MAN             = "MAN"
	SSDP_NT              = "NT"
	SSDP_NTS             = "NTS"
	SSDP_NTS_ALIVE       = "ssdp:alive"
	SSDP_NTS_BYEBYE      = "ssdp:byebye"
	SSDP_NTS_PROPCHANGE  = "upnp:propchange"
	SSDP_USN             = "USN"
	SSDP_EXT             = "EXT"
	SSDP_SID             = "SID"
	SSDP_SEQ             = "SEQ"
	SSDP_CALBACK         = "CALLBACK"
	SSDP_TIMEOUT         = "TIMEOUT"
	SSDP_SERVER          = "SERVER"
	SSDP_BOOTID_UPNP_ORG = "BOOTID.UPNP.ORG"
)
