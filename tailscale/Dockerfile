FROM ubuntu:latest

SHELL [ "/bin/bash", "-ceux", "-o", "pipefail" ]
RUN <<EOR
  version=$( . /etc/os-release && echo $VERSION_CODENAME )
  apt-get update
  apt-get install -y --no-install-recommends curl ca-certificates iputils-ping dnsutils jq
  curl -fsSL "https://pkgs.tailscale.com/stable/ubuntu/${version}.noarmor.gpg" -o /usr/share/keyrings/tailscale-archive-keyring.gpg
  curl -fsSL "https://pkgs.tailscale.com/stable/ubuntu/${version}.tailscale-keyring.list" -o /etc/apt/sources.list.d/tailscale.list
  apt-get update
  apt-get install -y --no-install-recommends tailscale
  apt-get clean -y
  rm -rf \
    /var/cache/debconf/* \
    /var/lib/apt/lists/* \
    /var/log/* \
    /tmp/* \
    /var/tmp/* \
    /usr/share/doc/* \
    /usr/share/doc-base/* \
    /usr/share/man/* \
    /usr/share/local/*
EOR

VOLUME /var/lib/tailscale

COPY entrypoint /bin/entrypoint

ENTRYPOINT ["entrypoint"]
