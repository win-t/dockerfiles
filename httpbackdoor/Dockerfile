# syntax=docker/dockerfile:1

FROM golang:1 AS builder
ENV GOCACHE=/gocache/build
ENV GOMODCACHE=/gocache/mod
WORKDIR /builder
COPY . .
RUN --mount=type=cache,target=/gocache  go build -o server .

FROM ubuntu:latest
RUN export DEBIAN_FRONTEND=noninteractive && apt update && apt install -y curl bind9-dnsutils iproute2 iputils-ping netcat-openbsd busybox
COPY --from=builder /builder/server /bin/server
ENTRYPOINT ["server"]
