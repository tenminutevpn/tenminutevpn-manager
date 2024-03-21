package utils

import (
	"fmt"
	"os"
)

func WriteToFile(filename string, mode os.FileMode, data string) error {
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

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, 0700)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}
