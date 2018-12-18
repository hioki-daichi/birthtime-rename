FROM ubuntu:16.04
MAINTAINER Daichi Hioki <hiokidaichi@gmail.com>

RUN apt-get update && apt-get install -y curl git bash make

WORKDIR /usr/local
RUN curl -O https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz \
  && tar -xzvf go1.11.4.linux-amd64.tar.gz \
  && rm -f go1.11.4.linux-amd64.tar.gz

ENV GOPATH /go
RUN mkdir -p "$GOPATH/src/github.com/hioki-daichi/birthtime-rename" && chmod -R 777 "$GOPATH"

ENV PATH /usr/local/go/bin:$GOPATH/bin:$PATH

WORKDIR $GOPATH/src/github.com/hioki-daichi/birthtime-rename
