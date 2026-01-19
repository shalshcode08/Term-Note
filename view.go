package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
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

// renderDeleteConfirm renders the delete confirmation dialog
func renderDeleteConfirm(filename string) string {
	dialogStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorError).
		Padding(2, 4).
		Width(60)

	titleStyle := lipgloss.NewStyle().
		Foreground(ColorError).
		Bold(true).
		Align(lipgloss.Center).
		Width(52)

	filenameStyle := lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true).
		Align(lipgloss.Center).
		Width(52)

	messageStyle := lipgloss.NewStyle().
		Foreground(ColorMuted).
		Align(lipgloss.Center).
		Width(52).
		MarginTop(1)

	buttonsStyle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(52).
		MarginTop(2)

	title := titleStyle.Render("⚠️  DELETE NOTE")
	file := filenameStyle.Render(filename)
	message := messageStyle.Render("Are you sure you want to delete this note?\nThis action cannot be undone.")

	yesButton := lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).
		Background(ColorError).
		Bold(true).
		Padding(0, 2).
		Render(" Yes (Y) ")

	noButton := lipgloss.NewStyle().
		Foreground(ColorText).
		Background(ColorMuted).
		Bold(true).
		Padding(0, 2).
		Render(" No (N) ")

	buttons := buttonsStyle.Render(yesButton + "  " + noButton)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		file,
		message,
		buttons,
	)

	return dialogStyle.Render(content)
}

// renderFileListView renders the file list with enhanced styling
func renderFileListView(fileList list.Model, showDeleteConfirm bool, fileToDelete string) string {
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
			Render("↑/↓: navigate  •  /: filter  •  Enter: open  •  d: delete  •  Esc: back  •  q: quit")

		listView = lipgloss.JoinVertical(lipgloss.Left, listView, helpText)
	}

	return listView
}

// renderFileListViewWithStatus renders the file list with status messages
func renderFileListViewWithStatus(fileList list.Model, showDeleteConfirm bool, fileToDelete string, statusMessage string, statusType string, windowWidth int, windowHeight int) string {
	listView := renderFileListView(fileList, false, "")

	// Show delete confirmation overlay if active - do this FIRST to center it
	if showDeleteConfirm {
		confirmDialog := renderDeleteConfirm(fileToDelete)
		return lipgloss.Place(
			windowWidth, windowHeight,
			lipgloss.Center,
			lipgloss.Center,
			confirmDialog,
		)
	}

	// Show status message if present
	if statusMessage != "" {
		var statusStyle lipgloss.Style
		switch statusType {
		case "success":
			statusStyle = SuccessStyle
		case "error":
			statusStyle = ErrorStyle
		case "warning":
			statusStyle = WarningStyle
		default:
			statusStyle = lipgloss.NewStyle().Foreground(ColorText)
		}

		statusBar := lipgloss.NewStyle().
			Padding(0, 2).
			Render(statusStyle.Render(statusMessage))

		return lipgloss.JoinVertical(lipgloss.Left, statusBar, listView)
	}

	return listView
}

