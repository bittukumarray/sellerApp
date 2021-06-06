FROM golang:1.16.3-alpine
RUN apk add git
RUN apk add build-base
WORKDIR /writer
COPY . /writer
WORKDIR ./writer
RUN go mod download
ENTRYPOINT go build && go run main.go