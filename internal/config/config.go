package config

import (
	"fmt"
	"log"
	"os"
)

var (
	VaultDir string
)

func InitConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	VaultDir = fmt.Sprintf("%s/.termnote", homeDir)

	err = os.MkdirAll(VaultDir, 0750)
	if err != nil {
		return fmt.Errorf("error creating vault directory: %w", err)
	}

	return nil
}

func init() {
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}
}
