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

VERSION_GO="./src/net/upnp/version.go"
USRAGNT_GO="./src/net/upnp/util/user_agent.go"

UPNPDUMP=${PREFIX}/bin/upnpdump
UPNPGWDUMP=${PREFIX}/bin/upnpgwdump
LIGHTDEV=${PREFIX}/bin/lightdev
UPNPCTRL=${PREFIX}/bin/upnpctrl

packages = net/upnp net/upnp/log net/upnp/ssdp net/upnp/util net/upnp/control
	
# .PHONY: ./src/net/upnp/version.go ./src/net/upnp/util/user_agent.go

all: build

${VERSION_GO}: ./src/net/upnp/version.gen
	$< > $@

${USRAGNT_GO}: ./src/net/upnp/util/user_agent.gen ${VERSION_GO}
	$< > $@

version: ${VERSION_GO} ${USRAGNT_GO}

format:
	gofmt -w src example

package: format $(shell find . -type f -name '*.go')
	go build -v ${packages}

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
	go test -v -cover ${packages}

install: build
	go install ${packages}

clean:
	rm ${PREFIX}/bin/upnpdump
	rm -rf _obj
	go clean -i ${packages}
