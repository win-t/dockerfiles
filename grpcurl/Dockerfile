# syntax=docker/dockerfile:1

FROM golang:1-alpine AS builder
ENV GOCACHE=/gocache/build
ENV GOMODCACHE=/gocache/mod
WORKDIR /builder
COPY . .
RUN --mount=type=cache,target=/gocache GOBIN=$PWD go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

FROM alpine
RUN apk -u add curl openssl bind-tools
COPY --from=builder /builder/grpcurl /bin/grpcurl
