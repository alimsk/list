package main

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	unreadStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EA4444")).Render

	normalTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#1A1A1A", Dark: "#DDDDDD"}).
			Render
	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#1A1A1A", Dark: "#DDDDDD"}).
			Padding(0, 0, 0, 2).
			Render
	selectedStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, false, true).
			BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
			Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
			Padding(0, 0, 0, 1).
			Render
)

type Status struct{ Text, Icon string }

var (
	online  = Status{"Online", lipgloss.NewStyle().Foreground(lipgloss.Color("#43B581")).Render("●")}
	offline = Status{"Offline", lipgloss.NewStyle().Foreground(lipgloss.Color("#3D426B")).Render("◯")}
	idle    = Status{"Idle", lipgloss.NewStyle().Foreground(lipgloss.Color("#F9A61A")).Render("◗")}
	dnd     = Status{"Do Not Disturb", lipgloss.NewStyle().Foreground(lipgloss.Color("#F04847")).Render("⊝")}
)

type CustomItem struct {
	ProfilePic rune
	Name       string
	Status     Status
	Desc       string
	Unread     int

	OnSelect tea.Cmd
}

type CustomAdapter []CustomItem

func (c CustomAdapter) Len() int {
	return len(c)
}

// separator
func (c CustomAdapter) Sep() string {
	return "\n\n"
}

func (c CustomAdapter) View(pos, focus int) string {
	const format = "" +
		"╭────%s\n" +
		"│ %c │ %s\n" +
		"╰────%s %s"

	item := c[pos]

	var style func(string) string
	if focus == pos {
		style = selectedStyle
	} else {
		style = normalStyle
	}

	var topRight, btmRight = "╮", "╯"
	if item.Unread > 0 {
		topRight = unreadStyle(strconv.Itoa(item.Unread))
	}
	if len(item.Status.Icon) > 0 {
		btmRight = item.Status.Icon
	}

	return fmt.Sprintf(style(format), topRight, item.ProfilePic,
		item.Name, btmRight, normalTextStyle(item.Status.Text))
}
