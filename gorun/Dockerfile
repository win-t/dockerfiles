FROM golang
RUN GOBIN=/usr/local/bin go install go.winto.dev/gorundir@latest

WORKDIR /code
COPY code/ .

ENTRYPOINT ["gorundir"]
CMD ["."]
