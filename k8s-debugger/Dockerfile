FROM alpine:latest

RUN <<EOF
set -eu

apk -U add \
  bash \
  bind-tools \
  busybox-suid \
  conntrack-tools \
  curl \
  htop \
  iotop \
  iproute2 \
  iptables \
  iptables-legacy \
  iputils \
  jq \
  lsof \
  openssh-client \
  procps-ng \
  socat \
  sysstat \
  traceroute \
  util-linux-misc

printf 'root\nroot\n' | passwd

EOF
