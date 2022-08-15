FROM golang:1.18-alpine as builder

RUN apk add --no-cache git gcc libc-dev curl
RUN mkdir /go-plugins
COPY . /go-plugins

RUN cd /go-plugins && \
   go mod download && \
   go build kong-go-plugin-GeoIP.go

FROM kong:2.8
COPY --from=builder /go-plugins/kong-go-plugin-GeoIP /usr/local/bin/