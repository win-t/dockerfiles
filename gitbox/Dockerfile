# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

ENV GOCACHE=/gocache/build
ENV GOMODCACHE=/gocache/mod

RUN --mount=type=cache,target=/cache \
    --mount=type=bind,dst=/build,rw \
    cd /build && go build -o /out/server .

FROM alpine
RUN apk -U add git-daemon
COPY --from=builder /out/server /server
WORKDIR /data
VOLUME /data
ENV GIT_PROJECT_ROOT=/data
ENTRYPOINT ["/server"]
