package list

import (
	tea "github.com/charmbracelet/bubbletea"
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

	selectedDimmedTitle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"}).
				Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
				Padding(0, 0, 0, 1)
	selectedDimmedDesc = selectedDimmedTitle.Copy().
				Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"})

	filterMatch = lipgloss.NewStyle().Underline(true)
)

type SimpleItem struct {
	Title, Desc string
	Selectable  bool
}

type SimpleAdapter struct {
	Items    []SimpleItem
	OnSelect func(pos int) tea.Cmd
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
	if item.Selectable {
		if focus == pos {
			// focused
			renderTitle, renderDesc = selectedTitle.Render, selectedDesc.Render
		} else if focus >= 0 {
			// blurred
			renderTitle, renderDesc = normalTitle.Render, normalDesc.Render
		} else {
			// disabled
			renderTitle, renderDesc = dimmedTitle.Render, dimmedDesc.Render
		}
	} else {
		if focus == pos {
			// focused
			renderTitle, renderDesc = selectedDimmedTitle.Render, selectedDimmedDesc.Render
		} else {
			// blurred, disabled
			renderTitle, renderDesc = dimmedTitle.Render, dimmedDesc.Render
		}
	}

	return renderTitle(item.Title) + "\n" + renderDesc(item.Desc)
}

func (s *SimpleAdapter) Select(pos int) tea.Cmd {
	if onselect := s.OnSelect; s.Items[pos].Selectable && onselect != nil {
		return onselect(pos)
	}
	return nil
}

func (s *SimpleAdapter) Append(item ...SimpleItem) {
	s.Items = append(s.Items, item...)
}

func (s *SimpleAdapter) Insert(i int, item ...SimpleItem) {
	s.Items = append(s.Items[:i+len(item)], s.Items[i:]...)
	for i2 := 0; i2 < len(item); i2++ {
		s.Items[i+i2] = item[i2]
	}
}

func (s *SimpleAdapter) Remove(i int) {
	s.Items = append(s.Items[:i], s.Items[i+1:]...)
}
