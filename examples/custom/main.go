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

type SelectMsg int

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	case SelectMsg:
		if items[msg].Name == "Rick" {
			return m, func() tea.Msg {
				openBrowser("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
				return nil
			}
		}
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
	l, _ := list.New(items)
	l.Focus()
	l.OnSelect = func(i int) tea.Cmd {
		return func() tea.Msg {
			return SelectMsg(i)
		}
	}
	if err := tea.NewProgram(model{list: l}).Start(); err != nil {
		panic(err)
	}
}
