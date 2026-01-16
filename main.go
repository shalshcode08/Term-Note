package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	valutDir    string
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	docStyle    = lipgloss.NewStyle().Margin(1, 2)
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home dir", err)
	}

	valutDir = fmt.Sprintf("%s/.termnote", homeDir)
}

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

	// list initialize
	notesList := listFiles()

	finalList := list.New(notesList, list.NewDefaultDelegate(), 0, 0)
	finalList.Title = "All Notes"
	finalList.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("16")).Background(lipgloss.Color("254")).Padding(0, 1)

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		textArea:               ta,
		fileList:               finalList,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.fileList.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "ctrl+n":
			m.createFileInputVisible = true
			return m, nil

		case "ctrl+l":
			notesList := listFiles()
			m.fileList.SetItems(notesList)
			m.showingList = true
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

		case "esc":
			if m.createFileInputVisible {
				m.createFileInputVisible = false
			}

			if m.currentFile != nil {
				m.textArea.SetValue("")
				m.currentFile = nil
			}

			if m.showingList {
				if m.fileList.FilterState() == list.Filtering {
					break
				}
				m.showingList = false
			}

			return m, nil

		case "enter":
			if m.currentFile != nil {
				break
			}

			if m.showingList {
				item, ok := m.fileList.SelectedItem().(item)
				if ok {
					filepath := fmt.Sprintf("%s/%s", valutDir, item.title)
					content, err := os.ReadFile(filepath)
					if err != nil {
						fmt.Printf("cannot read the file: %v", err)
						return m, nil
					}
					m.textArea.SetValue(string(content))

					file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
					if err != nil {
						fmt.Printf("cannot read the file: %v", err)
						return m, nil
					}

					m.currentFile = file
					m.showingList = false
				}
				return m, nil
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

	if m.showingList {
		m.fileList, cmd = m.fileList.Update(msg)
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

func main() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

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

			modifiedTime := info.ModTime().Format("2006-02-02 15:04")
			items = append(items, item{
				title: entry.Name(),
				desc:  fmt.Sprintf("Modified: %s", modifiedTime),
			})
		}
	}
	return items
}
