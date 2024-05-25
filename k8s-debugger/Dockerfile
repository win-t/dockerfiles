FROM alpine:latest

RUN <<EOF
apk -U add \
  bash \
  bind-tools \
  conntrack-tools \
  curl \
  htop \
  iproute2 \
  iptables \
  iptables-legacy \
  openssh-client \
  socat \
  sudo \
  traceroute \
  util-linux-misc

printf '\n\nALL ALL=NOPASSWD: ALL\n' >> /etc/sudoers
EOF
