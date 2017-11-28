SHELL := /bin/bash
MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1.0")
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
LDFLAGS:=-ldflags "-X github.com/jkomyno/palitrux/config.Version=${VERSION}"

install:
	dep ensure

clean:
	rm -rf palitrux palitrux.exe

test:
	go test $$(go list ./... | grep -v /vendor/)

build: clean
	go build ${LDFLAGS} -a -o bin/palitrux
