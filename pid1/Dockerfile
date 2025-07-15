# syntax=docker/dockerfile:1

FROM alpine AS builder
WORKDIR /data
RUN apk -U add gcc musl-dev curl
RUN curl -sSLf -o main.c https://raw.githubusercontent.com/win-t/docker_pid1/43ff425663314c9cf22b936dcf0802032385fda6/main.c
RUN gcc -std=c11 -Os -Wall -Wextra -Wpedantic -Werror -fdata-sections -ffunction-sections -Wl,--gc-sections -static -o docker_pid1 main.c \
  && strip -s docker_pid1

FROM busybox:musl
COPY --from=builder /data/docker_pid1 /docker_pid1
ENTRYPOINT ["/docker_pid1"]
