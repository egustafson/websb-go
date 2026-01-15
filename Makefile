##  -*- Makefile -*-
# --------------------------------------------------------------------

GIT_SUMMARY := $(shell git describe --tags --dirty --always)
BUILD_VER   := $(shell git describe --tags --always)
BUILD_DATE  := $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")

DIST_DIR = dist

PODMAN ?= podman

GO = go
GOFLAGS =

# --------------------------------------------------------------------

.PHONY: all        # default target
all: build test

.PHONY: preflight
preflight:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

.PHONY: build
build: websb

${DIST_DIR}:
	mkdir -p ${DIST_DIR}

.PHONY: websb  # force a rebuild always
websb: ${DIST_DIR}
	${GO} build -ldflags "-X main.GitSummary=$(GIT_SUMMARY) -X main.BuildDate=$(BUILD_DATE)" -o ${DIST_DIR}/$@

.PHONY: test
test: test-unit
#test: test-style

.PHONY: lint
lint: test-style

.PHONY: test-style
test-style:
	golangci-lint run ./...

.PHONY: test-unit
test-unit:
	go test ./...

.PHONY: package
package:
	${PODMAN} build -t websb:${BUILD_VER} .

.PHONY: clean
clean:
	go clean ./...
	rm -rf ${DIST_DIR}

.PHONY: real_clean
real_clean:
	go clean -cache
