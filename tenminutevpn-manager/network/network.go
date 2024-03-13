package network

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"os/exec"
	"runtime"
)

func GetDefaultInterface() (string, error) {
	var bashCmd string
	if runtime.GOOS == "linux" {
		bashCmd = "ip route | grep default | awk '{print $5}'"
	} else if runtime.GOOS == "darwin" {
		bashCmd = "route -n get default | grep interface | awk '{print $2}'"
	} else {
		return "", fmt.Errorf("unsupported os: %s", runtime.GOOS)
	}

	cmd := exec.Command("bash", "-c", bashCmd)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(out)), nil

}

func GetPublicIPv4() (net.IP, error) {
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
		err = fmt.Errorf("invalid ip address: %s", string(body))
		return nil, err
	}

	return ip, nil
}

func GetPrivateIPv4(ifaceName string) (net.IP, error) {
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return nil, err
	}

	addrs, err := iface.Addrs()
	if err != nil {

		return nil, err
	}

	for _, addr := range addrs {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			return nil, err
		}

		if ip.To4() != nil {
			return ip, nil
		}
	}

	return nil, fmt.Errorf("no private ipv4 address found")
}
