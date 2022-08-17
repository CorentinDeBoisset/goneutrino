# If version is undefined, we use "NoVersion as a default"
ifeq ($(VERSION),)
	VERSION = NoVersion
endif

PACKAGE_NAME := $(shell go list -m -f '{{.Path}}')
BIN_DIR := bin
DEB_FILES_DIR := deb_files
GO_FILES := $(shell find . -type f -name '*.go')

EXECUTABLE := neutrino


.PHONY: all
all: ${BIN_DIR}/${EXECUTABLE}

.PHONY: prod_build
prod_build: ${BIN_DIR}/${EXECUTABLE}_prod

.PHONY: clean
clean:
	rm -f ${BIN_DIR}/${EXECUTABLE} ${BIN_DIR}/${EXECUTABLE}_prod

# Build the executable
${BIN_DIR}/${EXECUTABLE}: ${GO_FILES}
	CGO_ENABLED=0 go build -ldflags "-X main.ExecutableName=${EXECUTABLE} -X main.Version=${VERSION}" -o $@ ./cmd/neutrino

${BIN_DIR}/${EXECUTABLE}_prod: ${GO_FILES}
	CGO_ENABLED=0 go build -tags prod -ldflags "-X main.ExecutableName=${EXECUTABLE} -X main.Version=${VERSION}" -o $@ ./cmd/neutrino
