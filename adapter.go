package list

type Adapter interface {
	// return the number of items, usually just return len(items)
	Count() int
	// separator
	Sep() string
	// item view, pos is position of the item, focus is currently focused item.
	//
	// tip: to check whether this item is focused or not, use focus == pos
	View(pos, focus int) string
}
