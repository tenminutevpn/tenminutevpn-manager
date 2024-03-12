# TenMinuteVPN - Tools

## Configuration

```json
{
  "wireguard": {
    "enabled": true,
    "interfaces": [
      {
        "name": "wg0",
        "privateKey": "",
        "listenPort": 51820,
        "peers": [
          {
            "name": "peer1",
            "privateKey": "",
            "allowedIPs": ""
          }
      }
    ]
  },
  "squid": {
    "enabled": true,
    "port": 3128,
  }
}
```
