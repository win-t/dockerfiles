FROM alpine/curl AS builder
RUN set -eux \
&& case "$(uname -m)" in \
      x86_64) suf=_amd64;; \
      aarch64) suf=_aarch64;; \
   esac \
&& curl -sSLf -o /docker_pid1 https://github.com/win-t/docker_pid1/releases/latest/download/docker_pid1$suf \
&& chmod +x /docker_pid1

FROM scratch
COPY --from=builder /docker_pid1 /docker_pid1
ENTRYPOINT ["/docker_pid1"]
