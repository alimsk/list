package list

type Adapter interface {
	// number of items, usually just return len(items)
	Count() int
	// separator
	Sep() string
	// item view, pos is position of the item, focus is current focused item.
	//
	// to check whether this item is focused or not, use focus == pos
	View(pos, focus int) string
}
