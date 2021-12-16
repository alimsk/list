package main

import (
	"github.com/alimsk/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list list.Model
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return m.list.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "q" || key.String() == "esc" || key.String() == "ctrl+c" {
		return m, tea.Quit
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

var items = CustomAdapter{
	{
		ProfilePic: '🙂',
		Name:       "David",
		Status:     online,
		Unread:     2,
	},
	{
		ProfilePic: '😎',
		Name:       "Alex",
		Status:     idle,
	},
	{
		ProfilePic: '😐',
		Name:       "Carl",
		Status:     online,
		Unread:     1,
	},
	{
		ProfilePic: '🙁',
		Name:       "Martin",
		Status:     dnd,
		Unread:     4,
	},
	{
		ProfilePic: '😎',
		Name:       "Rick",
		Status:     Status{"never gonna give you up", "🎧"},
		Unread:     12,
		OnSelect: func() tea.Msg {
			openBrowser("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
			return nil
		},
	},
	{
		ProfilePic: '😁',
		Name:       "Johnny",
		Status:     offline,
	},
	{
		ProfilePic: '😀',
		Name:       "Jimmy",
		Status:     offline,
		Unread:     1,
	},
}

func main() {
	l := list.New(items)
	l.Focus()
	if err := tea.NewProgram(model{list: l}).Start(); err != nil {
		panic(err)
	}
}
