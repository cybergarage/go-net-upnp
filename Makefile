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

VERSION_GO="./net/upnp/version.go"
USRAGNT_GO="./net/upnp/util/user_agent.go"

MODULE_ROOT=github.com/cybergarage/go-net-upnp

PKG_NAME=net/upnp
PKG_VER=$(shell git describe --abbrev=0 --tags)
PKG_COVER=net-upnp-cover
PKG_ID=${MODULE_ROOT}/${PKG_NAME}
PKG_SRC_DIR=${PKG_NAME}
PKG=${MODULE_ROOT}/${PKG_SRC_DIR}
	
all: test

${VERSION_GO}: ./net/upnp/version.gen
	$< > $@

${USRAGNT_GO}: ./net/upnp/util/user_agent.gen ${VERSION_GO}
	$< > $@

BIN_ROOT=examples
BIN_ID=${MODULE_ROOT}/${BIN_ROOT_DIR}
BIN_SRCS=\
	${BIN_ROOT}/ctrlpoint/upnpdump \
	${BIN_ROOT}/ctrlpoint/upnpsearch \
	${BIN_ROOT}/ctrlpoint/upnpgwlist \
	${BIN_ROOT}/ctrlpoint/upnpctrl \
	${BIN_ROOT}/device/upnplight \
	${BIN_ROOT}/device/upnpavserver
BINS=\
	${BIN_ID}/ctrlpoint/upnpdump \
	${BIN_ID}/ctrlpoint/upnpsearch \
	${BIN_ID}/ctrlpoint/upnpgwlist \
	${BIN_ID}/ctrlpoint/upnpctrl \
	${BIN_ID}/device/upnplight \
	${BIN_ID}/device/upnpavserver

version: ${VERSION_GO} ${USRAGNT_GO}

format:
	gofmt -s -w ${PKG_NAME} ${BIN_ROOT}

vet: format
	go vet ${PKG_ID}

lint: vet
	golangci-lint run --skip-files .*_test.go ${PKG_SRC_DIR}/... ${BIN_ROOT}/...

build: lint
	go build -v ${PKG}

test:
	go test -v -p 1 -DefaultTimeout 10m -cover -coverpkg=${PKG}/... -coverprofile=${PKG_COVER}.out ${PKG}/...
	go tool cover -html=${PKG_COVER}.out -o ${PKG_COVER}.html

install: test
	go install ${BINS}

clean:
	go clean -i ${PKGS}
