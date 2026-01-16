package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

func listFiles() []list.Item {
	items := make([]list.Item, 0)

	entries, err := os.ReadDir(valutDir)
	if err != nil {
		log.Fatal("error reading notes list")
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			modifiedTime := info.ModTime().Format("2006-01-02 15:04")
			items = append(items, item{
				title: entry.Name(),
				desc:  fmt.Sprintf("Modified: %s", modifiedTime),
			})
		}
	}
	return items
}
