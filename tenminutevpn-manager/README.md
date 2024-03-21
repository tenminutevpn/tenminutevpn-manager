# Config Sample

```yaml
---
kind: squid/v1
metadata:
  name: squid-1
  annotations:
    tenminutevpn.com/config-dir: /etc/squid/
spec:
  port: 3128
---
kind: wireguard/v1
metadata:
  name: wireguard-1
  annotations:
    tenminutevpn.com/config-dir: /etc/wireguard/
    tenminutevpn.com/peer-config-dir: /etc/wireguard/peers/
spec:
  device: wg0
  address: 100.96.0.1/24
  dns:
    - 1.1.1.1
    - 1.0.0.1
  peers:
    - allowedips:
        - 100.96.0.2/32
    - allowedips:
        - 100.86.0.3/32
```