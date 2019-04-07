# Builder
FROM golang:alpine AS build-env

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git
COPY . /go/src/github.com/Footters/hex-footters
WORKDIR /go/src/github.com/Footters/hex-footters
RUN GOOS=linux go build -o auth cmd/auth/main.go

# Exec 
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/Footters/hex-footters /app/
EXPOSE 8081
ENTRYPOINT ./auth