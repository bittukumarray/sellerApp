FROM golang:1.16.3-alpine
RUN apk add git
RUN apk add build-base
WORKDIR /fetcher
COPY . /fetcher
WORKDIR ./fetcher
RUN go mod download
ENTRYPOINT go build && go run main.go