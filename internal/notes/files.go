package notes

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
)

// Item represents a file list item
type Item struct {
	title, desc string
	filename    string // Full filename with extension
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }
func (i Item) Filename() string    { return i.filename }

// formatRelativeTime returns a human-readable relative time string
func formatRelativeTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	// Future time (shouldn't happen, but just in case)
	if diff < 0 {
		return "just now"
	}

	// Less than a minute
	if diff < time.Minute {
		return "just now"
	}

	// Less than an hour
	if diff < time.Hour {
		mins := int(diff.Minutes())
		if mins == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", mins)
	}

	// Less than a day
	if diff < 24*time.Hour {
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	}

	// Less than a week
	if diff < 7*24*time.Hour {
		days := int(diff.Hours() / 24)
		if days == 1 {
			return "yesterday"
		}
		return fmt.Sprintf("%d days ago", days)
	}

	// Less than a month
	if diff < 30*24*time.Hour {
		weeks := int(diff.Hours() / 24 / 7)
		if weeks == 1 {
			return "1 week ago"
		}
		return fmt.Sprintf("%d weeks ago", weeks)
	}

	// Less than a year
	if diff < 365*24*time.Hour {
		months := int(diff.Hours() / 24 / 30)
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	}

	// Over a year
	years := int(diff.Hours() / 24 / 365)
	if years == 1 {
		return "1 year ago"
	}
	return fmt.Sprintf("%d years ago", years)
}

// fileWithTime holds file info with modification time for sorting
type fileWithTime struct {
	name    string
	modTime time.Time
}

// ListFiles returns a list of all note files in the vault directory
func ListFiles(vaultDir string) []list.Item {
	items := make([]list.Item, 0)
	filesWithTime := make([]fileWithTime, 0)

	entries, err := os.ReadDir(vaultDir)
	if err != nil {
		log.Fatal("error reading notes list")
	}

	// Collect files with their modification times
	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			filesWithTime = append(filesWithTime, fileWithTime{
				name:    entry.Name(),
				modTime: info.ModTime(),
			})
		}
	}

	// Sort by modification time (most recent first)
	for i := 0; i < len(filesWithTime); i++ {
		for j := i + 1; j < len(filesWithTime); j++ {
			if filesWithTime[j].modTime.After(filesWithTime[i].modTime) {
				filesWithTime[i], filesWithTime[j] = filesWithTime[j], filesWithTime[i]
			}
		}
	}

	// Create list items from sorted files
	for _, file := range filesWithTime {
		items = append(items, Item{
			title:    file.name,
			desc:     fmt.Sprintf("Modified: %s", file.modTime.Format("2006-01-02 15:04")),
			filename: file.name, // Store full filename for opening
		})
	}

	return items
}
