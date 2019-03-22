# Builder
FROM golang:alpine AS build-env

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git
COPY . /go/src/github.com/Footters/hex-footters
WORKDIR /go/src/github.com/Footters/hex-footters
RUN GOOS=linux go build -o hex cmd/main.go

# Exec 
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/Footters/hex-footters /app/
EXPOSE 3000
ENTRYPOINT ./hex