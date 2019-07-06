# FROM golang:latest AS builder
FROM golang:1.10.1-alpine3.7 as builder

WORKDIR $GOPATH/src/github.com/Golang_Northwind_API
COPY . .
RUN apk --update upgrade  && \
  apk add --no-cache ca-certificates openssh-client curl git bash

RUN apk add --no-cache curl git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app 

FROM alpine:3.7  
CMD ["./app"]
COPY --from=builder /app .