# This Dockerfile builds on golang:alpine by building Trple from source
# using the current working directory.
#
# This produces a docker image that contains the Trple binaries along
# with all of its source code. This is not what produces the official releases.

FROM golang:alpine
LABEL maintainer="Brett Mostert <brettmostert@gmail.com>"
LABEL app=trple

RUN apk add --no-cache bash

ENV TRPLE_DEV=true
ENV TRPLE_RELEASE=1

WORKDIR $GOPATH/src/github.com/brettmostert/trple
COPY . .
RUN /bin/bash ./scripts/build.sh

WORKDIR $GOPATH/bin
ENTRYPOINT ["trple-server"]