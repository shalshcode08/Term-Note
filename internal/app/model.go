package app

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/shalshcode08/Term-Note/internal/config"
	"github.com/shalshcode08/Term-Note/internal/notes"
	"github.com/shalshcode08/Term-Note/internal/ui/styles"
)

var (
	DocStyle = lipgloss.NewStyle().Margin(1, 2)
)

// Model represents the application state
type Model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	textArea               textarea.Model
	fileList               list.Model
	showingList            bool
	statusMessage          string
	statusType             string // "success", "error", "warning", ""
	showHelp               bool   // Toggle help overlay
	showDeleteConfirm      bool   // Show delete confirmation dialog
	fileToDelete           string // Filename to delete
	windowWidth            int    // Terminal window width
	windowHeight           int    // Terminal window height
}

// New creates and initializes a new application model
func New() Model {
	ti := textinput.New()
	ti.Placeholder = "my-awesome-note"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 56
	ti.Prompt = "" // Hide default prompt, we'll add custom one in view
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(styles.ColorPrimary)
	ti.PromptStyle = lipgloss.NewStyle().Foreground(styles.ColorPrimary).Bold(true)
	ti.TextStyle = lipgloss.NewStyle().Foreground(styles.ColorText)
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(styles.ColorMuted).Italic(true)

	ta := textarea.New()
	ta.Placeholder = "Start writing your note..."
	ta.Focus()
	ta.ShowLineNumbers = false
	ta.CharLimit = 0 // No limit
	ta.SetWidth(80)
	ta.SetHeight(20)
	ta.Prompt = "" // Remove prompt to eliminate left line
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.FocusedStyle.Placeholder = lipgloss.NewStyle().Foreground(styles.ColorMuted)
	ta.FocusedStyle.Text = lipgloss.NewStyle().Foreground(styles.ColorText)
	ta.FocusedStyle.Prompt = lipgloss.NewStyle()
	ta.FocusedStyle.EndOfBuffer = lipgloss.NewStyle()
	ta.FocusedStyle.LineNumber = lipgloss.NewStyle()
	ta.BlurredStyle.Placeholder = lipgloss.NewStyle().Foreground(styles.ColorMuted)
	ta.BlurredStyle.Text = lipgloss.NewStyle().Foreground(styles.ColorText)
	ta.BlurredStyle.Prompt = lipgloss.NewStyle()
	ta.BlurredStyle.EndOfBuffer = lipgloss.NewStyle()
	ta.BlurredStyle.LineNumber = lipgloss.NewStyle()

	notesList := notes.ListFiles(config.VaultDir)
	finalList := list.New(notesList, list.NewDefaultDelegate(), 0, 0)
	finalList.Title = "All Notes"
	finalList.Styles.Title = lipgloss.NewStyle().
		Foreground(styles.ColorPrimary).
		Bold(true)
	finalList.Styles.FilterPrompt = lipgloss.NewStyle().Foreground(styles.ColorPrimary).Bold(true)
	finalList.Styles.FilterCursor = lipgloss.NewStyle().Foreground(styles.ColorPrimary)
	finalList.SetShowStatusBar(true)
	finalList.SetFilteringEnabled(true)
	finalList.Styles.StatusBar = lipgloss.NewStyle().
		Foreground(styles.ColorMuted).
		Padding(0, 2)
	finalList.Styles.HelpStyle = styles.ListHelpStyle
	finalList.SetStatusBarItemName("note", "notes")
	finalList.SetShowHelp(false) // Disable default help, we have custom help text

	return Model{
		newFileInput:           ti,
		createFileInputVisible: false,
		textArea:               ta,
		fileList:               finalList,
		showingList:            false,
		statusMessage:          "",
		statusType:             "",
		showHelp:               false,
		showDeleteConfirm:      false,
		fileToDelete:           "",
		windowWidth:            80,
		windowHeight:           24,
	}
}

// Init initializes the model (Bubble Tea interface)
func (m Model) Init() tea.Cmd {
	return nil
}
