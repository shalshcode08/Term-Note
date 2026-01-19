package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// renderLanding renders the beautiful landing page
func renderLanding() string {
	// Header with ASCII art (using "box" style, can be changed)
	header := HeaderStyle.Render(GetASCIIArt("box"))

	// Subtitle
	subtitle := SubtitleStyle.Render("📝 Your Terminal Note-Taking Companion")

	// Help section
	helpTitle := HelpTitleStyle.Render("⌨️  Keyboard Shortcuts")

	helpItems := []struct {
		key  string
		desc string
	}{
		{"Ctrl + N", "Create a new note"},
		{"Ctrl + L", "List all notes"},
		{"Ctrl + S", "Save current note"},
		{"Esc", "Go back / Close view"},
		{"Ctrl + C", "Quit application"},
	}

	var helpLines []string
	for _, item := range helpItems {
		line := lipgloss.JoinHorizontal(
			lipgloss.Left,
			KeyStyle.Render(item.key),
			DescStyle.Render(item.desc),
		)
		helpLines = append(helpLines, HelpItemStyle.Render(line))
	}

	helpSection := lipgloss.JoinVertical(
		lipgloss.Left,
		helpTitle,
		strings.Join(helpLines, "\n"),
	)

	// Combine everything
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		subtitle,
		BoxStyle.Render(helpSection),
	)

	return content
}

// renderCreateNoteDialog renders a beautiful dialog for creating new notes
func renderCreateNoteDialog(input textinput.Model, statusMsg string, statusType string) string {
	// Title with icon
	title := DialogTitleStyle.Render("📝  CREATE NEW NOTE")

	// Label for input with character counter
	charCount := len(input.Value())
	maxChars := input.CharLimit
	counterStyle := lipgloss.NewStyle().Foreground(ColorMuted)
	if charCount > maxChars-10 {
		counterStyle = counterStyle.Foreground(ColorWarning)
	}

	labelWithCounter := lipgloss.JoinHorizontal(
		lipgloss.Left,
		InputLabelStyle.Render("Note Name:"),
		counterStyle.Render(fmt.Sprintf("  %d/%d", charCount, maxChars)),
	)

	// Input with prompt symbol
	promptSymbol := lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true).
		Render("› ")

	inputValue := lipgloss.JoinHorizontal(lipgloss.Left, promptSymbol, input.View())
	inputBox := InputBoxStyle.Render(inputValue)

	// File extension hint
	extensionHint := FileExtensionStyle.Render(".md extension will be added automatically")

	// Status message (if any)
	var statusLine string
	if statusMsg != "" {
		switch statusType {
		case "error":
			statusLine = ErrorStyle.Render("❌ " + statusMsg)
		case "warning":
			statusLine = WarningStyle.Render("⚠️  " + statusMsg)
		case "success":
			statusLine = SuccessStyle.Render("✓ " + statusMsg)
		default:
			statusLine = ViewHelpStyle.Render(statusMsg)
		}
	}

	// Help text
	helpText := InputHelpStyle.Render("⏎ Enter to create  •  Esc to cancel")

	// Combine all elements
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		labelWithCounter,
		inputBox,
		extensionHint,
		"",
		statusLine,
		"",
		helpText,
	)

	// Wrap in dialog box
	dialog := DialogBoxStyle.Render(content)

	// Center on screen
	return lipgloss.Place(
		80, 20,
		lipgloss.Center, lipgloss.Center,
		dialog,
	)
}

// renderFileListView renders the file list with enhanced styling
func renderFileListView(fileList list.Model) string {
	// Check if list is empty
	if len(fileList.Items()) == 0 {
		emptyState := lipgloss.NewStyle().
			Foreground(ColorMuted).
			Align(lipgloss.Center).
			Padding(10, 2).
			Render("📝 No notes yet!\n\nPress Ctrl+N to create your first note")

		return emptyState
	}

	listView := fileList.View()

	// Add custom help text at bottom if not filtering
	if fileList.FilterState() != list.Filtering {
		helpText := lipgloss.NewStyle().
			Foreground(ColorMuted).
			Padding(1, 2).
			Render("↑/↓: navigate  •  /: filter  •  Enter: open  •  Esc: back  •  q: quit")

		return lipgloss.JoinVertical(lipgloss.Left, listView, helpText)
	}

	return listView
}

// View renders the current state of the application
func (m model) View() string {
	// If showing the file input
	if m.createFileInputVisible {
		return renderCreateNoteDialog(m.newFileInput, m.statusMessage, m.statusType)
	}

	// If editing a file
	if m.currentFile != nil {
		fileName := ViewTitleStyle.Render(fmt.Sprintf("✏️  Editing: %s", m.currentFile.Name()))

		help := ViewHelpStyle.Render("Ctrl+S: Save • Esc: Close without saving")

		return fmt.Sprintf("%s\n\n%s\n\n%s", fileName, m.textArea.View(), help)
	}

	// If showing the file list
	if m.showingList {
		return renderFileListView(m.fileList)
	}

	// Default: show landing page
	return renderLanding()
}
