FROM debian:12-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    bash \
    ca-certificates \
    curl \
    iproute2 \
    iptables \
    wireguard-tools

COPY --chmod=0744 ./src/tenminutevpn.bash /usr/local/bin/tenminutevpn
