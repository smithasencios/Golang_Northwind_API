# FROM golang:latest AS builder
FROM golang:1.10.1-alpine3.7 as builder

# Download and install the latest release of dep
RUN apk --update upgrade  && \
  apk add --no-cache ca-certificates openssh-client curl git bash

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR $GOPATH/src/github.com/Golang_Northwind_API

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app 

FROM scratch  
CMD ["./app"]
COPY --from=builder /app .