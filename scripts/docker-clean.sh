#!/usr/bin/env bash

# Removes all build related images for "trple"

source $(dirname "$0")/common.sh
getProjectDir

echo
echo "==> Checking trple images to remove:"
docker image ls -f label="app=trple" -f label="stage=build"

IMAGES=$(docker image ls -q -f label="app=trple" -f label="stage=build")

if [[ -n "${IMAGES}" ]]; then
    docker image rm -f ${IMAGES}
fi

# Done!
echo
echo "==> Results:"
docker image ls -f label="app=trple" -f label="stage=build"