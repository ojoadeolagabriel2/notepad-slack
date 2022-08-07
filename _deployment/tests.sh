#!/usr/bin/env bash

ROOT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && cd .. && pwd )
cd "$ROOT_DIR" || exit

# tool check
command -v k6 >/dev/null 2>&1 || { echo >&2 "load test requires k6 (https://k6.io/), terminating"; exit 1; }

# build and run load tests
# go build .
# "$ROOT_DIR"/notepad-slack
k6 run _deployment/load_test.js