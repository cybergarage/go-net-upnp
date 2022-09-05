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

BIN_ROOT=examples
BINS=\
	${GITHUB}/${BIN_ROOT}/ctrlpoint/upnpdump \
	${GITHUB}/${BIN_ROOT}/ctrlpoint/upnpsearch \
	${GITHUB}/${BIN_ROOT}/ctrlpoint/upnpgwlist \
	${GITHUB}/${BIN_ROOT}/ctrlpoint/upnpctrl \
	${GITHUB}/${BIN_ROOT}/device/upnplight \
	${GITHUB}/${BIN_ROOT}/device/upnpavserver

version: ${VERSION_GO} ${USRAGNT_GO}

format:
	gofmt -s -w ${PKG_NAME} ${BIN_ROOT}

vet: format
	go vet ${PKG_ID}

lint: vet
	golangci-lint run ${PKG_SRCS} ${BIN_SRCS} ${TEST_PKG_SRCS}

build:
	go build -v ${PKGS}

test: build
	go test -v -cover ${PKGS}

install: test
	go install ${BINS}

clean:
	go clean -i ${PKGS}
