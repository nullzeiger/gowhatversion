# Copyright 2022 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
# All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

BINARY_NAME=gowhatversion
BINARY_PATH=build

all: fmt vet test lint build

.PHONY: lint
lint:
	golangci-lint run --config .golangci.yml

.PHONY: run
run:
	go run ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	gofmt -d -e -s -w .

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build -o ${BINARY_PATH}/${BINARY_NAME} cmd/gowhatversion/main.go

clean:
	go clean
	rm -r ${BINARY_PATH}/${BINARY_NAME}
