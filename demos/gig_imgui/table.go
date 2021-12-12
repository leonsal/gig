package main

import (
	"fmt"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func NewTableDemo() gig.IWidget {

	createTableRow := func(tm *gig.TableModel, row int) {

		col := 0
		rdata := make([]*gig.Widget, 0)
		sel := gig.NewSelectable(fmt.Sprintf("%d", row)).
			SetFlags(imgui.SelectableFlags_SpanAllColumns | imgui.SelectableFlags_AllowItemOverlap).
			SetOnClick(func(s *gig.Selectable) {
				s.SetSelected(!s.Selected())
			}).
			SetOnFocused(func(w *gig.Widget) {
				s := w.GetControl().(*gig.Selectable)
				fmt.Println("focused:" + s.Label())
			})
		rdata = append(rdata, sel.GetWidget())

		col++
		rdata = append(rdata, gig.NewText("row:%d/col%d", row, col).GetWidget())
		//rdata = append(rdata, g.NewCheckbox(""))

		col++
		rdata = append(rdata, gig.NewText("row:%d/col%d", row, col).GetWidget())

		col++
		rdata = append(rdata, gig.NewText("row:%d/col%d", row, col).GetWidget())

		tm.AddRow(row, rdata)
		if row == 10 || row == 13 {
			tm.SetRowColor(row, imgui.TableBgTarget_RowBg0, imgui.ColorName("Green"))
		}
	}

	sortTable := func(t *gig.Table, ss []imgui.TableColumnSortSpecs) {

		fmt.Printf("sortTable: %+v\n", ss)
	}

	// Creates table model and add rows
	tmodel := gig.NewTableModel()
	for row := 0; row < 100; row++ {
		createTableRow(tmodel, row)
	}

	win := gig.NewWindow("Table").AddChildren(

		gig.NewTable().
			AddColumns(
				gig.NewTableColumn("ID").
					SetFlags(imgui.TableColumnFlags_WidthFixed|imgui.TableColumnFlags_NoHide).
					OrFlags(imgui.TableColumnFlags_DefaultSort).
					SetInitWidth(60),
				gig.NewTableColumn("Check").
					SetFlags(imgui.TableColumnFlags_WidthFixed).
					SetInitWidth(60),
				gig.NewTableColumn("Col3").
					SetFlags(imgui.TableColumnFlags_None).
					SetInitWidth(0),
				gig.NewTableColumn("Col4").
					SetFlags(imgui.TableColumnFlags_None).
					SetInitWidth(0),
			).
			SetShowHeader(true).
			SetRowHeight(0).
			SetFreezeRows(1).
			SetFreezeCols(0).
			SetFlags(imgui.TableFlags_Resizable | imgui.TableFlags_Reorderable | imgui.TableFlags_Hideable).
			OrFlags(imgui.TableFlags_BordersV).
			OrFlags(imgui.TableFlags_RowBg).
			OrFlags(imgui.TableFlags_ScrollY).
			OrFlags(imgui.TableFlags_Sortable).
			OrFlags(imgui.TableFlags_SortMulti).
			SetOnsort(sortTable).
			SetModel(tmodel),
	).SetVisible(false).SaveRef("tableDemo")
	return win
}
