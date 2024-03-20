package network

import (
	"fmt"
	"net"
	"strconv"
)

type Endpoint struct {
	IP   net.IP
	Port int
}

func (endpoint *Endpoint) String() string {
	if endpoint.IP == nil {
		return fmt.Sprintf(":%d", endpoint.Port)
	}

	return net.JoinHostPort(endpoint.IP.String(), fmt.Sprintf("%d", endpoint.Port))
}

func (endpoint *Endpoint) MarshalYAML() (interface{}, error) {
	return endpoint.String(), nil
}

func (endpoint *Endpoint) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var endpointStr string
	if err := unmarshal(&endpointStr); err != nil {
		return err
	}

	host, portStr, err := net.SplitHostPort(endpointStr)
	if err != nil {
		return fmt.Errorf("failed to parse endpoint: %w", err)
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return fmt.Errorf("invalid ip address: %s", host)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("failed to parse port: %w", err)
	}

	if port < 0 || port > 65535 {
		return fmt.Errorf("invalid port: %d", port)
	}

	endpoint.IP = ip
	endpoint.Port = port
	return nil
}

func NewEndpoint(ip net.IP, port int) *Endpoint {
	return &Endpoint{
		IP:   ip,
		Port: port,
	}
}
