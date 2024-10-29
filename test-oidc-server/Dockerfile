FROM golang:alpine AS builder
ENV GOCACHE=/gocache/build
ENV GOMODCACHE=/gocache/mod
WORKDIR /builder
RUN --mount=type=cache,target=/gocache GOBIN=$PWD CGO_ENABLED=0 go install go.winto.dev/test-oidc-server@latest

FROM scratch
COPY --from=builder /builder/test-oidc-server /server
ENV KEY=/data/key.pem
ENTRYPOINT ["/server"]
