FROM alpine:latest AS extra
RUN apk add -U curl gcc musl-dev
WORKDIR /work
RUN curl -sSLf -o swuser.c      https://raw.githubusercontent.com/win-t/switch-user/refs/heads/master/main.c \
 && curl -sSLf -o docker_pid1.c https://raw.githubusercontent.com/win-t/docker_pid1/refs/heads/master/main.c \
 && mkdir output \
 && gcc -std=c11 -Os -Wall -Wextra -Wpedantic -Werror -o output/swuser      swuser.c \
 && gcc -std=c11 -Os -Wall -Wextra -Wpedantic -Werror -o output/docker_pid1 docker_pid1.c

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
COPY --from=extra /work/output /usr/local/bin
RUN chmod u+s /usr/local/bin/swuser
