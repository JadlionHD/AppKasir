EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

BINARY=AppKasir
VERSION=1.0.0
BUILD=`git rev-parse HEAD`
PLATFORMS=linux windows
ARCHITECTURES=386 amd64
GO_EXEC=go1.20.12

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

default: build
all: clean build_all install

build:
	${GO_EXEC} build ${LDFLAGS} -o ./out/${BINARY}

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); export CGO_ENABLED=1; ${GO_EXEC} build -o ./out/$(BINARY)-$(GOOS)-$(GOARCH))))

install:
	${GO_EXEC} install ${LDFLAGS}

clean:
	$(shell rm -rf ${ROOT_DIR}/out)