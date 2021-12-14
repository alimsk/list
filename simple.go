package list

import (
	"github.com/charmbracelet/lipgloss"
)

// TODO: implement filter

var (
	normalTitle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
			Padding(0, 0, 0, 2)
	normalDesc = normalTitle.Copy().
			Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})

	selectedTitle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, false, true).
			BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
			Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
			Padding(0, 0, 0, 1)
	selectedDesc = selectedTitle.Copy().
			Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})

	dimmedTitle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
			Padding(0, 0, 0, 2)
	dimmedDesc = dimmedTitle.Copy().
			Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"})

	filterMatch = lipgloss.NewStyle().Underline(true)
)

type SimpleItem struct {
	Title string
	Desc  string
}

type SimpleAdapter struct {
	Items []SimpleItem
}

var _ Adapter = (*SimpleAdapter)(nil)

func (s *SimpleAdapter) Count() int {
	return len(s.Items)
}

func (s *SimpleAdapter) Sep() string {
	return "\n\n"
}

func (s *SimpleAdapter) View(pos, focus int) string {
	item := s.Items[pos]

	var renderTitle, renderDesc func(string) string
	if focus == pos {
		renderTitle, renderDesc = selectedTitle.Render, selectedDesc.Render
	} else if focus >= 0 {
		renderTitle, renderDesc = normalTitle.Render, normalDesc.Render
	} else {
		renderTitle, renderDesc = dimmedTitle.Render, dimmedDesc.Render
	}

	return renderTitle(item.Title) + "\n" + renderDesc(item.Desc)
}

func (s *SimpleAdapter) Append(item ...SimpleItem) {
	s.Items = append(s.Items, item...)
}

func (s *SimpleAdapter) Insert(i int, item ...SimpleItem) {
	s.Items = append(s.Items[:i+len(item)], s.Items[i:]...)
	for i2 := i; i2 < len(item); i2++ {
		s.Items[i+i2] = item[i2]
	}
}

func (s *SimpleAdapter) Remove(i int) {
	s.Items = append(s.Items[:i], s.Items[i+1:]...)
}
