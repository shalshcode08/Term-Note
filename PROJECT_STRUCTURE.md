# Project Structure

## Overview

TermNote is organized into a clean, modular architecture following Go best practices. Each component has a clear responsibility and can be modified independently.

## Directory Structure

```
term-note/
├── main.go                          # Application entry point
├── Makefile                         # Build automation
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
├── README.md                        # User documentation
├── PROJECT_STRUCTURE.md             # This file
│
└── internal/                        # Internal packages (not importable by other projects)
    │
    ├── app/                         # Core application logic
    │   ├── model.go                 # Application state (Bubble Tea model)
    │   ├── update.go                # Event handling and state updates
    │   └── view.go                  # UI rendering and all view functions
    │
    ├── config/                      # Configuration management
    │   └── config.go                # Vault directory setup and initialization
    │
    ├── notes/                       # Note operations
    │   ├── files.go                 # File listing, reading, and management
    │   └── markdown.go              # Markdown formatting helpers
    │
    └── ui/                          # User interface components
        └── styles/                  # Visual styling
            └── styles.go            # Colors, ASCII art, and style definitions
```

## Package Responsibilities

### `main.go`
**Purpose**: Application entry point  
**Responsibilities**:
- Initialize the Bubble Tea program
- Create the application model
- Handle top-level errors

**When to modify**:
- Changing program initialization
- Adding command-line flags
- Modifying error handling at startup

---

### `internal/app/`
**Purpose**: Core application logic and state management

#### `model.go`
**Responsibilities**:
- Define the application state (`Model` struct)
- Initialize all UI components (textarea, textinput, list)
- Set up default values and configurations

**Key exports**:
- `Model` - Main application state
- `New()` - Model initialization
- `Init()` - Bubble Tea initialization

**When to modify**:
- Adding new state fields
- Changing component initialization
- Modifying default settings

#### `update.go`
**Responsibilities**:
- Handle all user input and events
- Update application state based on events
- Coordinate between different views and modes
- Implement keyboard shortcuts
- Handle file operations (create, save, delete)

**Key function**:
- `Update(msg tea.Msg)` - Main event handler

**When to modify**:
- Adding new keyboard shortcuts
- Implementing new features that respond to user input
- Changing navigation between views
- Adding new event handlers

#### `view.go`
**Responsibilities**:
- Render all application views
- Compose UI elements using lipgloss
- Generate visual layouts for each mode

**Key function**:
- `View()` - Main render function

**Render functions**:
- `renderLanding()` - Landing page with logo and shortcuts
- `renderCreateNoteDialog()` - File creation dialog
- `renderEditorView()` - Note editing interface
- `renderFileListViewWithStatus()` - File list with status messages
- `renderHelpOverlay()` - Keyboard shortcuts help
- `renderDeleteConfirm()` - Delete confirmation dialog

**When to modify**:
- Changing visual appearance
- Adding new views or dialogs
- Modifying layout and styling
- Updating help text

---

### `internal/config/`
**Purpose**: Application configuration

#### `config.go`
**Responsibilities**:
- Set up vault directory (`~/.termnote`)
- Initialize configuration on startup
- Provide access to configuration values

**Key exports**:
- `VaultDir` - Path to notes directory
- `InitConfig()` - Initialize configuration

**When to modify**:
- Changing default directories
- Adding new configuration options
- Modifying initialization logic

---

### `internal/notes/`
**Purpose**: Note and file management

#### `files.go`
**Responsibilities**:
- List all notes in vault directory
- Sort files by modification time
- Format file metadata for display
- Provide file information to UI

**Key exports**:
- `Item` - File list item type
- `ListFiles(vaultDir)` - Get all notes

**When to modify**:
- Changing file listing logic
- Adding file metadata
- Modifying sort order
- Adding file filtering

#### `markdown.go`
**Responsibilities**:
- Insert markdown formatting elements
- Handle list continuation
- Toggle todo checkboxes
- Generate markdown templates

