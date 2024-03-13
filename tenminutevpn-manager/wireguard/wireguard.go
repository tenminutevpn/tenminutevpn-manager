package wireguard

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func GenKey() (string, error) {
	cmd := exec.Command("wg", "genkey")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	privkey := strings.TrimSpace(string(out))
	return privkey, nil
}

func GenPublicKey(privkey string) (string, error) {
	cmd := exec.Command("wg", "pubkey")
	cmd.Stdin = strings.NewReader(privkey)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	pubkey := strings.TrimSpace(string(out))
	return pubkey, nil
}

func GenKeypair() (string, string, error) {
	privkey, err := GenKey()
	if err != nil {
		return "", "", err
	}

	pubkey, err := GenPublicKey(privkey)
	if err != nil {
		return "", "", err
	}

	return privkey, pubkey, nil
}

func WriteKeypair(folder string, privkey string, pubkey string) error {
	privkey, pubkey, err := GenKeypair()
	if err != nil {
		return fmt.Errorf("failed to generate keypair: %w", err)
	}

	privkeyPath := path.Join(folder, "privkey")
	privkeyFile, err := os.Create(privkeyPath)
	if err != nil {
		return fmt.Errorf("failed to create private key file: %w", err)
	}
	defer privkeyFile.Close()

	_, err = privkeyFile.WriteString(privkey)
	if err != nil {
		return fmt.Errorf("failed to write private key to file: %w", err)
	}

	pubkeyPath := path.Join(folder, "pubkey")
	pubkeyFile, err := os.Create(pubkeyPath)
	if err != nil {
		return fmt.Errorf("failed to create public key file: %w", err)
	}
	defer pubkeyFile.Close()

	_, err = pubkeyFile.WriteString(pubkey)
	if err != nil {
		return fmt.Errorf("failed to write public key to file: %w", err)
	}

	return nil
}
