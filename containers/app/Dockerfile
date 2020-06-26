FROM golang:1.14.3

WORKDIR /app
ENV CGO_ENABLED 0

RUN go get -u golang.org/x/lint/golint

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download -x

COPY . /app
