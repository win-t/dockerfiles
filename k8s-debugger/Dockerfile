FROM alpine:latest

RUN <<EOF
apk -U add \
  bash \
  bind-tools \
  busybox-suid \
  conntrack-tools \
  curl \
  htop \
  iproute2 \
  iptables \
  iptables-legacy \
  iputils \
  openssh-client \
  socat \
  traceroute \
  util-linux-misc

printf 'root\nroot\n' | passwd

EOF
