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
UPNPGWDUMP=${PREFIX}/bin/upnpgwdump
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

${UPNPGWDUMP}: package $(shell find ./example/ctrlpoint/upnpgwdump -type f -name '*.go')
	go build -o $@ ./example/ctrlpoint/upnpgwdump

${LIGHTDEV}: package $(shell find ./example/device/light -type f -name '*.go')
	go build -o $@ ./example/device/light

${UPNPCTRL}: package $(shell find ./example/ctrlpoint/upnpctrl -type f -name '*.go')
	go build -o $@ ./example/ctrlpoint/upnpctrl

build: ${UPNPDUMP} ${UPNPGWDUMP} ${LIGHTDEV} ${UPNPCTRL}

test: package
	go test -v -cover ${PACKAGES}

install: build
	go install ${PACKAGES}

clean:
	rm ${PREFIX}/bin/*
	rm -rf _obj
	go clean -i ${PACKAGES}
