FROM golang:1.13.7-buster

ENV GO111MODULE=on
RUN go get github.com/ory/go-acc
ENV GOFLAGS=-mod=vendor

COPY main.go /main.go
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]