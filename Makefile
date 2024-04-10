SHELL=/usr/bin/env bash

.PHONY: build
build: build
	go mod tidy
	rm -rf nodedao-tool
	go build -o nodedao-tool main.go
