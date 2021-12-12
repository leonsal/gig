package gig

import (
	"strconv"
	"sync"

	"github.com/leonsal/gig/imgui"
)

type TableModelRow struct {
	Id       interface{}
	Row      []*Widget
	BgTarget imgui.TableBgTarget
	Color    imgui.U32
}

func (tmr *TableModelRow) SetColor(target imgui.TableBgTarget, color imgui.Vec4) {

	tmr.BgTarget = target
	tmr.Color = imgui.ColorConvertFloat4ToU32(color)
}

type TableModel struct {
	mut  sync.Mutex
	rows []TableModelRow
}

func NewTableModel() *TableModel {

	m := new(TableModel)
	return m
}

func (m *TableModel) Rows() []TableModelRow {

	return m.rows
}

func (m *TableModel) RowCount() int {

	m.mut.Lock()
	defer m.mut.Unlock()
	return len(m.rows)
}

func (m *TableModel) AddRow(id interface{}, row []*Widget) int {

	m.mut.Lock()
	defer m.mut.Unlock()
	m.rows = append(m.rows, TableModelRow{id, row, imgui.TableBgTarget_None, imgui.U32(0)})
	return len(m.rows) - 1
}

func (m *TableModel) InsRow(idx int, id interface{}, row []*Widget) bool {

	m.mut.Lock()
	defer m.mut.Unlock()
	if idx < 0 || idx > len(m.rows) {
		return false
	}
	m.rows = append(m.rows, TableModelRow{})
	copy(m.rows[idx+1:], m.rows[idx:])
	m.rows[idx] = TableModelRow{id, row, imgui.TableBgTarget_None, imgui.U32(0)}
	return true
}

func (m *TableModel) DelRow(idx int) bool {

	m.mut.Lock()
	defer m.mut.Unlock()
	if idx < 0 || idx >= len(m.rows) {
		return false
	}
	m.delRow(idx)
	return true
}

// Clear clears the model data without deallocating the previous allocated rows
func (m *TableModel) Clear() {

	m.mut.Lock()
	defer m.mut.Unlock()
	m.rows = m.rows[:0]
}

func (m *TableModel) SetRowColor(idx int, target imgui.TableBgTarget, color imgui.Vec4) bool {

	m.mut.Lock()
	defer m.mut.Unlock()
	if idx < 0 || idx >= len(m.rows) {
		return false
	}
	m.rows[idx].SetColor(target, color)
	return true
}

func (m *TableModel) delRow(idx int) {

	copy(m.rows[idx:], m.rows[idx+1:])
	m.rows[idx-1] = TableModelRow{}
	m.rows = m.rows[:len(m.rows)-1]
}

func (m *TableModel) ReadRow(idx int, rf func(*TableModelRow)) {

	m.mut.Lock()
	defer m.mut.Unlock()
	if idx < 0 || idx >= len(m.rows) {
		//rf(nil)
		return
	}
	rf(&m.rows[idx])
}

func (m *TableModel) UpdateRow(idx int, uf func(*TableModelRow)) {

	m.mut.Lock()
	defer m.mut.Unlock()
	if idx < 0 || idx >= len(m.rows) {
		uf(nil)
		return
	}
	uf(&m.rows[idx])
}

func (m *TableModel) UpdateRowById(id interface{}, uf func(*TableModelRow)) {

	m.mut.Lock()
	defer m.mut.Unlock()
	idx := m.findRow(id)
	if idx < 0 {
		uf(nil)
		return
	}
	uf(&m.rows[idx])
}

func (m *TableModel) findRow(id interface{}) int {

	for idx, rm := range m.rows {
		if rm.Id == id {
			return idx
		}
	}
	return -1
}

func (m *TableModel) DelRowByID(id interface{}) bool {

	m.mut.Lock()
	defer m.mut.Unlock()
	idx := m.findRow(id)
	if idx < 0 {
		return false
	}
	m.delRow(idx)
	return true
}

type TableColumn struct {
	label      imgui.CString
	flags      imgui.TableColumnFlags
	init_width float32
	user_id    imgui.ID
}

func NewTableColumn(label string) *TableColumn {

	c := new(TableColumn)
	c.label.SetString(label)
	return c
}

func (c *TableColumn) SetLabel(label string) *TableColumn {

	c.label.SetString(label)
	return c
}

func (c *TableColumn) SetFlags(flags imgui.TableColumnFlags) *TableColumn {

	c.flags = flags
	return c
}

func (c *TableColumn) OrFlags(flags imgui.TableColumnFlags) *TableColumn {

	c.flags |= imgui.TableColumnFlags(flags)
	return c
}

func (c *TableColumn) SetInitWidth(width float32) *TableColumn {

	c.init_width = width
	return c
}

