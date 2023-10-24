FROM golang:1.20-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY ./src .

COPY go.mod .
COPY go.sum .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o ./go-api

EXPOSE ${PORT}

CMD ["./go-api"]