FROM golang:alpine AS builder
WORKDIR /builder
RUN GOBIN=$PWD go install go.winto.dev/httpecho@latest

FROM alpine
COPY --from=builder /builder/httpecho /server
ENTRYPOINT ["/server"]
