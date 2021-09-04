ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BACKEND_LIBS_PATH := $(ROOT_DIR)/libs

install:
	GOMODCACHE=$(BACKEND_LIBS_PATH) go mod download

test:
	GOMODCACHE=$(BACKEND_LIBS_PATH) go test -cover $(ROOT_DIR)

fmt:
	GOMODCACHE=$(BACKEND_LIBS_PATH) go fmt $(ROOT_DIR)

check:
	GOMODCACHE=$(BACKEND_LIBS_PATH) go vet $(ROOT_DIR)
