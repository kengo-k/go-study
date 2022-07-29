FROM ubuntu:20.04
SHELL ["/bin/bash", "-c"]

ENV DEBIAN_FRONTEND=noninteractive

RUN \
  apt-get update \
    && apt-get -y install curl \
    && apt-get -y install vim

RUN \
  curl -OL https://go.dev/dl/go1.18.4.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz \
    && echo 'PATH=$PATH:/usr/local/go/bin' > ~/.bashrc

CMD /bin/bash
