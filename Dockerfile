FROM golang:1.17.7-alpine

RUN \
  apk update && apk add git;

WORKDIR /app/src

ENTRYPOINT ["/app/docker-entrypoint.sh"]
