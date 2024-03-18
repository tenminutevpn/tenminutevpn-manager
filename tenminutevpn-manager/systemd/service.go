package systemd

import (
	"fmt"
	"os/exec"
)

type Service struct {
	Name string
}

func NewService(name string) *Service {
	return &Service{Name: name}
}

func (s *Service) Enable() error {
	cmd := exec.Command("systemctl", "enable", s.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to enable %s: %w: %s", s.Name, err, string(out))
	}
	if len(out) > 0 {
		return fmt.Errorf("failed to enable %s: %s", s.Name, string(out))
	}
	return nil
}

func (s *Service) Start() error {
	cmd := exec.Command("systemctl", "start", s.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to start %s: %w: %s", s.Name, err, string(out))
	}
	if len(out) > 0 {
		return fmt.Errorf("failed to start %s: %s", s.Name, string(out))
	}
	return nil
}

func (s *Service) Stop() error {
	cmd := exec.Command("systemctl", "stop", s.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to stop %s: %w: %s", s.Name, err, string(out))
	}
	if len(out) > 0 {
		return fmt.Errorf("failed to stop %s: %s", s.Name, string(out))
	}
	return nil
}
