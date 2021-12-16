package list

import (
	. "github.com/charmbracelet/lipgloss"
)

var border = NewStyle().PaddingLeft(2)

type SimpleAdapterStyle struct {
	// normal border is just default style with left padding 2
	title, desc /* no border */, filterMatch                         Style
	titleSelected, descSelected, borderSelected, filterMatchSelected Style
}

func (s *SimpleAdapterStyle) Normal(title, desc Style) {
	s.title = title
	s.desc = desc
	s.filterMatch = title.Copy().Underline(true)
}

func (s *SimpleAdapterStyle) Selected(title, desc Style) {
	s.titleSelected = title
	s.descSelected = desc
	s.borderSelected = NewStyle().
		BorderStyle(NormalBorder()).
		BorderLeft(true).
		BorderForeground(title.GetForeground()).
		PaddingLeft(1)
	s.filterMatchSelected = title.Copy().Underline(true)
}

func SimpleDefaultStyle() (normal, dimmed *SimpleAdapterStyle) {
	normal = &SimpleAdapterStyle{}
	normal.Normal(
		NewStyle().
			Foreground(AdaptiveColor{Light: "#1A1A1A", Dark: "#DDDDDD"}),
		NewStyle().
			Foreground(AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}),
	)
	normal.Selected(
		NewStyle().
			Foreground(AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}),
		NewStyle().
			Foreground(AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}),
	)

	dimmed = &SimpleAdapterStyle{}
	dimmed.Normal(
		NewStyle().
			Foreground(AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}),
		NewStyle().
			Foreground(AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"}),
	)
	dimmed.Selected(
		NewStyle().
			Foreground(AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}),
		NewStyle().
			Foreground(AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"}),
	)

	return
}
