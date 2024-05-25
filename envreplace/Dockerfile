FROM golang:alpine AS builder
WORKDIR /builder
RUN GOBIN=$PWD CGO_ENABLED=0 go install go.winto.dev/envreplace@latest

FROM alpine
COPY --from=builder /builder/envreplace /envreplace
ENTRYPOINT ["/envreplace", "-"]
