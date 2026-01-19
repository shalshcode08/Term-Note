package main

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	docStyle    = lipgloss.NewStyle().Margin(1, 2)
)

type item struct {
	title, desc string
	filename    string // Full filename with extension
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }
func (i item) Filename() string    { return i.filename }

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	textArea               textarea.Model
	fileList               list.Model
	showingList            bool
	statusMessage          string
	statusType             string // "success", "error", "warning", ""
	showHelp               bool   // Toggle help overlay
}

func initializeModel() model {
	ti := textinput.New()
	ti.Placeholder = "my-awesome-note"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 56
	ti.Prompt = "" // Hide default prompt, we'll add custom one in view
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(ColorPrimary)
	ti.PromptStyle = lipgloss.NewStyle().Foreground(ColorPrimary).Bold(true)
	ti.TextStyle = lipgloss.NewStyle().Foreground(ColorText)
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(ColorMuted).Italic(true)

	ta := textarea.New()
	ta.Placeholder = "Start writing your note..."
	ta.Focus()
	ta.ShowLineNumbers = false
	ta.CharLimit = 0 // No limit
	ta.SetWidth(80)
	ta.SetHeight(20)
	ta.Prompt = "" // Remove prompt to eliminate left line
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.FocusedStyle.Placeholder = lipgloss.NewStyle().Foreground(ColorMuted)
	ta.FocusedStyle.Text = lipgloss.NewStyle().Foreground(ColorText)
	ta.FocusedStyle.Prompt = lipgloss.NewStyle()
	ta.FocusedStyle.EndOfBuffer = lipgloss.NewStyle()
	ta.FocusedStyle.LineNumber = lipgloss.NewStyle()
	ta.BlurredStyle.Placeholder = lipgloss.NewStyle().Foreground(ColorMuted)
	ta.BlurredStyle.Text = lipgloss.NewStyle().Foreground(ColorText)
	ta.BlurredStyle.Prompt = lipgloss.NewStyle()
	ta.BlurredStyle.EndOfBuffer = lipgloss.NewStyle()
	ta.BlurredStyle.LineNumber = lipgloss.NewStyle()

	notesList := listFiles()
	finalList := list.New(notesList, list.NewDefaultDelegate(), 0, 0)
	finalList.Title = "All Notes"
	finalList.Styles.Title = lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true)
	finalList.Styles.FilterPrompt = lipgloss.NewStyle().Foreground(ColorPrimary).Bold(true)
	finalList.Styles.FilterCursor = lipgloss.NewStyle().Foreground(ColorPrimary)
	finalList.SetShowStatusBar(true)
	finalList.SetFilteringEnabled(true)
	finalList.Styles.StatusBar = lipgloss.NewStyle().
		Foreground(ColorMuted).
		Padding(0, 2)
	finalList.Styles.HelpStyle = ListHelpStyle
	finalList.SetStatusBarItemName("note", "notes")
	finalList.SetShowHelp(false) // Disable default help, we have custom help text

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		textArea:               ta,
		fileList:               finalList,
		showingList:            false,
		statusMessage:          "",
		statusType:             "",
		showHelp:               false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
