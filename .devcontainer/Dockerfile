FROM debian:12-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    bash \
    ca-certificates \
    curl \
    iproute2 \
    iptables \
    make \
    shellcheck \
    wireguard-tools

COPY src/tenminutevpn.bash /usr/bin/tenminutevpn
COPY test /opt/test
COPY Makefile /opt/Makefile

WORKDIR /opt
