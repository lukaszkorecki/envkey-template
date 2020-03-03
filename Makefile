deps:
	go get -v github.com/envkey/envkeygo/loader

build:
	go build ./...


fmt:
	go fmt


all: deps fmt build
