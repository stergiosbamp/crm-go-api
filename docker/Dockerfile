FROM golang:1.20-alpine

WORKDIR /app

COPY ./api .

RUN go mod download

RUN CGO_ENABLED=0 go build -o ./go-api

EXPOSE ${PORT}

CMD ["./go-api"]