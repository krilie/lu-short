#!/bin/bash

docker run -it -v "$PWD":/app -v "$PWD"/.cache/gopath:/go -v "$PWD"/.cache/gocache:/root/.cache/ golang:1.16-alpine3.12 \
           /bin/sh -c "cd /app && export CGO_ENABLED=0 && export GOPROXY=https://goproxy.io,direct && go build -v -o lu-short"
if [ "$?" -eq 0 ];then
  echo "build ok"
else
  echo "no"
  exit 1
fi
