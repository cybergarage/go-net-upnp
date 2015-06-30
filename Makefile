###################################################################
#
# net/upnp
#
# Copyright (C) The go-net-upnp Authors 2015
#
# This is licensed under BSD-style license, see file COPYING.
#
###################################################################

packages = net/upnp net/upnp/ssdp
	
all: build

format:
	gofmt -w src

build: format 
	go build -v ${packages}

test: build
	go test -v ${packages}

install: build
	go install ${packages}

clean:
	rm -rf _obj
	go clean -i ${packages}
