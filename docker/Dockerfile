FROM golang:1.20-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY src src
COPY cmd cmd

RUN cd cmd && go mod tidy && CGO_ENABLED=0 go build -o bin/crm-go-api

EXPOSE ${PORT}

ENTRYPOINT ["./cmd/bin/crm-go-api"]