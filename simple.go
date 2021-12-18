package list

import (
	"github.com/sahilm/fuzzy"
)

type SimpleItem struct {
	Title, Desc string
	Disabled    bool
}

type SimpleItemList []SimpleItem

var _ fuzzy.Source = (SimpleItemList)(nil)

func (s SimpleItemList) Len() int {
	return len(s)
}

func (s SimpleItemList) String(i int) string {
	return s[i].Title
}

type SimpleAdapter struct {
	items             SimpleItemList
	filterResult      []fuzzy.Match
	lastFilterPattern string

	StyleNormal, StyleDimmed *SimpleAdapterStyle
}

func NewSimpleAdapter(items SimpleItemList) *SimpleAdapter {
	styleNormal, styleDimmed := SimpleDefaultStyle()
	return &SimpleAdapter{
		items: items,

		StyleNormal: styleNormal,
		StyleDimmed: styleDimmed,
	}
}

var _ Adapter = (*SimpleAdapter)(nil)

func (s *SimpleAdapter) Len() int {
	if s.filterResult == nil {
		return len(s.items)
	}
	return len(s.filterResult)
}

func (s *SimpleAdapter) Sep() string {
	return "\n\n"
}

func (s *SimpleAdapter) View(pos, focus int) string {
	var match *fuzzy.Match
	if s.filterResult != nil {
		match = &s.filterResult[pos]
		pos = match.Index
		if focus >= 0 {
			focus = s.filterResult[focus].Index
		}
	}

	item := s.items[pos]

	var renderTitle, renderDesc, renderBorder, renderFilterMatch func(string) string
	if !item.Disabled {
		if focus == pos {
			// focused
			st := s.StyleNormal
			renderTitle = st.titleSelected.Render
			renderDesc = st.descSelected.Render
			renderBorder = st.borderSelected.Render
			renderFilterMatch = st.filterMatchSelected.Render
		} else if focus == FocusDisabled {
			// disabled
			st := s.StyleDimmed
			renderTitle = st.title.Render
			renderDesc = st.desc.Render
			renderBorder = border.Render
			renderFilterMatch = st.filterMatch.Render
		} else {
			// blurred/viewmode
			st := s.StyleNormal
			renderTitle = st.title.Render
			renderDesc = st.desc.Render
			renderBorder = border.Render
			renderFilterMatch = st.filterMatch.Render
		}
	} else {
		st := s.StyleDimmed
		if focus == pos {
			// focused
			renderTitle = st.titleSelected.Render
			renderDesc = st.descSelected.Render
			renderBorder = st.borderSelected.Render
			renderFilterMatch = st.filterMatchSelected.Render
		} else {
			// blurred/disabled
			renderTitle = st.title.Render
			renderDesc = st.desc.Render
			renderBorder = border.Render
			renderFilterMatch = st.filterMatch.Render
		}
	}

	var title = item.Title
	if match != nil {
		var highlightedTitle string
	Outer:
		for i, r := range title {
			for _, i2 := range match.MatchedIndexes {
				if i == i2 {
					highlightedTitle += renderFilterMatch(string(r))
					continue Outer
				}
			}
			highlightedTitle += renderTitle(string(r))
		}
		title = highlightedTitle
	}

	return renderBorder(renderTitle(title) + "\n" + renderDesc(item.Desc))
}

func (s *SimpleAdapter) Append(item ...SimpleItem) {
	s.items = append(s.items, item...)
	if s.filterResult != nil {
		s.Filter(s.lastFilterPattern)
	}
}

func (s *SimpleAdapter) Insert(i int, item ...SimpleItem) {
	s.items = append(s.items[:i+len(item)], s.items[i:]...)
	for i2 := 0; i2 < len(item); i2++ {
		s.items[i+i2] = item[i2]
	}
	if s.filterResult != nil {
		s.Filter(s.lastFilterPattern)
	}
}

func (s *SimpleAdapter) Remove(i int) {
	s.items = append(s.items[:i], s.items[i+1:]...)
	if s.filterResult != nil {
		s.Filter(s.lastFilterPattern)
	}
}

func (s *SimpleAdapter) Filter(pattern string) {
	s.lastFilterPattern = pattern
	if len(pattern) > 0 {
		s.filterResult = fuzzy.FindFrom(pattern, s.items)
		if s.filterResult == nil {
			s.filterResult = make([]fuzzy.Match, 0)
		}
	} else {
		s.filterResult = nil
	}
}

func (s *SimpleAdapter) OriginalItemLen() int {
	return len(s.items)
}

func (s *SimpleAdapter) FilteredIndex(pos int) int {
	if s.filterResult == nil {
		return pos
	}
	return s.filterResult[pos].Index
}

func (s *SimpleAdapter) FilteredItemAt(pos int) SimpleItem {
	if s.filterResult == nil {
		return s.items[pos]
	}
	return s.items[s.filterResult[pos].Index]
}

func (s *SimpleAdapter) ItemAt(pos int) SimpleItem {
	return s.items[pos]
}

func (s *SimpleAdapter) SetItemAt(pos int, item SimpleItem) {
	s.items[pos] = item
	if s.filterResult != nil {
		s.Filter(s.lastFilterPattern)
	}
}

func (s *SimpleAdapter) SetItems(items SimpleItemList) {
	s.items = items
	if s.filterResult != nil {
		s.Filter(s.lastFilterPattern)
	}
}