type Table struct {
	Widget
	str_id       imgui.CString
	columns      []*TableColumn
	flags        imgui.TableFlags
	outer_size   imgui.Vec2
	inner_width  float32
	rowHeight    float32
	showHeader   bool
	freezeCols   int
	freezeRows   int
	sortSpecs    imgui.TableSortSpecs
	clipper      *imgui.ListClipper
	model        *TableModel
	onsort       func(*Table, []imgui.TableColumnSortSpecs)
	autoScroll   bool
	scrollBottom bool
}

func NewTable(cols ...*TableColumn) *Table {

	t := new(Table)
	t.Init(t, cols...)
	return t
}

// TODO finalizer -> destroy clipper

// Init initializes a table for widgets which embed a Table
func (t *Table) Init(wi IWidget, cols ...*TableColumn) {

	t.Widget.Init(wi)
	t.clipper = imgui.NewListClipper()
	t.str_id.SetString(strconv.Itoa(t.ID()))
	t.columns = make([]*TableColumn, 0)
	t.showHeader = true
	t.AddColumns(cols...)
	t.SetRender(t.Render)
}

func (t *Table) AddColumns(cols ...*TableColumn) *Table {

	for i := 0; i < len(cols); i++ {
		t.columns = append(t.columns, cols[i])
	}
	return t
}

func (t *Table) Column(idx int) *TableColumn {

	if idx < 0 || idx >= len(t.columns) {
		return nil
	}
	return t.columns[idx]
}

func (t *Table) Flags() imgui.TableFlags {

	return t.flags
}

func (t *Table) SetFlags(flags imgui.TableFlags) *Table {

	t.flags = flags
	return t
}

func (t *Table) OrFlags(flags imgui.TableFlags) *Table {

	t.flags |= flags
	return t
}

func (t *Table) SetOuterSize(size imgui.Vec2) *Table {

	t.outer_size = size
	return t
}

func (t *Table) SetInnerWidth(width float32) *Table {

	t.inner_width = width
	return t
}

func (t *Table) SetRowHeight(height float32) *Table {

	t.rowHeight = height
	return t
}

func (t *Table) SetShowHeader(show bool) *Table {

	t.showHeader = show
	return t
}

func (t *Table) SetFreezeCols(cols int) *Table {

	t.freezeCols = cols
	return t
}

func (t *Table) SetFreezeRows(rows int) *Table {

	t.freezeRows = rows
	return t
}

func (t *Table) SetAutoScroll(autoScroll bool) *Table {

	t.autoScroll = autoScroll
	return t
}

func (t *Table) ScrollToBottom() *Table {

	t.scrollBottom = true
	return t
}

func (t *Table) SetModel(m *TableModel) *Table {

	t.model = m
	return t
}

func (t *Table) SetOnsort(f func(*Table, []imgui.TableColumnSortSpecs)) *Table {

	t.onsort = f
	return t
}

func (t *Table) Render() {

	if imgui.BeginTableCS(t.str_id, len(t.columns), t.flags, t.outer_size, t.inner_width) {

		// Setup columns
		imgui.TableSetupScrollFreeze(t.freezeCols, t.freezeRows)
		for i := 0; i < len(t.columns); i++ {
			col := t.columns[i]
			imgui.TableSetupColumnCS(col.label, col.flags, col.init_width, col.user_id)
		}
		if t.showHeader {
			imgui.TableHeadersRow()
		}

		// Checks sorting
		imgui.TableGetSortSpecs(&t.sortSpecs)
		if t.sortSpecs.SpecsDirty && t.onsort != nil {
			t.onsort(t, t.sortSpecs.Specs)
		}

		// Draw rows
		if t.model != nil {
			t.clipper.Begin(t.model.RowCount(), -1.0)
			for t.clipper.Step() {
				for row := t.clipper.DisplayStart; row < t.clipper.DisplayEnd; row++ {
					imgui.TableNextRow(imgui.TableRowFlags_None, t.rowHeight)
					t.model.ReadRow(int(row), func(rm *TableModelRow) {
						if rm.BgTarget != imgui.TableBgTarget_None {
							imgui.TableSetBgColor(rm.BgTarget, rm.Color, -1)
						}
						for col := 0; col < len(t.columns); col++ {
							imgui.TableNextColumn()
							if col < len(rm.Row) {
								w := rm.Row[col]
								if w != nil {
									w.Render()
								}
							}
						}
					})
				}
			}
			if t.scrollBottom || (t.autoScroll && (imgui.GetScrollY() >= imgui.GetScrollMaxY())) {
				imgui.SetScrollHereY(1.0)
			}
			t.scrollBottom = false
		}
		imgui.EndTable()
	}
}
