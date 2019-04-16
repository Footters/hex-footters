# Builder
FROM golang:alpine AS build-env

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git
WORKDIR /go/src/github.com/Footters/hex-footters
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
# RUN CGO_ENABLED=0 go test ./pkg/auth/auth_test/
RUN GOOS=linux go build -o auth cmd/auth/main.go

# Exec 
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/Footters/hex-footters /app/
EXPOSE 8081
ENTRYPOINT ./auth