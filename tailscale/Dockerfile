FROM ubuntu:latest

SHELL [ "/bin/bash", "-ceux", "-o", "pipefail" ]
RUN <<EOR
  version=$( . /etc/os-release && echo $VERSION_CODENAME )
  apt-get update
  apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    dnsutils \
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

COPY entrypoint /bin/entrypoint

ENTRYPOINT ["entrypoint"]
