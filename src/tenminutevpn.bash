#!/bin/bash

set -e -o pipefail
if [ -n "$TENMINUTEVPN_DEBUG" ]; then
    set -x
fi

TENMINUTEVPN_PATH="$(readlink -f "${BASH_SOURCE[0]}")"
PATH="$(dirname "$TENMINUTEVPN_PATH"):$PATH"
export PATH

TENMINUTEVPN_CONFIG_PATH="${TENMINUTEVPN_CONFIG_PATH:-/etc/default/tenminutevpn}"
if [ -f "$TENMINUTEVPN_CONFIG_PATH" ]; then
    # shellcheck source=/dev/null
    source "$TENMINUTEVPN_CONFIG_PATH"
fi

### network ####################################################################

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

### wireguard ##################################################################

WIREGUARD_INTERFACE="${WIREGUARD_INTERFACE:-wg0}"
WIREGUARD_ADDRESS="${WIREGUARD_ADDRESS:-100.96.0.1/24}"
WIREGUARD_PORT="${WIREGUARD_PORT:-51820}"

# function that generates a new wireguard keypair in target directory
wireguard_generate_keypair() {
    local target_dir="$1"
    mkdir -p "$target_dir"
    wg genkey | tee "$target_dir/privatekey" | wg pubkey > "$target_dir/publickey"
    chmod 600 "$target_dir/privatekey"
}


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

    mkdir -p "$WIREGUARD_PEERS_PATH"
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

# function that adds a new wireguard peer
# shellcheck disable=SC2317
wireguard_add_peer() {
    local interface="$1"
    local publickey="$2"
    local allowed_ips="$3"

    cat <<EOF >> "/etc/wireguard/$interface.conf"
[Peer]
PublicKey = $publickey
AllowedIPs = $allowed_ips
EOF
}

# function that starts wireguard interface
wireguard_start() {
    systemctl start "wg-quick@$WIREGUARD_INTERFACE"
    systemctl enable "wg-quick@$WIREGUARD_INTERFACE"
}

# function that reloads wireguard interface
# shellcheck disable=SC2317
wireguard_reload() {
    systemctl reload "wg-quick@$WIREGUARD_INTERFACE"
}

# function that stops wireguard interface
# shellcheck disable=SC2317
wireguard_stop() {
    systemctl stop "wg-quick@$WIREGUARD_INTERFACE"
    systemctl disable "wg-quick@$WIREGUARD_INTERFACE"
}

# function that sets up wireguard interface

wireguard_setup() {
    local interface="$1"
    local ipv4_private="$2"

    wireguard_generate_keypair "/etc/wireguard"
    wireguard_generate_server_config "$interface" "$(cat /etc/wireguard/privatekey)"
    wireguard_generate_client_config "$interface" "$(cat /etc/wireguard/privatekey)" "$(cat /etc/wireguard/publickey)" "$(network_ipv4_public)"
    wireguard_start
}

### squid ######################################################################

# function that creates a new squid configuration file

squid_generate_config() {
    local interface="$1"
    local ipv4_private="$2"

    cat <<EOF > "/etc/squid/squid.conf"
http_port 3128
http_access allow all
EOF
}

# function that starts squid

squid_start() {
    systemctl start "squid"
    systemctl enable "squid"
}

# function that stops squid
# shellcheck disable=SC2317
squid_stop() {
    systemctl stop "squid"
    systemctl disable "squid"
}

# function that sets up squid
# shellcheck disable=SC2317
squid_setup() {
    local interface="$1"
    local ipv4_private="$2"

    squid_generate_config "$interface" "$ipv4_private"
    squid_start
}


### main #######################################################################

# function that prints usage

usage() {
    cat <<EOF
Usage: $0 [OPTION]...
EOF
}

# function that sets up the server

setup() {
    local interface
    interface="$(network_interface_default)"
    local ipv4_private
    ipv4_private="$(network_ipv4_private "$interface")"

    wireguard_setup "$interface" "$ipv4_private"
    squid_setup "$interface" "$ipv4_private"
}

# parses command line options

case "$1" in
    setup)
        setup
        ;;
    *)
        usage
        ;;
esac

exit 0
