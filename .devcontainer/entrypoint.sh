#!/bin/sh

echo ">>>entrypoint.sh..."

git config --global user.email $GIT_EMAIL
git config --global user.name $GIT_USER

exec "$@"