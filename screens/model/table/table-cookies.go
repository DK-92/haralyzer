package table

import (
	"sort"

	"github.com/DK-92/harviewer/structs"
	"github.com/lxn/walk"
)

type CookieTable struct {
	cookie *structs.Cookie
}

type CookieModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*CookieTable
}

func NewCookieModel(cookies *[]structs.Cookie) *CookieModel {
	m := new(CookieModel)
	m.FillRows(cookies)
	return m
}

func (m *CookieModel) RowCount() int {
	return len(m.items)
}

func (m *CookieModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.cookie.Name
	case 1:
		return item.cookie.Value
	case 2:
		return item.cookie.Path
	case 3:
		return item.cookie.Domain
	case 4:
		return item.cookie.Expires
	case 5:
		return item.cookie.HTTPOnly
	case 6:
		return item.cookie.Secure
	case 7:
		return item.cookie.Comment
	default:
		return ""
	}

	panic("unexpected col")
}

func (m *CookieModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.SliceStable(m.items, func(i, j int) bool {
		a, b := m.items[i], m.items[j]

		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}

			return !ls
		}

		switch m.sortColumn {
		case 0:
			return c(a.cookie.Name < b.cookie.Name)
		case 1:
			return c(a.cookie.Value < b.cookie.Value)
		case 2:
			return c(a.cookie.Path < b.cookie.Path)
		case 3:
			return c(a.cookie.Domain < b.cookie.Domain)
		case 4:
			return c(a.cookie.Expires < b.cookie.Expires)
		case 5:
			return true
		case 6:
			return true
		case 7:
			return c(a.cookie.Comment < b.cookie.Comment)
		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}

func (m *CookieModel) FillRows(cookies *[]structs.Cookie) {
	m.items = make([]*CookieTable, len((*cookies)))

	for i := range m.items {
		m.items[i] = &CookieTable{}
		m.items[i].cookie = &(*cookies)[i]
	}

	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}
