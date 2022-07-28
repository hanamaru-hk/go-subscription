FROM golang:1.16-alpine

WORKDIR /app

RUN apk update && apk add bash

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY .env ./

RUN go build -o go-subscription

ENTRYPOINT ["/app/go-subscription"]

EXPOSE 8090