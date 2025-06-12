# syntax=docker/dockerfile:1

FROM golang

ENV GOCACHE=/cache/go-build
ENV GOMODCACHE=/cache/go-mod
RUN GOBIN=/usr/local/bin go install go.winto.dev/gorundir@latest

WORKDIR /code
COPY code/ .

ENTRYPOINT ["gorundir"]
CMD ["."]
