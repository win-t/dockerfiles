# syntax=docker/dockerfile:1

FROM alpine AS builder
WORKDIR /data
RUN apk -U add gcc musl-dev curl
RUN curl -sSLf -o main.c https://raw.githubusercontent.com/win-t/docker_pid1/2056c8955a863a3e23655815f3fca53ac9d74975/main.c
RUN gcc -std=c11 -Os -Wall -Wextra -Wpedantic -Werror -fdata-sections -ffunction-sections -Wl,--gc-sections -static -o docker_pid1 main.c \
  && strip -s docker_pid1

FROM busybox:musl
COPY --from=builder /data/docker_pid1 /docker_pid1
ENTRYPOINT ["/docker_pid1"]
