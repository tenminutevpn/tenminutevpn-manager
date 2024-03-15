package wireguard

import (
	"fmt"
	"os"
)

func writeToFile(filename string, data string) error {
	if filename == "" {
		return fmt.Errorf("filename is empty")
	}

	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("file already exists: %s", filename)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	err = file.Chmod(0600)
	if err != nil {
		return fmt.Errorf("failed to change file permissions: %w", err)
	}

	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("failed to write private key to file: %w", err)
	}

	return nil
}
