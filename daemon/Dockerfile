FROM golang:1.17 AS builder

COPY . /go/src/github.com/xBalbinus/polygonburn-watcher/daemon
WORKDIR /go/src/github.com/xBalbinus/polygonburn-watcher/daemon
RUN go build

FROM debian:buster-slim

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/xBalbinus/polygonburn-watcher/daemon/daemon daemon

ENTRYPOINT ["/daemon"]
