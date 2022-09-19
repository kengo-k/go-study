FROM golang:1.19.1-alpine

RUN \
  apk update && apk add git;

WORKDIR /app/src

ENTRYPOINT ["/app/docker-entrypoint.sh"]
