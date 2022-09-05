#!/bin/sh

echo "> run docker-entrypoint.sh..."

export GOPATH=/app/.go
export PATH=$PATH:$GOPATH/bin

if [ $# != 0 ]; then
  echo "> $@"
  exec "$@"
else
  go install github.com/cosmtrek/air@v1.29.0;
  air -c .air.toml
fi
