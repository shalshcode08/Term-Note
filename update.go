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
		// Store window dimensions for centering dialogs
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height

		h, v := docStyle.GetFrameSize()
		m.fileList.SetSize(msg.Width-h, msg.Height-v)
		// Resize textarea for editor view - use full window
		m.textArea.SetWidth(msg.Width)
		m.textArea.SetHeight(msg.Height - 4) // Leave space for header and status bar

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
			m.statusMessage = ""
			m.statusType = ""
			return m, nil

		case "d", "delete":
			// Delete note - only works in list view
			if m.showingList && !m.showDeleteConfirm {
				selectedItem, ok := m.fileList.SelectedItem().(item)
				if ok {
					m.fileToDelete = selectedItem.Filename()
					m.showDeleteConfirm = true
				}
			}
			return m, nil

		case "y":
			// Confirm delete
			if m.showDeleteConfirm {
				filePath := fmt.Sprintf("%s/%s", valutDir, m.fileToDelete)
				if err := os.Remove(filePath); err != nil {
					m.statusMessage = "Failed to delete note"
					m.statusType = "error"
				} else {
					m.statusMessage = "Note deleted successfully"
					m.statusType = "success"
					// Refresh the list
					notesList := listFiles()
					m.fileList.SetItems(notesList)
				}
				m.showDeleteConfirm = false
				m.fileToDelete = ""
			}
			return m, nil

		case "n":
			// Cancel delete
			if m.showDeleteConfirm {
				m.showDeleteConfirm = false
				m.fileToDelete = ""
			}
			return m, nil

		case "ctrl+s":
			if m.currentFile == nil {
				break
			}
			// Save the file content
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

			// Sync to disk but keep file open
			if err := m.currentFile.Sync(); err != nil {
				fmt.Println("cannot sync the file")
			}

			// Don't close file, don't clear textarea - just save and continue editing
			return m, nil

		case "esc":
			if m.showDeleteConfirm {
				m.showDeleteConfirm = false
				m.fileToDelete = ""
				return m, nil
			}

			if m.createFileInputVisible {
				m.createFileInputVisible = false
				m.statusMessage = ""
				m.statusType = ""
				m.newFileInput.SetValue("")
			}

			if m.showHelp {
				m.showHelp = false
				return m, nil
			}

			if m.currentFile != nil {
				// Close the file before exiting
				if err := m.currentFile.Close(); err != nil {
					fmt.Println("error closing file")
				}
				m.currentFile = nil
				m.textArea.SetValue("")
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
		// Check for markdown shortcuts before passing to textarea
		if msg, ok := msg.(tea.KeyMsg); ok {
			switch msg.String() {
			case "ctrl+h":
				// Toggle help menu
				m.showHelp = !m.showHelp
				return m, nil
			case "ctrl+b":
				// Toggle bullet point - insert at cursor
				m.textArea.InsertString("- ")
				return m, nil
			case "ctrl+t":
				// Insert todo checkbox
				m.textArea.InsertString("- [ ] ")
				return m, nil
			case "ctrl+d":
				// Toggle todo checkbox (check/uncheck) on current cursor line
				// Get current line number from textarea
				currentLine := m.textArea.Line()
				newText := toggleTodo(m.textArea.Value(), currentLine)
				m.textArea.SetValue(newText)
				return m, nil
			case "ctrl+1":
				// Insert H1 header
				m.textArea.InsertString("# ")
				return m, nil
			case "ctrl+2":
				// Insert H2 header
				m.textArea.InsertString("## ")
				return m, nil
			case "ctrl+3":
				// Insert H3 header
				m.textArea.InsertString("### ")
				return m, nil
			case "alt+t":
				// Insert table
				m.textArea.InsertString(insertTable(3, 3))
				return m, nil
			case "alt+c":
				// Insert code block
				m.textArea.InsertString(insertCodeBlock(""))
				return m, nil
			case "alt+l":
				// Insert link
				m.textArea.InsertString(insertLink())
				return m, nil
			case "alt+i":
				// Insert image
				m.textArea.InsertString(insertImage())
				return m, nil
			case "alt+r":
				// Insert horizontal rule
				m.textArea.InsertString(insertHorizontalRule())
				return m, nil
			case "enter":
				// Auto-continue lists on Enter
				text := m.textArea.Value()
				lines := strings.Split(text, "\n")
				if len(lines) > 0 {
					lastLine := strings.TrimRight(lines[len(lines)-1], " \t")

					// Check for todo items FIRST (before bullets)
					if strings.HasPrefix(lastLine, "- [ ] ") && len(lastLine) > 6 {
						m.textArea.InsertString("\n- [ ] ")
						return m, nil
					}
					if strings.HasPrefix(lastLine, "- [x] ") && len(lastLine) > 6 {
						m.textArea.InsertString("\n- [ ] ")
						return m, nil
					}

					// Check for bullet points (after todos)
					if strings.HasPrefix(lastLine, "- ") && len(lastLine) > 2 {
						m.textArea.InsertString("\n- ")
						return m, nil
					}
					if strings.HasPrefix(lastLine, "* ") && len(lastLine) > 2 {
						m.textArea.InsertString("\n* ")
						return m, nil
					}

					// Check for numbered lists
					if len(lastLine) > 2 && lastLine[0] >= '0' && lastLine[0] <= '9' && lastLine[1] == '.' && lastLine[2] == ' ' {
						nextNum := int(lastLine[0]-'0') + 1
						if nextNum <= 9 {
							m.textArea.InsertString(fmt.Sprintf("\n%d. ", nextNum))
							return m, nil
						}
					}
				}
				// If not a list, just add newline normally
				m.textArea.InsertString("\n")
				return m, nil
			}
		}
		m.textArea, cmd = m.textArea.Update(msg)
	}

	if m.showingList {
		m.fileList, cmd = m.fileList.Update(msg)
	}

	return m, cmd
}
