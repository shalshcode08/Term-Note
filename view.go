package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// View renders the current state of the application
func (m model) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).
		PaddingLeft(2).
		PaddingRight(2)

	welcome := style.Render("Welcome to TERMNOTE 🗒️")

	helpKeys := "Ctrl+N : new file . Ctrl+L: list . Esc: back . Ctrl+S: save . Ctrl+Q: quit"

	view := ""
	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.textArea.View()
	}

	if m.showingList {
		view = m.fileList.View()
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, helpKeys)
}
