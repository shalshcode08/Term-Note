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
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	textArea               textarea.Model
	fileList               list.Model
	showingList            bool
}

func initializeModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter file name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	ti.TextStyle = cursorStyle

	ta := textarea.New()
	ta.Placeholder = "Write your note here..."
	ta.Focus()
	ta.ShowLineNumbers = false

	notesList := listFiles()
	finalList := list.New(notesList, list.NewDefaultDelegate(), 0, 0)
	finalList.Title = "All Notes"
	finalList.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("254")).
		Padding(0, 1)

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		textArea:               ta,
		fileList:               finalList,
		showingList:            false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
