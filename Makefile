EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

BINARY=AppKasir
VERSION=1.0.0
BUILD=`git rev-parse HEAD`
PLATFORMS=linux windows
ARCHITECTURES=386 amd64

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

default: build
all: clean build_all install

build:
	go1.20.12 build ${LDFLAGS} -o ./out/${BINARY}

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); export CGO_ENABLED=1; go1.20.12 build -v -o ./out/$(BINARY)-$(GOOS)-$(GOARCH))))

install:
	go1.20.12 install ${LDFLAGS}

clean:
	$(shell rm -rf ${ROOT_DIR}/out)