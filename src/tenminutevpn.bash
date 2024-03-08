#!/bin/bash

set -e -o pipefail

TENMINUTEVPN_PATH="$(readlink -f "${BASH_SOURCE[0]}")"
export PATH="$(dirname "$TENMINUTEVPN_PATH"):$PATH"

# function that finds default interface
network_interface_default() {
    ip route | awk '/^default/ {print $5}'
}

# function that finds public ipv4 address
network_ipv4_public() {
    curl -s https://ipinfo.io/ip
}

# function that finds private ipv4 address
network_ipv4_private() {
    local interface="$1"
    ip addr show dev "$interface" | awk '/inet / {print $2}' | cut -d/ -f1
}

# function that generates a new wireguard keypair in target directory
wireguard_generate_keypair() {
    local target_dir="$1"
    mkdir -p "$target_dir"
    wg genkey | tee "$target_dir/privatekey" | wg pubkey > "$target_dir/publickey"
    chmod 600 "$target_dir/privatekey"
}

# function that generates a new wireguard configuration file
wireguard_generate_config() {
    true;
}
