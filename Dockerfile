FROM golang:1.21

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o main ./cmd/server

ENV GIN_MODE=release
CMD ["./main"]
