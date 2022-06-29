FROM golang:1.17-alpine AS build-env

WORKDIR /go/src

ADD . /go/src

RUN apk update && apk add git \
    && cd /go/src \
    && go install github.com/swaggo/swag/cmd/swag@v1.7.9 \
    && swag init -g ./cmd/main.go \
    && go build -o ./build/bin/main ./cmd/main.go

FROM alpine:3.15

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /go/src
COPY --from=build-env /go/src/build/bin/main /go/src

EXPOSE 8080

CMD ["./main"]
