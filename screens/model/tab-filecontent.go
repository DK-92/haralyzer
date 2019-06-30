package model

import (
	"sort"
	"strconv"

	"github.com/DK-92/harviewer/screens/model/windows"

	"github.com/DK-92/harviewer/structs"
	"github.com/lxn/walk"

	_ "github.com/andlabs/ui/winmanifest"
)

var (
	harFile      *structs.MainLog
	filePosition int
)

type EntryTable struct {
	entry   *structs.Entry
	checked bool
}

type EntryModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*EntryTable
}

func NewEntryModel() *EntryModel {
	m := new(EntryModel)
	m.FillRows()
	return m
}

func (m *EntryModel) RowCount() int {
	return len(m.items)
}

func (m *EntryModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.entry.StartedDateTime[11:23]
	case 1:
		return strconv.Itoa(item.entry.Time) + " ms"
	case 2:
		return item.entry.Request.Method
	case 3:
		return item.entry.Request.URL
	case 4:
		return item.entry.Comment
	default:
		return ""
	}

	panic("unexpected col")
}

func (m *EntryModel) Sort(col int, order walk.SortOrder) error {
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
			return c(a.entry.StartedDateTime < b.entry.StartedDateTime)
		case 1:
			return c(a.entry.Time < b.entry.Time)
		case 2:
			return c(a.entry.Request.Method < b.entry.Request.Method)
		case 3:
			return c(a.entry.Request.URL < b.entry.Request.URL)
		case 4:
			return c(a.entry.Comment < b.entry.Comment)
		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}

func (m *EntryModel) FillRows() {
	m.items = make([]*EntryTable, len(harFile.Log.Entries))

	for i := range m.items {
		m.items[i] = &EntryTable{}
		m.items[i].entry = &harFile.Log.Entries[i]
	}

	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}

func MakeFileContentTab(page *walk.TabPage, position int, filename string) {
	filePosition = position
	harFile = &structs.GetData().HarFiles[filePosition]

	page.SetLayout(walk.NewVBoxLayout())

	label, _ := walk.NewLabel(page)
	label.SetText("Double click on a row to view its contents.")

	walk.NewVSeparator(page)

	/**
	 *	Build the table
	 */
	table, _ := walk.NewTableView(page)

	columnTime := walk.NewTableViewColumn()
	columnTime.SetTitle("Time")
	columnTime.SetWidth(100)
	table.Columns().Add(columnTime)

	columnDuration := walk.NewTableViewColumn()
	columnDuration.SetTitle("Duration")
	columnDuration.SetWidth(70)
	table.Columns().Add(columnDuration)

	columnMethod := walk.NewTableViewColumn()
	columnMethod.SetTitle("Method")
	columnMethod.SetWidth(50)
	table.Columns().Add(columnMethod)

	columnURL := walk.NewTableViewColumn()
	columnURL.SetTitle("URL")
	columnURL.SetWidth(400)
	table.Columns().Add(columnURL)

	columnComment := walk.NewTableViewColumn()
	columnComment.SetTitle("Comment")
	columnComment.SetWidth(100)
	table.Columns().Add(columnComment)

	/**
	 *	Fill the table
	 */
	table.SetModel(NewEntryModel())
	table.MouseUp().Attach(func(x, y int, button walk.MouseButton) {
		if len(table.SelectedIndexes()) > 0 {
			entry := harFile.Log.Entries[table.SelectedIndexes()[0]]

			windows.MakeEntityWindow(filename, &entry)
		}
	})
}
