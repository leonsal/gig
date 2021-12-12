package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// typedef struct xImGuiTableColumnSortSpecs {
// 		const ImGuiTableColumnSortSpecs* ss;
// 		int Direction;
// } xImGuiTableColumnSortSpecs;
// static xImGuiTableColumnSortSpecs xGetTableColumnSortSpec(const ImGuiTableColumnSortSpecs *specs, int index) {
//		xImGuiTableColumnSortSpecs xss;
//		xss.ss = &specs[index];
//		xss.Direction = specs[index].SortDirection;
//		return xss;
//	}
import "C"
import (
	"unsafe"
)

func BeginTable(strId string, column int, flags TableFlags, outerSize Vec2, innerWidth float32) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	return bool(C.igBeginTable(str_id, C.int(column), C.ImGuiTableFlags(flags), outerSize.ImVec2(), C.float(innerWidth)))
}

func BeginTableCS(strId CString, column int, flags TableFlags, outerSize Vec2, innerWidth float32) bool {

	return bool(C.igBeginTable(strId.CString(), C.int(column), C.ImGuiTableFlags(flags), outerSize.ImVec2(), C.float(innerWidth)))
}

func EndTable() {

	C.igEndTable()
}

func TableNextRow(rowFlags TableRowFlags, minRowHeight float32) {

	C.igTableNextRow(C.ImGuiTableRowFlags(rowFlags), C.float(minRowHeight))
}

func TableNextColumn() {

	C.igTableNextColumn()
}

func TableSetColumnIndex(columnN int) {

	C.igTableSetColumnIndex(C.int(columnN))
}

func TableSetupColumn(label string, flags TableColumnFlags, initWidthOrHeight float32, userId ID) {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.igTableSetupColumn(clabel, C.ImGuiTableColumnFlags(flags), C.float(initWidthOrHeight), C.ImGuiID(userId))
}

func TableSetupColumnCS(label CString, flags TableColumnFlags, initWidthOrHeight float32, userId ID) {

	C.igTableSetupColumn(label.CString(), C.ImGuiTableColumnFlags(flags), C.float(initWidthOrHeight), C.ImGuiID(userId))
}

func TableSetupScrollFreeze(cols int, rows int) {

	C.igTableSetupScrollFreeze(C.int(cols), C.int(rows))
}

func TableHeadersRow() {

	C.igTableHeadersRow()
}

func TableHeader(label string) {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.igTableHeader(clabel)
}

func TableHeaderCS(label CString) {

	C.igTableHeader(label.CString())
}

func TableGetSortSpecs(ss *TableSortSpecs) {

	cspecs := C.igTableGetSortSpecs()
	if cspecs == nil {
		ss.Specs = nil
		ss.SpecsDirty = false
		return
	}
	if len(ss.Specs) < int(cspecs.SpecsCount) {
		ss.Specs = make([]TableColumnSortSpecs, int(cspecs.SpecsCount))
	} else {
		ss.Specs = ss.Specs[:int(cspecs.SpecsCount)]
	}
	ss.SpecsDirty = bool(cspecs.SpecsDirty)
	cspecs.SpecsDirty = false

	for i := 0; i < len(ss.Specs); i++ {
		colSpec := C.xGetTableColumnSortSpec(cspecs.Specs, C.int(i))
		ss.Specs[i] = TableColumnSortSpecs{
			ColumnUserID: ID(colSpec.ss.ColumnUserID),
			ColumnIndex:  int(colSpec.ss.ColumnIndex),
			SortOrder:    int(colSpec.ss.SortOrder),
			Direction:    SortDirection(colSpec.Direction),
		}
	}
}

func TableGetColumnCount() int {

	return int(C.igTableGetColumnCount())
}

func TableGetColumnIndex() int {

	return int(C.igTableGetColumnIndex())
}

func TableGetRowIndex() int {

	return int(C.igTableGetRowIndex())
}

func TableGetColumnName(columnN int) string {

	return C.GoString(C.igTableGetColumnName_Int(C.int(columnN)))
}

func TableGetColumnFlags(columnN int) TableColumnFlags {

	return TableColumnFlags(C.igTableGetColumnFlags(C.int(columnN)))
}

func TableSetColumnEnables(columnN int, v bool) {

	C.igTableSetColumnEnabled(C.int(columnN), C.bool(v))
}

func TableSetBgColor(target TableBgTarget, color U32, columnN int) {

	C.igTableSetBgColor(C.ImGuiTableBgTarget(target), C.ImU32(color), C.int(columnN))
}

func Columns(count int, id string, border bool) {

	cid := C.CString(id)
	defer C.free(unsafe.Pointer(cid))
	C.igColumns(C.int(count), cid, C.bool(border))
}

func NextColumn() {

	C.igNextColumn()
}

func GetColumnIndex() int {

	return int(C.igGetColumnIndex())
}

func GetColumnWidth(columnIndex int) float32 {

	return float32(C.igGetColumnWidth(C.int(columnIndex)))
}

func SetColumnWidth(columnIndex int, width float32) {

	C.igSetColumnWidth(C.int(columnIndex), C.float(width))
}

func GetColumnOffset(columnIndex int) float32 {

	return float32(C.igGetColumnOffset(C.int(columnIndex)))
}

func SetColumnOffset(columnIndex int, offsetX int) {

	C.igSetColumnOffset(C.int(columnIndex), C.float(offsetX))
}

func GetColumnsCount() int {

	return int(C.igGetColumnsCount())
}
