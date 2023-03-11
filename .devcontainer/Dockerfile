FROM golang:1.19-alpine

# setup user
ARG USERNAME=workspace-user
ARG USER_UID=3000
ARG USER_GID=$USER_UID
RUN addgroup --gid $USER_GID $USERNAME \
    && adduser --uid $USER_UID --ingroup $USERNAME -h /home/$USERNAME -s /bin/ash -D $USERNAME

# install required packages
RUN apk update && \
  apk add  --no-cache git openssh-client

USER workspace-user

RUN git config --global user.email 31839234+kengo-k@users.noreply.github.com && \
    git config --global user.name "kengo-k"
