FROM alpine:latest AS swuser-builder
RUN apk add -U curl gcc musl-dev
WORKDIR /work
RUN curl -sSLf https://raw.githubusercontent.com/win-t/switch-user/e048cd00269fa9e1fb66f1c633ca4baca980a03c/main.c -o swuser.c \
 && gcc -o swuser swuser.c

FROM wint/pid1 AS pid1

FROM alpine:latest
RUN apk -U add \
  bash \
  bind-tools \
  conntrack-tools \
  curl \
  gawk \
  htop \
  iftop \
  iotop \
  iproute2 \
  iptables \
  iptables-legacy \
  iputils \
  jq \
  lsof \
  openssh-client \
  procps-ng \
  sed \
  socat \
  sysstat \
  tcpdump \
  traceroute \
  util-linux-misc
COPY --from=swuser-builder /work/swuser /usr/local/bin
RUN chmod u+s /usr/local/bin/swuser
COPY --from=pid1 /docker_pid1 /usr/local/bin/docker_pid1
