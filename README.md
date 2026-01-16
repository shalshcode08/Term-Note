# Term-Note 🗒️

A terminal-based note-taking application built with Go and Bubble Tea.

## Features

- ✨ Create new markdown notes
- 📝 Edit notes with a full-screen text editor
- 📋 List and browse all your notes
- 💾 Save notes with Ctrl+S
- 🎨 Beautiful terminal UI with syntax highlighting

## Installation

```bash
# Clone the repository
git clone https://github.com/shalshcode08/Term-Note.git
cd term-note

# Build the application
make build

# Run the application
make run
```

## Usage

### Keyboard Shortcuts

- **Ctrl+N**: Create a new note
- **Ctrl+L**: List all notes
- **Ctrl+S**: Save current note
- **Esc**: Go back / Close current view
- **Ctrl+C** or **Q**: Quit application
- **Enter**: Confirm file name / Open selected note

### Workflow

1. Launch the application with `make run`
2. Press `Ctrl+N` to create a new note
3. Enter a filename (without extension, `.md` is added automatically)
4. Write your note in the text editor
5. Press `Ctrl+S` to save
6. Press `Ctrl+L` to list all notes
7. Use arrow keys to navigate and `Enter` to open a note

## Project Structure

The codebase is organized into separate files for better maintainability:

```
term-note/
├── main.go          # Application entry point
├── config.go        # Configuration and vault directory setup
├── model.go         # Bubble Tea model and state management
├── update.go        # Event handling and state updates
├── view.go          # UI rendering and styles
├── files.go         # File operations (listing, reading)
├── Makefile         # Build automation
├── go.mod           # Go module dependencies
└── README.md        # This file
```

### File Descriptions

- **main.go**: Minimal entry point that initializes and runs the Bubble Tea program
- **config.go**: Handles configuration initialization, sets up the `~/.termnote` vault directory
- **model.go**: Defines the application state struct, initializes UI components (text input, textarea, file list)
- **update.go**: Contains all the event handling logic for keyboard shortcuts and state transitions
- **view.go**: Renders the UI based on current state (welcome screen, file input, editor, file list)
- **files.go**: Utility functions for file operations like listing notes from the vault directory

## Data Storage

All notes are stored in `~/.termnote/` as markdown files with the `.md` extension.

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions

## Development

### Building

```bash
make build
```

### Running

```bash
make run
```

### Cleaning

```bash
make clean
```

## Architecture

The application follows the Elm Architecture pattern via Bubble Tea:

1. **Model**: Holds the application state
2. **Update**: Handles messages and updates the model
3. **View**: Renders the current state

This separation of concerns makes the code:
- Easy to test
- Simple to extend with new features
- Clear to understand and maintain
