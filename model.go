package list

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	VisibleItemCount int
	InfiniteScroll   bool
	ScrollBarStyle   lipgloss.Style
	Adapter          Adapter

	focus            int
	visibleItemStart int
	hasFocus         bool

	OnSelect func(int) tea.Cmd
}

func New(adapter Adapter) (Model, error) {
	if adapter == nil {
		return Model{}, fmt.Errorf("adapter is nil")
	}
	return Model{
		VisibleItemCount: 7,
		Adapter:          adapter,
		ScrollBarStyle: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
				Light: "#333B4E", Dark: "#FBB3BD",
			}).
			SetString("\u2502"), // '│', U+2502, BOX DRAWINGS LIGHT VERTICAL
	}, nil
}

func (m Model) View() string {
	var bob strings.Builder

	for i := m.visibleItemStart; i < m.Adapter.Count() && i < m.visibleItemStart+m.VisibleItemCount; i++ {
		focus := -1
		if m.hasFocus {
			focus = m.focus
		}
		bob.WriteString(m.Adapter.View(i, focus) + m.Adapter.Sep())
	}

	s := bob.String()
	s = s[:max(0, len(s)-len(m.Adapter.Sep()))] // remove trailing separator

	if m.Adapter.Count() > m.VisibleItemCount {
		/* draw scrollbar */
		bob.Reset()

		height := lipgloss.Height(s)
		scrollbarpos := int((float32(m.visibleItemStart) / float32(m.Adapter.Count()-m.VisibleItemCount)) *
			float32(height-1)) // -1 because it start from 0 not 1

		// first line
		if scrollbarpos == 0 {
			bob.WriteString(m.ScrollBarStyle.String())
		} else {
			bob.WriteByte(' ')
		}
		var lineno int
		for _, r := range s {
			bob.WriteRune(r)
			if r == '\n' {
				lineno++
				if lineno == scrollbarpos {
					bob.WriteString(m.ScrollBarStyle.String())
				} else {
					bob.WriteByte(' ')
				}
			}
		}

		return bob.String()
	}

	return s
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.hasFocus {
		return m, nil
	}

	if m.Adapter.Count() <= m.visibleItemStart+m.VisibleItemCount {
		m.visibleItemStart = max(0, m.Adapter.Count()-m.VisibleItemCount)
	}
	if m.Adapter.Count() <= m.focus {
		m.focus = max(0, m.Adapter.Count()-1)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.updateFocus(-1)
		case "down", "tab", "shift+tab":
			m.updateFocus(+1)
		case "enter":
			if m.OnSelect != nil {
				m.adjustView()
				return m, m.OnSelect(m.focus)
			}
		case "home":
			m.SetItemFocus(0)
		case "end":
			m.SetItemFocus(m.Adapter.Count() - 1)
		case "pgup":
			m.visibleItemStart = max(0, m.visibleItemStart-m.VisibleItemCount)
		case "pgdown":
			m.visibleItemStart = min(max(0, m.Adapter.Count()-m.VisibleItemCount), m.visibleItemStart+m.VisibleItemCount)
		}
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseWheelUp:
			m.visibleItemStart = max(0, m.visibleItemStart-1)
		case tea.MouseWheelDown:
			m.visibleItemStart = min(max(0, m.Adapter.Count()-m.VisibleItemCount), m.visibleItemStart+1)
		}
	}

	return m, nil
}

func (m *Model) updateFocus(i int) {
	if m.InfiniteScroll {
		m.ShiftItemFocus(i)
	} else {
		m.SetItemFocus(m.focus + i)
	}
}

func (m *Model) Focus() {
	m.hasFocus = true
}

func (m *Model) Blur() {
	m.hasFocus = false
}

func (m *Model) SetItemFocus(i int) error {
	m.focus = max(0, min(i, m.Adapter.Count()-1))
	m.adjustView()
	return nil
}

func (m *Model) ItemFocus() int {
	return m.focus
}

func (m *Model) ShiftItemFocus(i int) {
	m.focus = mod(m.focus+i, m.Adapter.Count())
	m.adjustView()
}

func (m *Model) VisibleItemStart() int {
	return m.visibleItemStart
}

func (m *Model) adjustView() {
	if m.focus < m.visibleItemStart {
		m.visibleItemStart = m.focus
	} else if m.focus >= m.visibleItemStart+m.VisibleItemCount {
		m.visibleItemStart = m.focus - (m.VisibleItemCount - 1)
	}
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
