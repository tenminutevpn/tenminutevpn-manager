#!/bin/bash

set -e -o pipefail

TENMINUTEVPN_PATH="$(readlink -f "${BASH_SOURCE[0]}")"
export PATH="$(dirname "$TENMINUTEVPN_PATH"):$PATH"

# function that finds default interface
find_default_interface() {
    ip route | awk '/^default/ {print $5}'
}


# function that generates a new wireguard keypair in target directory
wireguard_keygen() {
    local target_dir="$1"
    mkdir -p "$target_dir"
    wg genkey | tee "$target_dir/privatekey" | wg pubkey > "$target_dir/publickey"
}
