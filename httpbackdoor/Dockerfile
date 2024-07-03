FROM golang:1 AS builder
ENV GOCACHE=/gocache/build
ENV GOMODCACHE=/gocache/modWORKDIR /builder
WORKDIR /builder
COPY . .
RUN --mount=type=cache,target=/gocache  go build -o server .

FROM ubuntu:jammy
RUN apt update && apt install -y curl bind9-dnsutils iproute2 iputils-ping netcat busybox
COPY --from=builder /builder/server /bin/server
ENTRYPOINT ["server"]
