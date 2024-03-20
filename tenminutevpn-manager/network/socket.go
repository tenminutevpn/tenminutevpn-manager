package network

import "net"

type Socket struct {
	Protocol string `yaml:"protocol"`
	Address  net.IP `yaml:"address"`
	Port     int    `yaml:"port"`
}

func (s *Socket) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type socket Socket
	var tmp socket
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	s.Protocol = tmp.Protocol
	s.Address = net.ParseIP(tmp.Address.String())
	s.Port = tmp.Port
	return nil
}
