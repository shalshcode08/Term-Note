package styles

import "github.com/charmbracelet/lipgloss"

// ASCII Art Options for TERMNOTE Header
var (
	asciiArtBox = `
 ████████╗███████╗██████╗ ███╗   ███╗███╗   ██╗ ██████╗ ████████╗███████╗
 ╚══██╔══╝██╔════╝██╔══██╗████╗ ████║████╗  ██║██╔═══██╗╚══██╔══╝██╔════╝
    ██║   █████╗  ██████╔╝██╔████╔██║██╔██╗ ██║██║   ██║   ██║   █████╗
    ██║   ██╔══╝  ██╔══██╗██║╚██╔╝██║██║╚██╗██║██║   ██║   ██║   ██╔══╝
    ██║   ███████╗██║  ██║██║ ╚═╝ ██║██║ ╚████║╚██████╔╝   ██║   ███████╗
    ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═══╝ ╚═════╝    ╚═╝   ╚══════╝
`

	asciiArtSimple = `
 ▀▀█▀▀ ██▀▀▀ ██▀▀█ ▄▀▀▄▀▀▄ ▄▀▀▄  █ ▄▀▀▄ ▀▀█▀▀ ██▀▀▀
   █   █▀▀   █▄▄▀  █  ▀  █ █  █▀▄█ █  █   █   █▀▀
   █   ██▄▄▄ ▀  ▀▄ ▀    ▀  ▀  ▀  ▀  ▀▀    ▀   ██▄▄▄
`

	asciiArtSlant = `
  ______                   _   __      __
 /_  __/__ ______ _  ___  / | / /__  / /____
  / / / -_) __/  ' \/ _ \/  |/ / _ \/ __/ -_)
 /_/  \__/_/ /_/_/_/_//_/_/|___/\___/\__/\__/
`

	asciiArtDouble = `
 ████████╗███████╗██████╗ ███╗   ███╗███╗   ██╗ ██████╗ ████████╗███████╗
 ╚══██╔══╝██╔════╝██╔══██╗████╗ ████║████╗  ██║██╔═══██╗╚══██╔══╝██╔════╝
    ██║   █████╗  ██████╔╝██╔████╔██║██╔██╗ ██║██║   ██║   ██║   █████╗
    ██║   ██╔══╝  ██╔══██╗██║╚██╔╝██║██║╚██╗██║██║   ██║   ██║   ██╔══╝
    ██║   ███████╗██║  ██║██║ ╚═╝ ██║██║ ╚████║╚██████╔╝   ██║   ███████╗
    ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═══╝ ╚═════╝    ╚═╝   ╚══════╝
`

	asciiArtMinimal = `
  ████████ ███████ ████  ████ ███  █ █████ ████████ ███████
     ██    ██      ██  █ █  █ ██ █ █ ██  █    ██    ██
     ██    █████   ████  ████ █ ██ █ ██  █    ██    █████
     ██    ██      ██ █  █ █  █  ███ ██  █    ██    ██
     ██    ███████ ██  █ █  █ █   ██ █████    ██    ███████
`
)

// Color Scheme - All exported for external use
var (
	// Primary colors
	ColorPrimary   = lipgloss.Color("205") // Pink/Magenta
	ColorSecondary = lipgloss.Color("141") // Purple
	ColorAccent    = lipgloss.Color("212") // Light pink
	ColorMuted     = lipgloss.Color("241") // Gray
	ColorText      = lipgloss.Color("252") // Light gray
	ColorSuccess   = lipgloss.Color("42")  // Green
	ColorWarning   = lipgloss.Color("214") // Orange
	ColorError     = lipgloss.Color("196") // Red
	ColorBorder    = lipgloss.Color("99")  // Soft purple for borders
	ColorBg        = lipgloss.Color("235") // Dark background
)

// Global Styles - All exported for external use
var (
	HeaderStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true).
			Align(lipgloss.Center)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(ColorMuted).
			Align(lipgloss.Center)

	HelpTitleStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true).
			MarginBottom(1)

	KeyStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true).
			Width(12).
			Align(lipgloss.Left)

	DescStyle = lipgloss.NewStyle().
			Foreground(ColorText).
			Width(30)

	HelpItemStyle = lipgloss.NewStyle().
			PaddingLeft(1)

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPrimary).
			Padding(1, 2).
			MarginTop(1)

	ViewTitleStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true).
			MarginBottom(1)

	ViewHelpStyle = lipgloss.NewStyle().
			Foreground(ColorMuted)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(ColorWarning).
			Bold(true)

	// Input Dialog Styles
	DialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(ColorPrimary).
			Padding(2, 4).
			Width(70).
			Align(lipgloss.Center)

	DialogTitleStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true).
				Align(lipgloss.Center).
				Width(62).
				MarginBottom(2)

	InputLabelStyle = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true).
			MarginBottom(1)

	InputBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Padding(0, 1).
			Width(60).
			MarginBottom(2)

	InputHelpStyle = lipgloss.NewStyle().
			Foreground(ColorMuted).
			Italic(true).
			Align(lipgloss.Center).
			Width(62).
			MarginTop(1)

	InputTipStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			Align(lipgloss.Center).
			Width(62).
			MarginTop(1)

	// File creation specific
	FileIconStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true).
			MarginRight(1)

	FileExtensionStyle = lipgloss.NewStyle().
				Foreground(ColorMuted).
				Italic(true)

	// List styles
	ListTitleStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true).
			Padding(0, 0).
			MarginLeft(0).
			MarginTop(0).
			MarginBottom(0)

	ListItemTitleStyle = lipgloss.NewStyle().
				Foreground(ColorText).
				Bold(true)

	ListItemDescStyle = lipgloss.NewStyle().
				Foreground(ColorMuted)

	ListItemSelectedTitleStyle = lipgloss.NewStyle().
					Foreground(ColorPrimary).
					Bold(true)

	ListItemSelectedDescStyle = lipgloss.NewStyle().
					Foreground(ColorAccent)

	ListHelpStyle = lipgloss.NewStyle().
			Foreground(ColorMuted).
			Padding(1, 2)
)

// GetASCIIArt returns the selected ASCII art style
// Options: "box", "simple", "slant", "double", "minimal"
func GetASCIIArt(style string) string {
	switch style {
	case "simple":
		return asciiArtSimple
	case "slant":
		return asciiArtSlant
	case "double":
		return asciiArtDouble
	case "minimal":
		return asciiArtMinimal
	default:
		return asciiArtBox
	}
}
