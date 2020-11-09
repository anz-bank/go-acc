#!/bin/bash

docker build . -t coverage-action
export INPUT_HARD_TARGET=80
export INPUT_SOFT_TARGET=95

docker run --rm \
  --workdir /github/workspace \
  -v $(pwd):/github/workspace \
  -e INPUT_HARD_TARGET \
  -e INPUT_SOFT_TARGET \
  coverage-action