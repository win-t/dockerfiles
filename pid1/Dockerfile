FROM alpine AS builder
WORKDIR /data
RUN apk -U add gcc musl-dev curl
RUN curl -sSLf -o main.c https://raw.githubusercontent.com/win-t/docker_pid1/refs/heads/master/main.c
RUN gcc -std=c11 -Os -Wall -Wextra -Wpedantic -Werror -fdata-sections -ffunction-sections -Wl,--gc-sections -static -o docker_pid1 main.c \
 && strip -s docker_pid1

FROM scratch
COPY --from=builder /data/docker_pid1 /docker_pid1
ENTRYPOINT ["/docker_pid1"]
