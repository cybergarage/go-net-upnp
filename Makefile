###################################################################
#
# go-net-upnp
#
# Copyright (C) The go-net-upnp Authors 2015
#
# This is licensed under BSD-style license, see file COPYING.
#
###################################################################

PREFIX?=$(shell pwd)
GOPATH=$(shell pwd)

VERSION_GO="./net/upnp/version.go"
USRAGNT_GO="./net/upnp/util/user_agent.go"

GITHUB=github.com/cybergarage/go-net-upnp

UPNPDUMP=${PREFIX}/bin/upnpdump
UPNPSEARCH=${PREFIX}/bin/upnpsearch
UPNPGWLIST=${PREFIX}/bin/upnpgwdump
LIGHTDEV=${PREFIX}/bin/lightdev
UPNPCTRL=${PREFIX}/bin/upnpctrl

PACKAGES=${GITHUB}/net/upnp ${GITHUB}/net/upnp/ssdp ${GITHUB}/net/upnp/control ${GITHUB}/net/upnp/log ${GITHUB}/net/upnp/http ${GITHUB}/net/upnp/http
	
all: build

${VERSION_GO}: ./net/upnp/version.gen
	$< > $@

${USRAGNT_GO}: ./net/upnp/util/user_agent.gen ${VERSION_GO}
	$< > $@

version: ${VERSION_GO} ${USRAGNT_GO}

goget:
	go get -u ${GITHUB}/{net/upnp/log,net/upnp/util,net/upnp/http,net/upnp/ssdp,net/upnp/control,net/upnp}

format:
	gofmt -w src net example

package: format $(shell find . -type f -name '*.go')
	go build -v ${PACKAGES}

${UPNPDUMP}: package $(shell find ./example/ctrlpoint/upnpdump -type f -name '*.go')
	go build -o $@ ./example/ctrlpoint/upnpdump

${UPNPSEARCH}: package $(shell find ./example/ctrlpoint/upnpsearch -type f -name '*.go')
	go build -o $@ ./example/ctrlpoint/upnpsearch

${UPNPGWLIST}: package $(shell find ./example/ctrlpoint/upnpgwlist -type f -name '*.go')
	go build -o $@ ./example/ctrlpoint/upnpgwlist

${LIGHTDEV}: package $(shell find ./example/device/light -type f -name '*.go')
	go build -o $@ ./example/device/light

${UPNPCTRL}: package $(shell find ./example/ctrlpoint/upnpctrl -type f -name '*.go')
	go build -o $@ ./example/ctrlpoint/upnpctrl

build: ${UPNPDUMP} ${UPNPSEARCH} ${UPNPGWLIST} ${LIGHTDEV} ${UPNPCTRL}

test: package
	go test -v -cover ${PACKAGES}

install: build
	go install ${PACKAGES}

clean:
	rm ${PREFIX}/bin/*
	rm -rf _obj
	go clean -i ${PACKAGES}
