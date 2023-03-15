FROM golang:1.19-alpine

COPY entrypoint.sh /entrypoint.sh

# setup user
ARG USERNAME=container-user
ARG USER_UID=3000
ARG USER_GID=$USER_UID
RUN addgroup --gid $USER_GID $USERNAME \
    && adduser --uid $USER_UID --ingroup $USERNAME -h /home/$USERNAME -s /bin/ash -D $USERNAME \
    && chmod +x /entrypoint.sh

# install required packages
RUN apk update && \
  apk add  --no-cache git openssh-client

USER $USERNAME

ENTRYPOINT [ "/entrypoint.sh" ]
