# syntax=docker/dockerfile:1

FROM alpine:latest AS extra
RUN apk add -U curl gcc musl-dev
WORKDIR /work
RUN curl -sSLf -o docker_pid1.c https://raw.githubusercontent.com/win-t/docker_pid1/2056c8955a863a3e23655815f3fca53ac9d74975/main.c \
  && mkdir output \
  && gcc -Wall -Wextra -Wpedantic -o output/docker_pid1 docker_pid1.c

FROM alpine:latest
RUN apk -U add wireshark-common
COPY --from=extra /work/output /usr/local/bin
ENTRYPOINT ["docker_pid1"]
