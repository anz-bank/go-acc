#!/bin/bash
set -e

if [ -z "$INPUT_IGNORE" ]; then
  # generate coverage profile
  /go/bin/go-acc ./... -o=cover.out -- ${INPUT_FLAGS}
else
  # generate coverage profile (with ignores)
  /go/bin/go-acc ./... --ignore "${INPUT_IGNORE//[[:space:]]/}" -o=cover.out -- ${INPUT_FLAGS}
fi

go tool cover -func=cover.out | go run /main.go