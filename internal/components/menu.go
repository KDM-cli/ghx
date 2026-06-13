package components

import (
	"fmt"
	"strings"

	"github.com/utkarshpatrikar/ghx/internal/styles"
)

type MenuItem struct {
	Title       string
	Description string
	Target      string
}

type Menu struct {
	items    []MenuItem
	selected int
}

func NewMenu(items []MenuItem) Menu {
	return Menu{items: items}
}

func (m *Menu) Next() {
	if len(m.items) == 0 {
		return
	}
	m.selected = (m.selected + 1) % len(m.items)
}

func (m *Menu) Prev() {
	if len(m.items) == 0 {
		return
	}
	m.selected--
	if m.selected < 0 {
		m.selected = len(m.items) - 1
	}
}

func (m Menu) Selected() MenuItem {
	if len(m.items) == 0 {
		return MenuItem{}
	}
	return m.items[m.selected]
}

func (m Menu) View() string {
	lines := make([]string, 0, len(m.items))
	for i, item := range m.items {
		cursor := "  "
		title := item.Title
		if i == m.selected {
			cursor = styles.Cursor.Render("> ")
			title = styles.Selected.Render(item.Title)
		}
		lines = append(lines, fmt.Sprintf("%s%-12s %s", cursor, title, styles.Muted.Render(item.Description)))
	}
	return strings.Join(lines, "\n")
}