// renderHelpOverlay renders the help menu with all shortcuts
func renderHelpOverlay() string {
	helpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorPrimary).
		Padding(2, 4).
		Width(66)

	titleStyle := lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true).
		Align(lipgloss.Center).
		Width(58)

	sectionStyle := lipgloss.NewStyle().
		Foreground(ColorAccent).
		Bold(true).
		MarginTop(1)

	title := titleStyle.Render("⌨️  KEYBOARD SHORTCUTS")

	// Format shortcuts with proper alignment
	shortcuts := []struct {
		section string
		items   [][2]string
	}{
		{
			section: "Basic Commands:",
			items: [][2]string{
				{"Ctrl+S", "Save note"},
				{"Ctrl+H", "Toggle this help"},
				{"Esc   ", "Close without saving"},
			},
		},
		{
			section: "Markdown Formatting:",
			items: [][2]string{
				{"Ctrl+B", "Insert bullet point (- )"},
				{"Ctrl+T", "Insert todo checkbox (- [ ] )"},
				{"Ctrl+D", "Toggle todo (check/uncheck)"},
				{"Ctrl+1", "Insert H1 header (# )"},
				{"Ctrl+2", "Insert H2 header (## )"},
				{"Ctrl+3", "Insert H3 header (### )"},
			},
		},
		{
			section: "Advanced Features:",
			items: [][2]string{
				{"Alt+T ", "Insert table"},
				{"Alt+C ", "Insert code block"},
				{"Alt+L ", "Insert link template"},
				{"Alt+I ", "Insert image template"},
				{"Alt+R ", "Insert horizontal rule"},
			},
		},
	}

	var sections []string
	for _, s := range shortcuts {
		sections = append(sections, "")
		sections = append(sections, sectionStyle.Render(s.section))
		for _, item := range s.items {
			key := lipgloss.NewStyle().Foreground(ColorText).Bold(true).Render(item[0])
			desc := lipgloss.NewStyle().Foreground(ColorMuted).Render("  " + item[1])
			sections = append(sections, key+desc)
		}
	}

	closeHint := lipgloss.NewStyle().
		Foreground(ColorMuted).
		Italic(true).
		Align(lipgloss.Center).
		Width(58).
		MarginTop(2).
		Render("Press Ctrl+H or Esc to close")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		strings.Join(sections, "\n"),
		"",
		closeHint,
	)

	return helpStyle.Render(content)
}

// renderEditorView renders the note editing interface
func renderEditorView(currentFile *os.File, textArea textarea.Model, showHelp bool) string {
	// Extract just the filename from the full path
	fullPath := currentFile.Name()
	fileName := fullPath
	if idx := strings.LastIndex(fullPath, "/"); idx != -1 {
		fileName = fullPath[idx+1:]
	}

	// Header section with file info
	headerStyle := lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true)

	header := headerStyle.Render("✏️  " + fileName)

	// Editor without border - clean and minimal
	editor := textArea.View()

	// Status bar at bottom with markdown shortcuts
	statusBarStyle := lipgloss.NewStyle().
		Foreground(ColorMuted)

	// Main commands
	statusLeft := lipgloss.NewStyle().Foreground(ColorText).Render("Ctrl+S")
	statusLeftDesc := lipgloss.NewStyle().Foreground(ColorMuted).Render(" Save")

	statusRight := lipgloss.NewStyle().Foreground(ColorText).Render("  •  Esc")
	statusRightDesc := lipgloss.NewStyle().Foreground(ColorMuted).Render(" Close")

	// Help hint
	helpHint := lipgloss.NewStyle().Foreground(ColorMuted).Render("  •  Ctrl+H Help")

	statusBar := statusBarStyle.Render(statusLeft + statusLeftDesc + statusRight + statusRightDesc + helpHint)

	// Combine all parts
	view := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		"",
		editor,
		"",
		statusBar,
	)

	// Overlay help if toggled
	if showHelp {
		helpOverlay := renderHelpOverlay()
		// Place help overlay centered on top of the editor view
		return lipgloss.Place(
			lipgloss.Width(view),
			lipgloss.Height(view),
			lipgloss.Center,
			lipgloss.Center,
			helpOverlay,
			lipgloss.WithWhitespaceChars(" "),
			lipgloss.WithWhitespaceForeground(lipgloss.Color("0")),
		)
	}

	return view
}

// View renders the current state of the application
func (m model) View() string {
	// If showing the file input
	if m.createFileInputVisible {
		return renderCreateNoteDialog(m.newFileInput, m.statusMessage, m.statusType)
	}

	// If editing a file
	if m.currentFile != nil {
		return renderEditorView(m.currentFile, m.textArea, m.showHelp)
	}

	// If showing the file list
	if m.showingList {
		return renderFileListViewWithStatus(m.fileList, m.showDeleteConfirm, m.fileToDelete, m.statusMessage, m.statusType, m.windowWidth, m.windowHeight)
	}

	// Default: show landing page
	return renderLanding()
}
