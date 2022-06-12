# BUILD Stage
FROM golang:1.18-alpine as build
ENV CGO_ENABLED 1

RUN apk add git wget gcc g++

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go install github.com/go-delve/delve/cmd/dlv@v1.8.3

COPY . .
RUN go build -gcflags "all=-N -l" -o /server ./cmd/testc/main.go
ENTRYPOINT ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/server"]
