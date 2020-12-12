FROM golang:1.15-buster

ENV GO111MODULE=on
RUN go get github.com/ory/go-acc

COPY main.go /main.go
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
