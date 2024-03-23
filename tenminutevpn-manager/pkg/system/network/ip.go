package network

import "net"

type IPNet net.IPNet

func (n *IPNet) MarshalYAML() (interface{}, error) {
	return n.String(), nil
}

func (n *IPNet) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var cidr string
	if err := unmarshal(&cidr); err != nil {
		return err
	}

	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}

	*n = IPNet(*ipNet)
	return nil
}

func (n *IPNet) Contains(ip net.IP) bool {
	return (*net.IPNet)(n).Contains(ip)
}

func (n *IPNet) Network() net.IP {
	return (*net.IPNet)(n).IP
}

func (n *IPNet) String() string {
	return (*net.IPNet)(n).String()
}
