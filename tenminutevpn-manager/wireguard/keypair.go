package wireguard

import (
	"fmt"
	"os/exec"
	"strings"
)

type KeyPair struct {
	PrivateKey string
	PublicKey  string
}

func (kp *KeyPair) WritePrivateKey(filename string) error {
	return writeToFile(filename, 0600, kp.PrivateKey)
}

func (kp *KeyPair) WritePublicKey(filename string) error {
	return writeToFile(filename, 0644, kp.PublicKey)
}

func NewKeyPair(privateKey string) (*KeyPair, error) {
	if privateKey == "" {
		return nil, fmt.Errorf("private key cannot be empty")
	}

	publicKey, err := GeneratePublicKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate public key: %w", err)
	}

	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func GeneratePrivateKey() (string, error) {
	cmd := exec.Command("wg", "genkey")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %w", err)
	}
	privkey := strings.TrimSpace(string(out))
	return privkey, nil
}

func GeneratePublicKey(privkey string) (string, error) {
	cmd := exec.Command("wg", "pubkey")
	cmd.Stdin = strings.NewReader(privkey)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to generate public key: %w", err)
	}
	pubkey := strings.TrimSpace(string(out))
	return pubkey, nil
}
