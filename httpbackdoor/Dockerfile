FROM golang:1 AS builder
WORKDIR /builder
COPY . .
RUN go build -o server .

FROM ubuntu
RUN apt update && apt install -y curl bind9-dnsutils iproute2 iputils-ping netcat busybox
COPY --from=builder /builder/server /bin/server
ENTRYPOINT ["server"]
