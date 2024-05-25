FROM golang:alpine AS builder
WORKDIR /builder
RUN GOBIN=$PWD CGO_ENABLED=0 go install go.winto.dev/httpecho@latest

FROM scratch
COPY --from=builder /builder/httpecho /server
ENTRYPOINT ["/server"]
