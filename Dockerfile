# Builder
FROM golang:alpine AS build-env
RUN apk update && apk add --no-cache git \
    && go get -u github.com/golang/dep/cmd/dep
COPY . /go/src/github.com/Footters/hex-footters
WORKDIR /go/src/github.com/Footters/hex-footters
RUN go get
RUN GOOS=linux go build -o hex

# Exec 
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/Footters/hex-footters /app/
EXPOSE 3000
ENTRYPOINT ./hex