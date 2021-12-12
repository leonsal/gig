package dialog

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
	"golang.org/x/text/message"
)

type FileDialog struct {
	gig.PopupModal
	path         string
	nextPath     string
	comboPath    *gig.Combo
	grpNewFolder *gig.Group
	btnNewFolder *gig.Button
	filterInput  *gig.InputText
	checkHidden  *gig.Checkbox
	table        *gig.Table
	model        *gig.TableModel
	fileInput    *gig.InputText
	btnOK        *gig.Button
	btnCancel    *gig.Button
	footerStart  imgui.Vec2
	footerEnd    imgui.Vec2
	printer      *message.Printer
	dlgNewFolder *InputDialog
	dlgMsg       *MsgDialog
	timeFormat   string
	sortSpecs    []imgui.TableColumnSortSpecs
	onSelect     func(fd *FileDialog)
	onCancel     func(fd *FileDialog)
}

// NewFileDialog creates and returns a pointer to a new FileDialog
// showing the contents of the specified directory.
// It return error if the specified directory is invalid.
func NewFileDialog(title, dir, filters string) (*FileDialog, error) {

	d := new(FileDialog)
	d.Init(d, title)

	d.table = gig.NewTable(
		gig.NewTableColumn("Filename").
			SetFlags(imgui.TableColumnFlags_NoHide).
			SetInitWidth(4),
		gig.NewTableColumn("Size").
			SetInitWidth(1),
		gig.NewTableColumn("Modification Time").
			SetFlags(imgui.TableColumnFlags_DefaultHide).
			SetInitWidth(2),
	)
	d.table.SetFlags(imgui.TableFlags_Sortable | imgui.TableFlags_Resizable | imgui.TableFlags_ScrollY |
		imgui.TableFlags_Borders | imgui.TableFlags_Hideable | imgui.TableFlags_SizingStretchSame)
	d.table.SetFreezeRows(1)

	// PreRender adjusts the size of the table to current available space
	d.table.SetPreRender(func(wi gig.IWidget) {
		avail := imgui.GetContentRegionAvail()
		lineHeight := imgui.GetTextLineHeightWithSpacing()
		avail.Y -= (d.footerEnd.Y - d.footerStart.Y)
		nlines := int(avail.Y / lineHeight)
		height := float32(nlines) * lineHeight
		d.table.SetOuterSize(imgui.Vec2{-1, height})
	})

	d.table.SetPostRender(func(wi gig.IWidget) {
		if len(d.nextPath) > 0 {
			d.SetPath(d.nextPath)
			d.nextPath = ""
		}
		if len(d.sortSpecs) > 0 {
			d.sortTable()
		}
	})
	d.model = gig.NewTableModel()
	d.table.SetModel(d.model)
	d.table.SetOnsort(func(t *gig.Table, s []imgui.TableColumnSortSpecs) {
		d.sortSpecs = s
	})

	d.setHeader()
	d.AddChildren(d.table)
	d.setFooter()
	d.setDlgNewFolder()

	d.printer = message.NewPrinter(message.MatchLanguage("en"))
	d.timeFormat = "2006-Jan-02 15:04:05 -0700"

	d.SetFilters(filters)
	err := d.SetPath(dir)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Table returns the Table widget used by the FileDialog
func (d *FileDialog) Table() *gig.Table {

	return d.table
}

// SetOnselect sets a function which will be called when the user selects a file.
func (d *FileDialog) SetOnSelect(f func(d *FileDialog)) *FileDialog {

	d.onSelect = f
	return d
}

// SetOncancel sets a function with will be called the user cancels the Dialog
func (d *FileDialog) SetOnCancel(f func(d *FileDialog)) *FileDialog {

	d.onCancel = f
	return d
}

// ShowNewFolder shows button to create new folders.
// If the label is empty the button is hidden
func (d *FileDialog) ShowNewFolder(label string) *FileDialog {

	if label == "" {
		d.grpNewFolder.SetVisible(false)
		return d
	}
	d.grpNewFolder.SetVisible(true)
	d.btnNewFolder.SetLabel(label)
	return d
}

// Filepath returns the complete filename of the currently selected file
func (d *FileDialog) Filepath() string {

	fname := strings.Trim(d.fileInput.Text(), " ")
	return filepath.Join(d.path, fname)
}

// SetFile sets the current file name
func (d *FileDialog) SetFile(filename string) *FileDialog {

	d.fileInput.SetText(filename)
	return d
}

// SetFilters sets the current file filters
// Example of a filter string: *.go; *.c
func (d *FileDialog) SetFilters(filters string) *FileDialog {

	d.filterInput.SetText(filters)
	return d
}

func (d *FileDialog) AddPath(path string) error {

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	d.comboPath.AddChildren(gig.NewSelectable(absPath))
	return nil
}

func (d *FileDialog) SetPath(path string) error {

	// Read specified directory
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	d.path = path
	d.comboPath.SetPreview(path)

	// Get current filename filters
	filtext := strings.Trim(d.filterInput.Text(), "")
	filters := []string{}
	if len(filtext) > 0 {
		filters = strings.Split(filtext, ";")
	}

	// Updates table model with files from directory
	d.model.Clear()
	for _, f := range files {

		// Get file info
		info, err := f.Info()
		if err != nil {
			fmt.Println("info error", err)
			continue
		}
		fname := f.Name()
		fpath := filepath.Join(path, fname)
		isDir := info.IsDir()

		// If file is a symbolic link reads the target to check if it is a directory
		if info.Mode()&os.ModeSymlink != 0 {
			dstname, err := os.Readlink(fpath)
			if err != nil {
				continue
			}
			fstat, err := os.Stat(filepath.Join(path, dstname))
			if err != nil {
				continue
			}
			if fstat.IsDir() {
				isDir = true
			}
		}

		// Check if current file/directory matches the filters
		if !d.checkFilters(filters, fname, isDir) {
			continue
		}

		row := make([]*gig.Widget, 0)

		// File name selectable column
		if isDir {
			fname += "/"
		}
		fsel := gig.NewSelectable(fname).
			SetFlags(imgui.SelectableFlags_SpanAllColumns | imgui.SelectableFlags_AllowItemOverlap | imgui.SelectableFlags_AllowDoubleClick).
			SetOnDoubleclick(func(s *gig.Selectable) {
				if isDir {
					d.nextPath = filepath.Join(d.path, fname)
				} else {
					if d.onSelect != nil {
						d.onSelect(d)
					}
				}
			}).
			SetOnClick(func(s *gig.Selectable) {
			}).
			SetOnFocused(func(w *gig.Widget) {
				if !isDir {
					d.fileInput.SetText(fname)
				}
			})
		fsel.SetUserdata(fname)
		row = append(row, fsel.GetWidget())

		// File size column
		size := d.printer.Sprintf("%d", info.Size())
		txtSize := gig.NewText(size)
		txtSize.SetUserdata(info.Size())
		row = append(row, txtSize.GetWidget())

		// Modification time column
		txtTime := gig.NewText(info.ModTime().Format(d.timeFormat))
		txtTime.SetUserdata(info.ModTime())
		row = append(row, txtTime.GetWidget())

		d.model.AddRow(nil, row)
	}
	return nil
}

func (d *FileDialog) setHeader() {

	// Creates Combo for showing the current path and allow
	// selection from standard paths
	d.comboPath = gig.NewCombo("").SetOnChange(func(c *gig.Combo) {
		newPath := c.Selected().Label()
		err := d.SetPath(newPath)
		if err != nil {
			fmt.Println("Error setting path", newPath)
		}
	})

	// Creates group containing Button to allow creation of new folders
	d.grpNewFolder = gig.NewGroup()
	d.grpNewFolder.SetVisible(false)
	d.btnNewFolder = gig.NewButton("New Folder").
		SetOnClick(func(b *gig.Button) {
			d.dlgNewFolder.SetText("")
			d.dlgNewFolder.Open()
		})
	d.grpNewFolder.AddChildren(
		gig.NewSameLine(),
		d.btnNewFolder,
	)

	// Creates InpuText for entering file filters
	d.filterInput = gig.NewInputText("##filter").
		SetFlags(imgui.InputTextFlags_EnterReturnsTrue).
		SetOnChange(func(it *gig.InputText) {
			d.nextPath = d.path
		})
	d.checkHidden = gig.NewCheckbox("Show hidden").
		SetOnClick(func(c *gig.Checkbox) {
			d.nextPath = d.path
		})

	const offset = 70
	d.AddChildren(
		gig.NewText("Path:"),
		gig.NewSameLine().SetOffset(offset),
		d.comboPath,
		gig.NewSameLine(),
		gig.NewArrowButton(imgui.Dir_Up).
			SetOnClick(func(b *gig.ArrowButton) {
				d.upPath()
			}),
		d.grpNewFolder,
		gig.NewText("Filters:"),
		gig.NewSameLine().SetOffset(offset),
		d.filterInput,
		gig.NewSameLine(),
		d.checkHidden,
	)
}

func (d *FileDialog) setFooter() {

	const spaceY = 8
	filePrompt := gig.NewText("File:")
	filePrompt.SetPreRender(func(wi gig.IWidget) {
		d.footerStart = imgui.GetCursorPos()
		d.footerStart.Y += spaceY
		imgui.SetCursorPos(d.footerStart)
	})
	d.fileInput = gig.NewInputText("##filepath").
		SetFlags(imgui.InputTextFlags_EnterReturnsTrue).
		SetOnChange(func(it *gig.InputText) {
			if d.onSelect != nil {
				d.onSelect(d)
			}
		})
	d.fileInput.SetItemWidth(-1)

	const bwidth = 60
	d.btnOK = gig.NewButton("OK").
		SetSize(imgui.Vec2{bwidth, 0}).
		SetOnClick(func(b *gig.Button) {
			if d.onSelect != nil {
				d.onSelect(d)
			}
		})
	d.btnCancel = gig.NewButton("Cancel").
		SetSize(imgui.Vec2{bwidth, 0}).
		SetOnClick(func(b *gig.Button) {
			if d.onCancel != nil {
				d.onCancel(d)
			}
		})
	d.btnCancel.SetPostRender(func(wi gig.IWidget) {
		d.footerEnd = imgui.GetCursorPos()
		d.footerEnd.Y += spaceY
	})

	const offset = 50
	d.AddChildren(
		filePrompt,
		gig.NewSameLine().SetOffset(offset),
		d.fileInput,
		gig.NewSpacing(),
		gig.NewIndent(offset),
		d.btnOK, gig.NewSameLine(), d.btnCancel,
	)
}

func (d *FileDialog) setDlgNewFolder() {

	createDir := func(name string) {
		name = strings.Trim(name, " \t")
		if len(name) == 0 {
			return
		}
		path := filepath.Join(d.path, name)
		err := os.Mkdir(path, 0755)
		if err == nil {
			d.dlgNewFolder.Close()
			d.nextPath = path
			return
		}
		d.dlgMsg.SetTitle("Error")
		d.dlgMsg.SetMsg(fmt.Sprintf("Error creating folder:%s", err))
		d.dlgMsg.Open()
	}

	// Creates MsgDialog for error messages
	d.dlgMsg = NewMsgDialog("title", "msg", "OK")
	d.dlgMsg.SetOnClick(func(d *MsgDialog, idx int) {
		d.Close()
	})

	// Creates InputDialog for reading the new folder name
	d.dlgNewFolder = NewInputDialog("Create New Folder", "Folder name:", "OK", "Cancel").
		SetOnEnter(func(d *InputDialog) {
			createDir(d.Text())
		}).
		SetOnClick(func(d *InputDialog, idx int) {
			switch idx {
			case 0:
				createDir(d.Text())
			case 1:
				d.Close()
			}
		})
	d.dlgNewFolder.AddChildren(d.dlgMsg)
	d.AddChildren(d.dlgNewFolder)
}

func (d *FileDialog) upPath() {

	nextPath := filepath.Dir(d.path)
	if nextPath != d.path {
		d.nextPath = nextPath
	}
}

func (d *FileDialog) checkFilters(filters []string, filename string, isDir bool) bool {

	// Check hidden files or directories
	if strings.HasPrefix(filename, ".") {
		if !d.checkHidden.Value() {
			return false
		}
		return true
	}
	if isDir {
		return true
	}

	// Check other files
	if len(filters) == 0 {
		return true
	}
	for _, filter := range filters {
		res, err := filepath.Match(filter, filename)
		if err != nil {
			return true
		}
		if res {
			return true
		}
	}
	return false
}

func (d *FileDialog) sortTable() {

	fmt.Printf("sort table:%+v\n", d.sortSpecs)
	col := d.sortSpecs[0]
	rows := d.model.Rows()
	switch col.ColumnIndex {
	// Dir/file name
	case 0:
		sort.Slice(rows, func(i, j int) bool {
			fieldI := rows[i].Row[0].Userdata().(string)
			fieldJ := rows[j].Row[0].Userdata().(string)
			if col.Direction == imgui.SortDirection_Ascending {
				return fieldJ < fieldI
			} else {
				return fieldI < fieldJ
			}
		})
	// Dir/file size
	case 1:
		sort.Slice(rows, func(i, j int) bool {
			fieldI := rows[i].Row[1].Userdata().(int64)
			fieldJ := rows[j].Row[1].Userdata().(int64)
			if col.Direction == imgui.SortDirection_Ascending {
				return fieldJ < fieldI
			} else {
				return fieldI < fieldJ
			}
		})
	// Modification time
	case 2:
		sort.Slice(rows, func(i, j int) bool {
			fieldI := rows[i].Row[2].Userdata().(time.Time)
			fieldJ := rows[j].Row[2].Userdata().(time.Time)
			if col.Direction == imgui.SortDirection_Ascending {
				return fieldJ.Before(fieldI)
			} else {
				return fieldI.Before(fieldJ)
			}
		})
	}
	d.sortSpecs = nil
}

type FileDialogModal struct {
	gig.PopupModal
	D *FileDialog
}

//func NewFileDialogModal(title string, dir string, filters string) (*FileDialogModal, error) {
//
//	d := new(FileDialogModal)
//	d.PopupModal.Init(d, title)
//	d.AddChildren(gig.NewCombo("combo").AddChildren(
//		gig.NewSelectable("1"),
//		gig.NewSelectable("2"),
//	))
//
//	var err error
//	d.D, err = NewFileDialog(dir, filters)
//	if err != nil {
//		return nil, err
//	}
//	d.AddChildren(d.D)
//
//	//d := new(FileDialogModal)
//	//d.PopupModal.Init(d, title)
//	//var err error
//	//d.D, err = NewFileDialog(dir, filters)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//d.AddChildren(d.D)
//	//d.SetSize(imgui.Vec2{600, 400}, imgui.Cond_Once)
//	return d, nil
//}
