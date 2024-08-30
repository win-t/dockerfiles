FROM golang:alpine AS builder
ENV GOCACHE=/gocache/build
ENV GOMODCACHE=/gocache/mod
WORKDIR /builder
RUN --mount=type=cache,target=/gocache GOBIN=$PWD CGO_ENABLED=0 go install go.winto.dev/envreplace@latest

FROM scratch
COPY --from=builder /builder/envreplace /envreplace
ENTRYPOINT ["/envreplace", "-"]
