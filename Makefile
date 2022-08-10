# If version is undefined, we use "NoVersion as a default"
ifeq ($(VERSION),)
	VERSION = NoVersion
endif

PACKAGE_NAME := $(shell go list -m -f '{{.Path}}')
BIN_DIR := bin
DEB_FILES_DIR := deb_files
GO_FILES := $(shell find . -type f -name '*.go')
STATIC_ASSETS := $(shell find web -type f)

EXECUTABLE := secret-send


.PHONY: all
all: ${BIN_DIR}/${EXECUTABLE}

.PHONY: clean
clean:
	rm -f ${BIN_DIR}/${EXECUTABLE}
	rm -rf pkg/server/static/bindata.go

pkg/server/static/bindata.go: ${STATIC_ASSETS}
	go-bindata -fs -pkg static -prefix "web" -o $@ web/...

# Build the executable
${BIN_DIR}/${EXECUTABLE}: ${GO_FILES} pkg/server/static/bindata.go
	CGO_ENABLED=0 go build -ldflags "-X main.ExecutableName=${EXECUTABLE} -X main.Version=${VERSION}" -o $@ ./cmd/secret-send
