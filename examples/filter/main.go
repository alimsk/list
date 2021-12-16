package main

import (
	"fmt"
	"math/rand"
	"unicode"

	"github.com/alimsk/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list          list.Model
	exitMessage   string
	filterPattern string
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	if len(m.exitMessage) > 0 {
		return m.exitMessage
	}
	return "Search: " + m.filterPattern + "\n\n" +
		m.list.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "backspace":
			if len(m.filterPattern) > 0 {
				m.filterPattern = m.filterPattern[:len(m.filterPattern)-1]
				m.list.Adapter.(*list.SimpleAdapter).Filter(m.filterPattern)
			}
		default:
			if len(msg.Runes) > 0 {
				if r := msg.Runes[0]; unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' {
					m.filterPattern += string(r)
					m.list.Adapter.(*list.SimpleAdapter).Filter(m.filterPattern)
				}
			}
		}
	case SelectMsg:
		m.exitMessage = fmt.Sprintln("You selected", m.list.Adapter.(*list.SimpleAdapter).ItemAt(int(msg)).Title)
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

var random = list.SimpleItemList{
	// Title, Desc, Selectable
	{"Unselectable", "Pressing enter will do nothing", false},
	{"Pocky", "Expesive", true},
	{"Ginger", "Exquisite", true},
	{"Plantains", "Questionable", true},
	{"Honey Dew", "Delectable", true},
	{"Pineapple", "Kind of spicy", true},
	{"Snow Peas", "Bold flavour", true},
	{"Party Gherkin", "My favorite", true},
	{"Bananas", "Looks fresh", true},
}

type SelectMsg int

func onSelect(pos int) tea.Cmd {
	return func() tea.Msg {
		return SelectMsg(pos)
	}
}

// generate random items
func RandomItems(n int) *list.SimpleAdapter {
	a := list.NewSimpleAdapter(make(list.SimpleItemList, n))
	a.OnSelect = onSelect
	for i := 0; i < n; i++ {
		a.SetItemAt(i, random[rand.Intn(len(random))])
	}
	return a
}

func main() {
	l := list.New(RandomItems(26))
	// l.ViewMode = true
	// enable focus, so you can interact with it
	l.Focus()
	if err := tea.NewProgram(model{list: l}, tea.WithMouseCellMotion()).Start(); err != nil {
		panic(err)
	}
}
