# syntax=docker/dockerfile:1

FROM alpine:latest

RUN set -eux \
  && apk --no-cache add socat

ENTRYPOINT ["socat"]
