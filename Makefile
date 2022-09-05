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
UPNPCTRL=${PREFIX}/bin/upnpctrl

LIGHTDEV=${PREFIX}/bin/lightdev
MEDIASERVER=${PREFIX}/bin/mediaserver

PKG_NAME=net/upnp
PKG_ID=${GITHUB}/${PKG_NAME}
PKGS=\
	${PKG_ID} \
	${PKG_ID}/ssdp \
	${PKG_ID}/control \
	${PKG_ID}/http
	
all: test

${VERSION_GO}: ./net/upnp/version.gen
	$< > $@

${USRAGNT_GO}: ./net/upnp/util/user_agent.gen ${VERSION_GO}
	$< > $@

version: ${VERSION_GO} ${USRAGNT_GO}

format:
	gofmt -s -w net example

vet: format
	go vet ${PKG_ID}

lint: vet
	golangci-lint run ${PKG_SRCS} ${BIN_SRCS} ${TEST_PKG_SRCS}

package: format $(shell find . -type f -name '*.go')
	go build -v ${PKGS}

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
	go test -v -cover ${PKGS}

install: build
	go install ${PKGS}

clean:
	rm ${PREFIX}/bin/*
	rm -rf _obj
	go clean -i ${PKGS}
