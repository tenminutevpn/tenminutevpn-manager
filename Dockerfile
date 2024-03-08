FROM debian:12-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    bash \
    ca-certificates \
    curl \
    iproute2 \
    iptables \
    make \
    wireguard-tools

COPY src /opt/src
COPY test /opt/test
COPY Makefile /opt/Makefile
WORKDIR /opt
