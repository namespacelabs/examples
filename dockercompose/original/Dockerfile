# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-server

EXPOSE 5000
ENV HTTP_PORT 5000

CMD [ "/go-server" ]