# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

ENV GOCACHE=/gocache/build
ENV GOMODCACHE=/gocache/mod

RUN --mount=type=cache,target=/cache \
    --mount=type=bind,dst=/build,rw \
    cd /build && go build -o /out/server .

FROM alpine
COPY --from=builder /out/server /server
ENTRYPOINT ["/server"]
