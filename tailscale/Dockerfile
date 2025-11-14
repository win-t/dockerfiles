# syntax=docker/dockerfile:1

FROM ubuntu:latest

SHELL [ "/bin/bash", "-ceux", "-o", "pipefail" ]
RUN <<EOR
  version=$( . /etc/os-release && echo $VERSION_CODENAME )
  export DEBIAN_FRONTEND=noninteractive
  apt-get update
  apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    dnsutils \
    iproute2 \
    iputils-ping \
    jq \
    tcpdump
  curl -fsSL "https://pkgs.tailscale.com/stable/ubuntu/${version}.noarmor.gpg" -o /usr/share/keyrings/tailscale-archive-keyring.gpg
  curl -fsSL "https://pkgs.tailscale.com/stable/ubuntu/${version}.tailscale-keyring.list" -o /etc/apt/sources.list.d/tailscale.list
  apt-get update
  apt-get install -y --no-install-recommends tailscale
  apt-get clean -y
  rm -rf /var/lib/apt/lists/*
EOR

VOLUME /var/lib/tailscale

COPY --from=wint/pid1 /docker_pid1 /usr/local/bin/docker_pid1
COPY entrypoint /entrypoint

ENTRYPOINT ["docker_pid1", "/entrypoint"]
