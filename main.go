package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	valutDir    string
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home dir", err)
	}

	valutDir = fmt.Sprintf("%s/.termnote", homeDir)
}

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	textArea               textarea.Model
}

func initializeModel() model {
	err := os.MkdirAll(valutDir, 0750)
	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	// initialize text-input with styles
	ti := textinput.New()
	ti.Placeholder = "Enter file name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	ti.TextStyle = cursorStyle

	// initilize textArea
	ta := textarea.New()
	ta.Placeholder = "Write your note here..."
	ta.Focus()
	ta.ShowLineNumbers = false

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		currentFile:            nil,
		textArea:               ta,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "ctrl+n":
			m.createFileInputVisible = true
			return m, nil

		case "ctrl+s":
			if m.currentFile == nil {
				break
			}
			if err := m.currentFile.Truncate(0); err != nil {
				fmt.Println("cannot save the file :(")
				return m, nil
			}

			if _, err := m.currentFile.Seek(0, 0); err != nil {
				fmt.Println("cannot save the file :(")
				return m, nil
			}

			if _, err := m.currentFile.WriteString(m.textArea.Value()); err != nil {
				fmt.Println("cannot save the file :(")
				return m, nil
			}

			if err := m.currentFile.Close(); err != nil {
				fmt.Println("cannot close the file")
			}

			m.currentFile = nil
			m.textArea.SetValue("")

			return m, nil

		case "enter":
			if m.currentFile != nil {
				break
			}

			filename := m.newFileInput.Value()
			if filename != "" {
				filePath := fmt.Sprintf("%s/%s.md", valutDir, filename)
				_, err := os.Stat(filePath)
				if err == nil {
					log.Fatalf("File already exists: %s", filePath)
					return m, nil
				}

				f, err := os.Create(filePath)
				if err != nil {
					log.Fatalf("%v", err)
				}

				m.currentFile = f
				m.createFileInputVisible = false
				m.newFileInput.SetValue("")
			}
			return m, nil
		}
	}

	if m.createFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
	}

	if m.currentFile != nil {
		m.textArea, cmd = m.textArea.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).
		PaddingLeft(2).
		PaddingRight(2)

	welcome := style.Render("Welcome to TERMNOTE 🗒️")

	helpKeys := "Ctrl+N : new file . Ctrl+L: list . Esc: back/save . Ctrl+S: save . Ctrl+Q: quit"

	view := ""
	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.textArea.View()
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, helpKeys)

}

func main() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
