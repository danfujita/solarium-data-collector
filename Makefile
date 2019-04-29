# Makefile

build:
	go get -u github.com/golang/dep/cmd/dep
	dep init
	go build $(GOPATH)/src/solarium-data-collector/cmd/solarium/main.go
