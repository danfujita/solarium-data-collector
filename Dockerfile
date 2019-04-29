FROM golang:1.12.4

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR $GOPATH/src/solarium-data-collector

COPY . .
#COPY ["cmd", "configs", "internal", "vendor", "web", "./"]

RUN dep init

RUN go build cmd/solarium/main.go
EXPOSE 8080
CMD ["./main"]
