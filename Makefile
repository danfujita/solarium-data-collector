# Makefile

export GOPATH := $(shell pwd)

build:
	go get $GOPATH/src/solarium-golang/cmd/solarium
	go build $GOPATH/src/solarium-golang/cmd/solarium/main.go
