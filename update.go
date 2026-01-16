package main

import (
	"fmt"
	"log"
	"os"

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
