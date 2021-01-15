#!/bin/bash

export GO111MODULE=on
export GOARCH="amd64"
export CGO_ENABLED=0

cmd=$1

binary="newschain"

if [ "$cmd" = "build" ]; then
  echo "Executing build command"
  go get .
  go build -v -o ${binary}
  exit
fi

if [ "$cmd" = "run" ]; then
  echo "Executing run command"
  ./${binary} serve
  exit
fi

if [ "$cmd" = "worker" ]; then
  echo "Executing worker command"
  ./${binary} worker
  exit
fi

if [ "$cmd" = "auto" ]; then
  echo "Executing migration auto command"
  ./${binary} migration auto
  exit
fi

if [ "$cmd" = "init" ]; then
  echo "Executing migration init command"
  ./${binary} migration init
  exit
fi

if [ "$cmd" = "extra" ]; then
  echo "Executing migration extra command"
  ./${binary} migration extra
  exit
fi

if [ "$cmd" = "populate" ]; then
  echo "Executing migration populate command"
  ./${binary} migration populate
  exit
fi

if [ "$cmd" = "drop" ]; then
  echo "Executing migration drop command"
  ./${binary} migration drop
  exit
fi

if [ "$cmd" = "gen" ]; then
  echo "Executing generate ent schema command"
  go generate ./ent
  exit
fi

echo "No command specified"
