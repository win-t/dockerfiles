FROM golang AS derper-build
WORKDIR /build
RUN GOBIN=$PWD go install tailscale.com/cmd/derper@latest

FROM ubuntu:latest

RUN --mount=type=bind,from=derper-build,source=/build,dst=/derper-build,rw \
    --mount=type=bind,from=wint/pid1,source=/,dst=/pid1-build,rw \
    --mount=type=bind,source=/,dst=/context \
    <<'EOR'
#!/bin/bash

set -euxo pipefail
version=$( . /etc/os-release && echo $VERSION_CODENAME )
export DEBIAN_FRONTEND=noninteractive
apt-get update
apt-get install -y \
  curl \
  dnsutils \
  iproute2 \
  iputils-ping \
  jq
curl -fsSL "https://pkgs.tailscale.com/stable/ubuntu/${version}.noarmor.gpg" -o /usr/share/keyrings/tailscale-archive-keyring.gpg
curl -fsSL "https://pkgs.tailscale.com/stable/ubuntu/${version}.tailscale-keyring.list" -o /etc/apt/sources.list.d/tailscale.list
apt-get update
apt-get install -y tailscale
apt-get clean -y
rm -rf /var/lib/apt/lists/*

cp /derper-build/derper /usr/local/bin/derper
cp /pid1-build/docker_pid1 /usr/local/bin/docker_pid1
cp /context/entrypoint /entrypoint

EOR

ENTRYPOINT ["docker_pid1", "/entrypoint"]
