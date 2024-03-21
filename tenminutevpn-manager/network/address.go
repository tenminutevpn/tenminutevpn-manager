package network

import (
	"fmt"
	"net"
)

type Address struct {
	IP      net.IP
	Network IPNet
}

func (addr Address) Mask() int {
	ones, _ := addr.Network.Mask.Size()
	return ones
}

func (addr Address) String() string {
	return fmt.Sprintf("%s/%d", addr.IP, addr.Mask())
}

func (addr Address) MarshalYAML() (interface{}, error) {
	return addr.String(), nil
}

func (addr *Address) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var addrStr string
	if err := unmarshal(&addrStr); err != nil {
		return err
	}

	ip, ipNet, err := net.ParseCIDR(addrStr)
	if err != nil {
		return fmt.Errorf("failed to parse address: %w", err)
	}

	addr.IP = ip
	addr.Network = IPNet(*ipNet)
	return nil
}

func NewAddress(ip net.IP, network *net.IPNet) *Address {
	return &Address{
		IP:      ip,
		Network: IPNet(*network),
	}
}

func NewAddressFromString(addr string) (*Address, error) {
	ip, ipNet, err := net.ParseCIDR(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse address: %w", err)
	}

	return &Address{
		IP:      ip,
		Network: IPNet(*ipNet),
	}, nil
}
