FROM golang:1.20

WORKDIR /app

COPY api/ ./

RUN go mod download

RUN CGO_ENABLED=0 go build -o ./go-api

EXPOSE 8080

CMD ["./go-api"]