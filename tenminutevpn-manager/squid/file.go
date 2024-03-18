package squid

import (
	"fmt"
	"os"
)

func writeToFile(filename string, mode os.FileMode, data string) error {
	if filename == "" {
		return fmt.Errorf("filename is empty")
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	err = file.Chmod(mode)
	if err != nil {
		return fmt.Errorf("failed to change file permissions: %w", err)
	}

	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("failed to write private key to file: %w", err)
	}

	return nil
}
