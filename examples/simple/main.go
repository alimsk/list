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
	if m.exitMessage != "" {
		return m.exitMessage
	}
	return strconv.Itoa(m.list.Adapter.Count()) + " items\n\n" +
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
		m.exitMessage = fmt.Sprintln("You selected", adapter.Items[msg].Title)
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

var random = []list.SimpleItem{
	{"Pocky", "Expesive"},
	{"Ginger", "Exquisite"},
	{"Plantains", "Questionable"},
	{"Honey Dew", "Delectable"},
	{"Pineapple", "Kind of spicy"},
	{"Snow Peas", "Bold flavour"},
	{"Party Gherkin", "My favorite"},
	{"Bananas", "Looks fresh"},
}

// generate random items
func RandomItems(n int) *list.SimpleAdapter {
	a := &list.SimpleAdapter{}
	a.Items = make([]list.SimpleItem, n)
	for i := range a.Items {
		a.Items[i] = random[rand.Intn(len(random))]
	}
	return a
}

func NumberList(n int) *list.SimpleAdapter {
	a := &list.SimpleAdapter{}
	a.Items = make([]list.SimpleItem, n)
	for i := range a.Items {
		a.Items[i] = list.SimpleItem{strconv.Itoa(i), "an item"}
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
	l, _ := list.New(RandomItems(26))
	l.OnSelect = onSelect
	// enable focus, so you can interact with it
	l.Focus()
	if err := tea.NewProgram(model{list: l}, tea.WithMouseCellMotion()).Start(); err != nil {
		panic(err)
	}
}
