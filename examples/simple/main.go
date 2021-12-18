package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/alimsk/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list        list.Model
	exitMessage string
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	if len(m.exitMessage) > 0 {
		return m.exitMessage
	}
	return strconv.Itoa(m.list.Adapter.Len()) + " items\n\n" +
		m.list.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "enter":
			adapter := m.list.Adapter.(*list.SimpleAdapter)
			item := adapter.FilteredItemAt(m.list.ItemFocus())
			if !item.Disabled {
				m.exitMessage = fmt.Sprintln("You selected", item.Title)
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

var random = list.SimpleItemList{
	// Title, Desc, Disabled
	{"Disabled", "Pressing enter will do nothing", true},
	{"Pocky", "Expesive", false},
	{"Ginger", "Exquisite", false},
	{"Plantains", "Questionable", false},
	{"Honey Dew", "Delectable", false},
	{"Pineapple", "Kind of spicy", false},
	{"Snow Peas", "Bold flavour", false},
	{"Party Gherkin", "My favorite", false},
	{"Bananas", "Looks fresh", false},
}

// generate random items
func RandomItems(n int) *list.SimpleAdapter {
	a := list.NewSimpleAdapter(make(list.SimpleItemList, n))
	for i := 0; i < n; i++ {
		a.SetItemAt(i, random[rand.Intn(len(random))])
	}
	return a
}

func main() {
	l := list.New(RandomItems(26))
	// enable focus, so you can interact with it
	l.Focus()
	if err := tea.NewProgram(model{list: l}, tea.WithMouseCellMotion()).Start(); err != nil {
		panic(err)
	}
}
