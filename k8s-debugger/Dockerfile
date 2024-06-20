FROM alpine:latest

RUN <<EOF
apk -U add \
  bash \
  bind-tools \
  busybox-suid \
  conntrack-tools \
  curl \
  htop \
  iostat \
  iproute2 \
  iptables \
  iptables-legacy \
  iputils \
  openssh-client \
  procps-ng \
  socat \
  sysstat \
  traceroute \
  util-linux-misc

printf 'root\nroot\n' | passwd

EOF
