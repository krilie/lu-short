#!/bin/bash

docker build -t "$1" -f Dockerfile.full .
if [ "$?" -eq 0 ];then
  echo "build ok"
else
  echo "no"
  exit 1
fi
