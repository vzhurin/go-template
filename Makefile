# TODO linter, build

default: build

HASH := $(shell git rev-parse --short=10 HEAD)

.PHONY: build
build:
	echo "build"