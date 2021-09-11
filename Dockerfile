FROM golang:1.17.5

ENV GO111MODULE=on

RUN apt-get update

RUN mkdir -p /sdk

WORKDIR /sdk

CMD ["make", "e2e"]