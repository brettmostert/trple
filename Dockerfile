# This Dockerfile builds on golang:alpine by building Trple from source
# using the current working directory.
#

# This produces a docker image that contains the Trple binaries along
# with all of its source code.
FROM golang:alpine as dev

LABEL app=trple \
      stage=build \
      maintainer="Brett Mostert <brettmostert@gmail.com>"

RUN apk add --no-cache bash

ARG RELEASE
ENV TRPLE_DEV=true
ENV TRPLE_RELEASE=$RELEASE

WORKDIR /go/src/app
COPY . .
RUN /bin/bash ./scripts/build.sh

ENTRYPOINT ["/go/src/app/bin/trple-server"]

FROM scratch as release

LABEL app=trple \
      stage=release \    
      maintainer="Brett Mostert <brettmostert@gmail.com>"

COPY --from=dev /go/src/app/bin/. /trple
ENTRYPOINT ["/trple/trple-server"]
