package network

import (
	"io"
	"net"
	"net/http"
	"os/exec"
)

func GetDefaultInterface() (string, error) {
	cmd := exec.Command("bash", "-c", "ip route | grep default | awk '{print $5}'")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func GetPublicIPv4() (net.IP, error) {
	// get public ipv4 using https://ipinfo.io/ip

	resp, err := http.Get("https://ipinfo.io/ip")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ip := net.ParseIP(string(body))
	if ip == nil {
		return nil, err
	}

	return ip, nil
}
