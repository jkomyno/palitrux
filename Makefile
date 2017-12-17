SHELL:=/bin/bash
GOOS:=linux
GOARCH:=amd64
APP:=palitrux
PROJECT:=github.com/jkomyno/${APP}
MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1.0")
VERSION:=${MAIN_VERSION}
BUILD_DATE:=$(shell date -u '+%Y-%m-%d')
LDFLAGS:=-ldflags "-s -w -X ${PROJECT}/config.Version=${VERSION} \
									 -X ${PROJECT}/config.BuildDate=${BUILD_DATE}"

install:
	dep ensure

clean:
	rm -rf palitrux palitrux.exe

test:
	go test $$(go list ./... | grep -v /vendor/)

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build ${LDFLAGS} -a -o ./${APP}

container: build
	docker build -t $(APP):$(VERSION) .

run: container
	docker stop $(APP):$(VERSION) || true && \
	docker rm $(APP):$(VERSION) || true && \
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
						 -e "PORT=${PORT}" \
						 $(APP):$(VERSION)