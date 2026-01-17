package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

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
			m.statusMessage = ""
			m.statusType = ""
			m.newFileInput.SetValue("")
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
				m.statusMessage = ""
				m.statusType = ""
				m.newFileInput.SetValue("")
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
					filepath := fmt.Sprintf("%s/%s", valutDir, item.filename)
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

			if m.createFileInputVisible {
				filename := strings.TrimSpace(m.newFileInput.Value())

				// Validate filename
				if filename == "" {
					m.statusMessage = "Please enter a note name"
					m.statusType = "error"
					return m, nil
				}

				// Check for invalid characters
				invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
				for _, char := range invalidChars {
					if strings.Contains(filename, char) {
						m.statusMessage = "Filename contains invalid characters"
						m.statusType = "error"
						return m, nil
					}
				}

				filePath := fmt.Sprintf("%s/%s.md", valutDir, filename)

				// Check if file already exists
				_, err := os.Stat(filePath)
				if err == nil {
					m.statusMessage = "File already exists with this name"
					m.statusType = "error"
					return m, nil
				}

				// Create the file
				f, err := os.Create(filePath)
				if err != nil {
					m.statusMessage = fmt.Sprintf("Failed to create file: %v", err)
					m.statusType = "error"
					return m, nil
				}

				m.currentFile = f
				m.createFileInputVisible = false
				m.newFileInput.SetValue("")
				m.statusMessage = ""
				m.statusType = ""
				return m, nil
			}
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
