# TermNote

A minimal terminal-based note-taking application built with Go and Bubble Tea.

## Features

- Create and edit Markdown notes
- List and browse all notes
- Markdown formatting shortcuts (bullets, todos, headers, tables)
- Delete notes with confirmation
- Full-text editing with syntax support
- Auto-save functionality
- Keyboard-driven interface

## Installation

```bash
git clone https://github.com/shalshcode08/Term-Note.git
cd term-note
make build
```

## Usage

```bash
make run
```

### Keyboard Shortcuts

#### Basic Commands
- `Ctrl+N` - Create new note
- `Ctrl+L` - List all notes
- `Ctrl+S` - Save current note
- `Ctrl+H` - Show help menu
- `Esc` - Go back / Close current view
- `q` - Quit application

#### Markdown Formatting
- `Ctrl+B` - Insert bullet point
- `Ctrl+T` - Insert todo checkbox
- `Ctrl+D` - Toggle todo (check/uncheck)
- `Ctrl+1/2/3` - Insert headers (H1, H2, H3)

#### Advanced Features
- `Alt+T` - Insert table
- `Alt+C` - Insert code block
- `Alt+L` - Insert link template
- `Alt+I` - Insert image template
- `Alt+R` - Insert horizontal rule

#### File Management
- `Enter` - Open selected note
- `d` - Delete selected note
- `/` - Filter notes

## Data Storage

All notes are stored as Markdown files in `~/.termnote/`

## Development

### Build
```bash
make build
```

### Run
```bash
make run
```

### Clean
```bash
make clean
```

## Project Structure

```
term-note/
├── main.go              # Application entry point
├── Makefile             # Build automation
└── internal/            # Internal packages
    ├── app/             # Core application logic
    ├── config/          # Configuration management
    ├── notes/           # Note operations
    └── ui/              # User interface components
```

See `PROJECT_STRUCTURE.md` for detailed documentation.

## Architecture

TermNote follows the Elm Architecture pattern via Bubble Tea:
- Model: Application state
- Update: Event handling and state updates
- View: UI rendering

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions
