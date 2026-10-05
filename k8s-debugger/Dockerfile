# syntax=docker/dockerfile:1

FROM alpine:latest AS extra
RUN apk add -U curl gcc musl-dev
WORKDIR /work
RUN curl -sSLf -o swuser.c      https://raw.githubusercontent.com/win-t/switch-user/refs/heads/master/main.c \
 && curl -sSLf -o docker_pid1.c https://raw.githubusercontent.com/win-t/docker_pid1/refs/heads/master/main.c \
 && mkdir output \
 && gcc -Wall -Wextra -Wpedantic -o output/swuser      swuser.c \
 && gcc -Wall -Wextra -Wpedantic -o output/docker_pid1 docker_pid1.c

FROM alpine:latest
RUN --mount=type=bind,from=extra,source=/work/output,dst=/extra <<'EOR'
#!/bin/sh

set -eux

cp /extra/* /usr/local/bin
chmod u+s /usr/local/bin/swuser

apk -U add \
  bash \
  bind-tools \
  busybox-static \
  curl \
  findmnt \
  iproute2 \
  iputils \
  jq \
  procps-ng \
  sysstat \
  tmux \
  traceroute \
  util-linux-misc

mkdir /tmp/grpcurl
cd /tmp/grpcurl
curl -L https://github.com/fullstorydev/grpcurl/releases/download/v1.9.4/grpcurl_1.9.4_linux_x86_64.tar.gz | tar xz
ls -lha
mv grpcurl /usr/local/bin
cd /
rm -rf /tmp/grpcurl
EOR

ENTRYPOINT [ "docker_pid1" ]
