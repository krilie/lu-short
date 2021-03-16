#!/bin/bash

chmod u+x ./build_binary.sh
./build_binary.sh
if [ "$?" -eq 0 ];then
  echo "build ok"
else
  echo "no"
  exit 1
fi

docker build -t "$1" .
