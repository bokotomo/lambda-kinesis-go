FROM golang:1.12.7-stretch

LABEL maintainer boko

SHELL ["/bin/bash", "-c"]
ENV GO111MODULE on
WORKDIR /lambda-kinesis-go
RUN apt update

RUN apt-get install -y default-libmysqlclient-dev vim zip unzip