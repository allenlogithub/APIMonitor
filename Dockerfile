FROM golang:1.17-buster

# linux
RUN apt-get update \
    && apt-get install -y htop

ENV PATH="$PATH:$(go env GOPATH)/bin"

EXPOSE 80
