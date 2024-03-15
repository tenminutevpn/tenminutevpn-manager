package wireguard

import (
	"fmt"
	"net"
)

type Address struct {
	IP    net.IP
	IPNet *net.IPNet
}

func (a *Address) Mask() int {
	ones, _ := a.IPNet.Mask.Size()
	return ones
}

func (a *Address) String() string {
	return fmt.Sprintf("%s/%d", a.IP, a.Mask())
}

func NewAddress(ip net.IP, ipNet *net.IPNet) *Address {
	return &Address{
		IP:    ip,
		IPNet: ipNet,
	}
}

func NewAddressFromString(address string) (*Address, error) {
	ip, ipNet, err := net.ParseCIDR(address)
	if err != nil {
		return nil, fmt.Errorf("failed to parse address: %w", err)
	}

	return &Address{
		IP:    ip,
		IPNet: ipNet,
	}, nil
}
