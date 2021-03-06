###################################################################
#
# go-net-upnp
#
# Copyright (C) Satoshi Konno 2015
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
UPNPCTRL=${PREFIX}/bin/upnpctrl

LIGHTDEV=${PREFIX}/bin/lightdev
MEDIASERVER=${PREFIX}/bin/mediaserver

PACKAGES=${GITHUB}/net/upnp ${GITHUB}/net/upnp/ssdp ${GITHUB}/net/upnp/control ${GITHUB}/net/upnp/log ${GITHUB}/net/upnp/http ${GITHUB}/net/upnp/http
	
all: build

${VERSION_GO}: ./net/upnp/version.gen
	$< > $@

${USRAGNT_GO}: ./net/upnp/util/user_agent.gen ${VERSION_GO}
	$< > $@

version: ${VERSION_GO} ${USRAGNT_GO}

setup:
	go get -u ${GITHUB}/net/upnp

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

${UPNPCTRL}: package $(shell find ./example/ctrlpoint/upnpctrl -type f -name '*.go')
	go build -o $@ ./example/ctrlpoint/upnpctrl

${LIGHTDEV}: package $(shell find ./example/dev/lightdev -type f -name '*.go')
	go build -o $@ ./example/dev/lightdev

${MEDIASERVER}: package $(shell find ./example/dev/mediaserver -type f -name '*.go')
	go build -o $@ ./example/dev/mediaserver

build: ${UPNPDUMP} ${UPNPSEARCH} ${UPNPGWLIST} ${LIGHTDEV} ${UPNPCTRL} ${MEDIASERVER}

test: package
	go test -v -cover ${PACKAGES}

install: build
	go install ${PACKAGES}

clean:
	rm ${PREFIX}/bin/*
	rm -rf _obj
	go clean -i ${PACKAGES}
