package notes

import (
	"strings"
)

// InsertBulletPoint inserts a bullet point at the current line
func InsertBulletPoint(text string, cursorLine int) string {
	lines := strings.Split(text, "\n")
	if cursorLine >= len(lines) {
		cursorLine = len(lines) - 1
	}

	// Check if current line already has a bullet
	trimmed := strings.TrimSpace(lines[cursorLine])
	if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") {
		return text
	}

	// Add bullet point at the beginning of the line
	lines[cursorLine] = "- " + lines[cursorLine]
	return strings.Join(lines, "\n")
}

// InsertNumberedList inserts a numbered list item
func InsertNumberedList(text string, cursorLine int) string {
	lines := strings.Split(text, "\n")
	if cursorLine >= len(lines) {
		cursorLine = len(lines) - 1
	}

	// Find the number to use (look at previous line)
	number := 1
	if cursorLine > 0 {
		prevLine := strings.TrimSpace(lines[cursorLine-1])
		if len(prevLine) > 2 && prevLine[0] >= '0' && prevLine[0] <= '9' && prevLine[1] == '.' {
			// Extract number from previous line
			if n := int(prevLine[0] - '0'); n > 0 {
				number = n + 1
			}
		}
	}

	lines[cursorLine] = strings.Repeat(" ", 0) + string('0'+rune(number)) + ". " + lines[cursorLine]
	return strings.Join(lines, "\n")
}

// InsertTodo inserts a todo checkbox
func InsertTodo(text string, cursorLine int) string {
	lines := strings.Split(text, "\n")
	if cursorLine >= len(lines) {
		cursorLine = len(lines) - 1
	}

	trimmed := strings.TrimSpace(lines[cursorLine])
	if strings.HasPrefix(trimmed, "- [ ] ") || strings.HasPrefix(trimmed, "- [x] ") {
		return text
	}

	lines[cursorLine] = "- [ ] " + lines[cursorLine]
	return strings.Join(lines, "\n")
}

// ToggleTodo toggles a todo item between checked and unchecked
func ToggleTodo(text string, cursorLine int) string {
	lines := strings.Split(text, "\n")
	if cursorLine >= len(lines) {
		cursorLine = len(lines) - 1
	}

	line := lines[cursorLine]
	trimmed := strings.TrimSpace(line)

	if strings.HasPrefix(trimmed, "- [ ] ") {
		// Check it
		lines[cursorLine] = strings.Replace(line, "- [ ]", "- [x]", 1)
	} else if strings.HasPrefix(trimmed, "- [x] ") {
		// Uncheck it
		lines[cursorLine] = strings.Replace(line, "- [x]", "- [ ]", 1)
	}

	return strings.Join(lines, "\n")
}

// InsertHeader inserts a markdown header
func InsertHeader(text string, cursorLine int, level int) string {
	lines := strings.Split(text, "\n")
	if cursorLine >= len(lines) {
		cursorLine = len(lines) - 1
	}

	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}

	prefix := strings.Repeat("#", level) + " "
	lines[cursorLine] = prefix + lines[cursorLine]
	return strings.Join(lines, "\n")
}

// InsertTable inserts a simple markdown table
func InsertTable(rows, cols int) string {
	if rows < 2 {
		rows = 2
	}
	if cols < 2 {
		cols = 2
	}

	var table strings.Builder

	// Header row
	table.WriteString("|")
	for i := 0; i < cols; i++ {
		table.WriteString(" Header " + string('A'+rune(i)) + " |")
	}
	table.WriteString("\n")

	// Separator row
	table.WriteString("|")
	for i := 0; i < cols; i++ {
		table.WriteString("---------|")
	}
	table.WriteString("\n")

	// Data rows
	for r := 0; r < rows-1; r++ {
		table.WriteString("|")
		for c := 0; c < cols; c++ {
			table.WriteString(" Cell     |")
		}
		table.WriteString("\n")
	}

	return table.String()
}

// WrapSelection wraps selected text with markers (for bold, italic, code)
func WrapSelection(text string, marker string) string {
	// If text is already wrapped, unwrap it
	if strings.HasPrefix(text, marker) && strings.HasSuffix(text, marker) && len(text) > len(marker)*2 {
		return text[len(marker) : len(text)-len(marker)]
	}
	return marker + text + marker
}

// InsertCodeBlock inserts a code block with language
func InsertCodeBlock(language string) string {
	if language == "" {
		language = "go"
	}
	return "```" + language + "\n\n```"
}

// InsertHorizontalRule inserts a horizontal rule
func InsertHorizontalRule() string {
	return "---\n"
}

// InsertLink inserts a markdown link template
func InsertLink() string {
	return "[link text](url)"
}

// InsertImage inserts a markdown image template
func InsertImage() string {
	return "![alt text](image-url)"
}

// GetLineAtCursor returns the line number where cursor is positioned
func GetLineAtCursor(text string, cursorPos int) int {
	if cursorPos <= 0 {
		return 0
	}

	lines := strings.Split(text[:cursorPos], "\n")
	return len(lines) - 1
}
