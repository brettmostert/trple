#!/usr/bin/env bash

source $(dirname "$0")/common.sh
getProjectDir

# Build a image with src code and debug
docker build . --rm --force-rm --target release -t trple:latest --build-arg RELEASE=1

# Done!
echo
echo -e "==> ${YELLOW}Results:${NC}"
docker image ls -f label="app=trple"
echo