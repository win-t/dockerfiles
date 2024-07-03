FROM golang:alpine AS builder
ENV GOCACHE=/gocache/build
ENV GOMODCACHE=/gocache/mod
WORKDIR /builder
RUN --mount=type=cache,target=/gocache GOBIN=$PWD CGO_ENABLED=0 go install go.winto.dev/httpecho@latest

FROM scratch
COPY --from=builder /builder/httpecho /server
ENTRYPOINT ["/server"]
