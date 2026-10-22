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
COPY --from=wint/pid1 /docker_pid1 /usr/local/bin/docker_pid1
WORKDIR /data
VOLUME /data
ENV GIT_PROJECT_ROOT=/data
ENTRYPOINT ["docker_pid1", "/server"]
