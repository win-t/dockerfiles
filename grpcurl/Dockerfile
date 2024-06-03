FROM golang:1-alpine AS builder
WORKDIR /builder
COPY . .
RUN GOBIN=$PWD go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

FROM alpine
RUN apk -u add curl openssl bind-tools
COPY --from=builder /builder/grpcurl /bin/grpcurl
