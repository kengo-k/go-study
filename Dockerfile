FROM golang:1.19.1-alpine

RUN \
  apk update; \
  apk add git; \
  apk add make;

WORKDIR /app

ENTRYPOINT ["/app/docker-entrypoint.sh"]
CMD ["/bin/ash"]
