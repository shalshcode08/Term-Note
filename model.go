package main

import (
	"fmt"
	"io"
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

// customDelegate is a custom list item delegate with enhanced styling
type customDelegate struct{}

func (d customDelegate) Height() int                               { return 2 }
func (d customDelegate) Spacing() int                              { return 1 }
func (d customDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d customDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	// Get styles based on selection
	var title, desc string
	if index == m.Index() {
		// Selected item
		title = ListItemSelectedTitleStyle.Render(i.Title())
		desc = ListItemSelectedDescStyle.Render("  " + i.Description())
	} else {
		// Normal item
		title = ListItemTitleStyle.Render(i.Title())
		desc = ListItemDescStyle.Render("  " + i.Description())
	}

	fmt.Fprintf(w, "%s\n%s", title, desc)
}

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	textArea               textarea.Model
	fileList               list.Model
	showingList            bool
	statusMessage          string
	statusType             string // "success", "error", "warning", ""
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
	ta.Placeholder = "Write your note here..."
	ta.Focus()
	ta.ShowLineNumbers = false

	notesList := listFiles()
	delegate := customDelegate{}
	finalList := list.New(notesList, delegate, 0, 0)
	finalList.Title = "📋 All Notes"
	finalList.Styles.Title = ListTitleStyle
	finalList.Styles.FilterPrompt = lipgloss.NewStyle().Foreground(ColorPrimary).Bold(true)
	finalList.Styles.FilterCursor = lipgloss.NewStyle().Foreground(ColorPrimary)
	finalList.SetShowStatusBar(true)
	finalList.SetFilteringEnabled(true)
	finalList.Styles.StatusBar = lipgloss.NewStyle().
		Foreground(ColorMuted).
		Padding(0, 2)
	finalList.Styles.HelpStyle = ListHelpStyle
	finalList.SetStatusBarItemName("note", "notes")

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		textArea:               ta,
		fileList:               finalList,
		showingList:            false,
		statusMessage:          "",
		statusType:             "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
