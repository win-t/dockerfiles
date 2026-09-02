# syntax=docker/dockerfile:1

FROM alpine AS builder
WORKDIR /data
RUN apk -U add gcc musl-dev curl
RUN curl -sSLf -o main.c https://raw.githubusercontent.com/win-t/docker_pid1/476ef430bd88c42da18de2bf7ea82492b78bccf5/main.c
RUN gcc -std=c23 -Os -Wall -Wextra -Wpedantic -Wconversion -Warith-conversion -Werror -fdata-sections -ffunction-sections -Wl,--gc-sections -static -o docker_pid1 main.c \
  && strip -s docker_pid1

FROM busybox:musl
COPY --from=builder /data/docker_pid1 /docker_pid1
ENTRYPOINT ["/docker_pid1"]
