# syntax=docker/dockerfile:1

FROM alpine:latest AS extra
RUN apk add -U curl gcc musl-dev
WORKDIR /work
RUN curl -sSLf -o swuser.c      https://raw.githubusercontent.com/win-t/switch-user/92697545616fff34670c15490da37ee15f860b0b/main.c \
  && curl -sSLf -o docker_pid1.c https://raw.githubusercontent.com/win-t/docker_pid1/43ff425663314c9cf22b936dcf0802032385fda6/main.c \
  && mkdir output \
  && gcc -Wall -Wextra -Wpedantic -o output/swuser      swuser.c \
  && gcc -Wall -Wextra -Wpedantic -o output/docker_pid1 docker_pid1.c

FROM alpine:latest
RUN apk -U add \
  bash \
  bind-tools \
  curl \
  findmnt \
  iproute2 \
  iputils \
  jq \
  procps-ng \
  sysstat \
  tmux \
  traceroute \
  util-linux-misc
COPY --from=extra /work/output /usr/local/bin
RUN chmod u+s /usr/local/bin/swuser

ENTRYPOINT [ "docker_pid1" ]
