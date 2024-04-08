# Makefile

.PHONY: all build run clean test

BINARY_NAME=app01
MAIN_FILE=app01.go

build-t:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

clean:
	go clean
	rm -f $(BINARY_NAME)

test:
	go test ./...

build-main:
	@sh script/build.sh

run: build-main
	cd bin && ./out-darwin-arm64

all: run