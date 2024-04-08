# Makefile

.PHONY: all build run clean test

APP=app01
MAIN_FILE=app01.go

clean:
	go clean
	rm -rf bin/*

test:
	go test ./...

build-main:
	@sh script/build.sh

run: clean build-main
	bin/${APP}-darwin-arm64

all: run