FROM golang:alpine AS builder
ENV GOCACHE=/gocache/build
ENV GOMODCACHE=/gocache/mod
WORKDIR /builder
COPY . .
RUN --mount=type=cache,target=/gocache go build -o server .

FROM alpine
RUN apk -U add git-daemon
COPY --from=builder /builder/server /server
ENV GIT_PROJECT_ROOT=/data
ENTRYPOINT ["/server"]
