#!/usr/bin/env bash

echo "==> Checking that code complies with static analysis requirements..."
# skip=""
# packages=$(go list ./... | egrep -v ${skip})
packages=$(go list ./...)

# We are skipping style-related checks, since terraform intentionally breaks
# some of these. The goal here is to find issues that reduce code clarity, or
# may result in bugs.
go run honnef.co/go/tools/cmd/staticcheck -checks 'all' ${packages}