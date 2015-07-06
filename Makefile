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

packages = net/upnp net/upnp/log net/upnp/ssdp net/upnp/util
	
.PHONY: ./src/net/upnp/version.go ./src/net/upnp/util/user_agent.go

all: build

./src/net/upnp/version.go: ./src/net/upnp/version.gen
	$< > $@

./src/net/upnp/util/user_agent.go: ./src/net/upnp/util/user_agent.gen
	$< > $@

versions: ./src/net/upnp/version.go ./src/net/upnp/util/user_agent.go

format: versions
	gofmt -w src cmd

package: format $(shell find . -type f -name '*.go')
	go build -v ${packages}

${PREFIX}/bin/upnpdump: package $(shell find ./cmd/upnpdump -type f -name '*.go')
	go build -o $@ ./cmd/upnpdump

build: ${PREFIX}/bin/upnpdump

test: package
	go test -v ${packages}

install: build
	go install ${packages}

clean:
	rm ${PREFIX}/bin/upnpdump
	rm -rf _obj
	go clean -i ${packages}
