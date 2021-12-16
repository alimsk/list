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
		}
	case SelectMsg:
		adapter := m.list.Adapter.(*list.SimpleAdapter)
		m.exitMessage = fmt.Sprintln("You selected", adapter.ItemAt(int(msg)).Title)
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

// generate random items
func RandomItems(n int) *list.SimpleAdapter {
	a := list.NewSimpleAdapter(make(list.SimpleItemList, n))
	a.OnSelect = onSelect
	for i := 0; i < n; i++ {
		a.SetItemAt(i, random[rand.Intn(len(random))])
	}
	return a
}

type SelectMsg int

func onSelect(i int) tea.Cmd {
	return func() tea.Msg {
		return SelectMsg(i)
	}
}

func main() {
	l := list.New(RandomItems(26))
	// enable focus, so you can interact with it
	l.Focus()
	if err := tea.NewProgram(model{list: l}, tea.WithMouseCellMotion()).Start(); err != nil {
		panic(err)
	}
}
