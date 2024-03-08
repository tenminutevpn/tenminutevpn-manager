#!/bin/bash

set -e -o pipefail

TENMINUTEVPN_PATH="$(readlink -f "${BASH_SOURCE[0]}")"
export PATH="$(dirname "$TENMINUTEVPN_PATH"):$PATH"
