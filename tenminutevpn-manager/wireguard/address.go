package wireguard

import (
	"fmt"
	"net"
)

type Address struct {
	IP    net.IP
	IPNet *net.IPNet
}

func (addr *Address) Mask() int {
	ones, _ := addr.IPNet.Mask.Size()
	return ones
}

func (addr *Address) String() string {
	return fmt.Sprintf("%s/%d", addr.IP, addr.Mask())
}

func NewAddress(ip net.IP, ipNet *net.IPNet) *Address {
	return &Address{
		IP:    ip,
		IPNet: ipNet,
	}
}

func NewAddressFromString(addr string) (*Address, error) {
	ip, ipNet, err := net.ParseCIDR(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse address: %w", err)
	}

	return &Address{
		IP:    ip,
		IPNet: ipNet,
	}, nil
}
