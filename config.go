package main

import (
	"fmt"
	"log"
	"os"
)

var (
	valutDir string
)

func InitConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	valutDir = fmt.Sprintf("%s/.termnote", homeDir)

	err = os.MkdirAll(valutDir, 0750)
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
