#!/usr/bin/env sh

ROOT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && cd .. && pwd )
cd "$ROOT_DIR" || exit

export PATH=$PATH:/usr/local/go/bin

TAG=$(date +"%s.%3N")

docker buildx create --use
docker buildx build --platform linux/amd64,linux/arm64 -t ojoadeolagabriel/notebook-slack:"$TAG" .