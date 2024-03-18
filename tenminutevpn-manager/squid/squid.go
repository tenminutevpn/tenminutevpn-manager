package squid

import (
	"github.com/tenminutevpn/tenminutevpn-manager/systemd"
)

type Squid struct {
	Port int
}

func NewSquid(port int) *Squid {
	return &Squid{Port: port}
}

func (s *Squid) SystemdService() *systemd.Service {
	return systemd.NewService("squid")
}

func (s *Squid) Render() string {
	return makeTemplateSquidData(s).Render()
}

func (s *Squid) Write(path string) error {
	return writeToFile(path, 0644, s.Render())
}
