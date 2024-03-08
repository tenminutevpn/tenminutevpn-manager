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

WIREGUARD_INTERFACE="wg0"
WIREGUARD_ADDRESS="100.96.0.1/24"
WIREGUARD_PORT="51820"

# function that generates a new wireguard configuration file
wireguard_generate_server_config() {
    local interface="$1"
    local privatekey="$2"

    cat <<EOF > "/etc/wireguard/$WIREGUARD_INTERFACE.conf"
[Interface]
PrivateKey = $privatekey
Address = $WIREGUARD_ADDRESS
ListenPort = $WIREGUARD_PORT
PostUp = iptables -A FORWARD -i wg0 -j ACCEPT; iptables -t nat -A POSTROUTING -o ${interface} -j MASQUERADE
PostDown = iptables -D FORWARD -i wg0 -j ACCEPT; iptables -t nat -D POSTROUTING -o ${interface} -j MASQUERADE
EOF

    chmod 600 "/etc/wireguard/$WIREGUARD_INTERFACE.conf"
}

WIREGUARD_PEERS_PATH="/etc/wireguard/peers"

# function that generates a new wireguard client configuration file
wireguard_generate_client_config() {
    local interface="$1"
    local privatekey="$2"
    local publickey="$3"
    local endpoint="$4"

    cat <<EOF > "$WIREGUARD_PEERS_PATH/client-1.conf"
[Interface]
PrivateKey = $privatekey
Address = 100.96.0.2/32
DNS = 1.1.1.1,1.0.0.1

[Peer]
PublicKey = $publickey
AllowedIPs = 0.0.0.0/0
Endpoint = $endpoint
PersistentKeepalive = 25
EOF

    chmod 600 "$WIREGUARD_PEERS_PATH/client-1.conf"
}

# function that starts wireguard interface
wireguard_start() {
    local interface="$1"
    systemctl start "wg-quick@$interface"
    systemctl enable "wg-quick@$interface"
}
