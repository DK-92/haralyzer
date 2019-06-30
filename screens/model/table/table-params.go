package table

import (
	"sort"

	"github.com/DK-92/harviewer/structs"
	"github.com/lxn/walk"
)

type ParamTable struct {
	header *structs.Param
}

type ParamModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*ParamTable
}

func NewParamModel(params *[]structs.Param) *ParamModel {
	m := new(ParamModel)
	m.FillRows(params)
	return m
}

func (m *ParamModel) RowCount() int {
	return len(m.items)
}

func (m *ParamModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.header.Name
	case 1:
		return item.header.Value
	case 2:
		return item.header.Comment
	default:
		return ""
	}

	panic("unexpected col")
}

func (m *ParamModel) Sort(col int, order walk.SortOrder) error {
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
			return c(a.header.Name < b.header.Name)
		case 1:
			return c(a.header.Value < b.header.Value)
		case 2:
			return c(a.header.Comment < b.header.Comment)
		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}

func (m *ParamModel) FillRows(params *[]structs.Param) {
	m.items = make([]*ParamTable, len((*params)))

	for i := range m.items {
		m.items[i] = &ParamTable{}
		m.items[i].header = &(*params)[i]
	}

	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}