**Key exports**:
- `InsertBulletPoint()` - Add bullet points
- `InsertTodo()` - Add todo checkboxes
- `ToggleTodo()` - Toggle checkbox state
- `InsertHeader()` - Add headers (H1, H2, H3)
- `InsertTable()` - Generate table template
- `InsertCodeBlock()` - Add code blocks
- `InsertLink()` - Add link template
- `InsertImage()` - Add image template
- `InsertHorizontalRule()` - Add horizontal rule

**When to modify**:
- Adding new markdown features
- Changing formatting behavior
- Implementing auto-formatting
- Adding smart text manipulation

---

### `internal/ui/styles/`
**Purpose**: Visual styling and theming

#### `styles.go`
**Responsibilities**:
- Define color scheme
- Provide reusable style components
- Store ASCII art variations
- Export style constants for consistency

**Key exports**:

**Colors**:
- `ColorPrimary` - Pink/Magenta (main accent)
- `ColorSecondary` - Purple
- `ColorAccent` - Light pink
- `ColorText` - Light gray
- `ColorMuted` - Gray
- `ColorSuccess` - Green
- `ColorWarning` - Orange
- `ColorError` - Red

**Styles**:
- `HeaderStyle` - Page headers
- `HelpTitleStyle` - Help section titles
- `DialogBoxStyle` - Dialog containers
- `ListTitleStyle` - List headers
- `ErrorStyle`, `SuccessStyle`, `WarningStyle` - Status messages

**ASCII Art**:
- Multiple header styles (box, simple, slant, etc.)
- `GetASCIIArt(style)` - Retrieve art by name

**When to modify**:
- Changing color scheme
- Adding new styles
- Creating new ASCII art
- Updating visual consistency

---

## Adding New Features

### Adding a New View
1. Create render function in `internal/app/view.go`
2. Add state field in `internal/app/model.go`
3. Add navigation logic in `internal/app/update.go`
4. Add styles in `internal/ui/styles/styles.go` if needed

### Adding a New Keyboard Shortcut
1. Add handler in `internal/app/update.go` (in the appropriate section)
2. Update help overlay in `internal/app/view.go` (`renderHelpOverlay()`)
3. Add any new functions to `internal/notes/` if markdown-related

### Adding File Operations
1. Add function to `internal/notes/files.go`
2. Call from `internal/app/update.go` event handlers
3. Update UI in `internal/app/view.go` if needed

### Changing Colors or Styles
1. Modify `internal/ui/styles/styles.go` color constants
2. All components automatically use updated colors

---

## Testing Strategy

### Unit Testing Locations
- `internal/notes/markdown_test.go` - Test markdown functions
- `internal/notes/files_test.go` - Test file operations
- `internal/config/config_test.go` - Test configuration

### Integration Testing
- Test full flow in `internal/app/`
- Mock file system for testing

---

## Dependencies

### External Packages
- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/bubbles` - TUI components
- `github.com/charmbracelet/lipgloss` - Styling

### Internal Dependencies
```
main.go
  └── internal/app
        ├── internal/config
        ├── internal/notes
        └── internal/ui/styles
```

---

## Design Principles

1. **Separation of Concerns**: Each package has a single, clear responsibility
2. **Encapsulation**: Internal packages prevent external dependencies
3. **Modularity**: Features can be modified independently
4. **Testability**: Pure functions and clear interfaces
5. **Scalability**: Easy to add new features without breaking existing code

---

## Future Enhancements

Suggested areas for expansion:
- `internal/search/` - Note search functionality
- `internal/export/` - Export notes to PDF, HTML
- `internal/sync/` - Cloud synchronization
- `internal/tags/` - Note tagging and categorization
- `internal/themes/` - Multiple color themes
- `internal/plugins/` - Plugin system

---

## Build and Run

```bash
# Build
make build

# Run
make run

# Clean
make clean
```

---

## Troubleshooting

**Can't find package errors**:
- Ensure `go.mod` module path matches import paths
- Run `go mod tidy`

**Import cycle errors**:
- Check package dependencies
- Avoid circular imports between packages

**Undefined references**:
- Ensure exported names start with capital letters
- Check import statements