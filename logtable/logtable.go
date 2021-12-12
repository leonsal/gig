package logtable

import (
	"fmt"
	"time"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

type Level int

const (
	DEBUG Level = 0
	INFO  Level = 1
	WARN  Level = 2
	ERROR Level = 3
)

var levelNames = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
}

var levelColors = []imgui.Vec4{
	imgui.ColorName("Green"),
	imgui.ColorName("White"),
	imgui.ColorName("Yellow"),
	imgui.ColorName("Red"),
}

// Global singleton log instance
var instance *LogTable

// Get returns the single instance of the Log table widget
func Get() *LogTable {

	if instance == nil {
		instance = New()
	}
	return instance
}

// LogTable is a Table widget specialized for display of log events
// and containing its own data model
type LogTable struct {
	gig.Table
	model      *gig.TableModel
	timeFormat string
}

// New creates and returns a LogTable widget
func New() *LogTable {

	l := new(LogTable)
	l.Init(l)
	l.SetFlags(imgui.TableFlags_Resizable | imgui.TableFlags_ScrollY | imgui.TableFlags_SizingStretchSame)
	l.SetFreezeRows(1)
	l.AddColumns(
		gig.NewTableColumn("Time").SetInitWidth(2),
		gig.NewTableColumn("Level").SetInitWidth(1),
		gig.NewTableColumn("Message").SetInitWidth(8),
	)
	l.model = gig.NewTableModel()
	l.SetModel(l.model)
	l.timeFormat = "15:04:05.000"
	return l
}

func (l *LogTable) SetTimeFormat(format string) *LogTable {

	l.timeFormat = format
	return l
}

func (l *LogTable) Clear() {

	l.model.Clear()
}

func (l *LogTable) AppendLog(level Level, args ...interface{}) {

	l.appendLog(level, fmt.Sprint(args...))
}

func (l *LogTable) AppendLogf(level Level, format string, args ...interface{}) {

	l.appendLog(level, fmt.Sprintf(format, args...))
}

func (l *LogTable) Debug(args ...interface{}) {

	l.AppendLog(DEBUG, args...)
}

func (l *LogTable) Debugf(format string, args ...interface{}) {

	l.AppendLogf(DEBUG, format, args...)
}

func (l *LogTable) Info(args ...interface{}) {

	l.AppendLog(INFO, args...)
}

func (l *LogTable) Infof(format string, args ...interface{}) {

	l.AppendLogf(INFO, format, args...)
}

func (l *LogTable) Warn(args ...interface{}) {

	l.AppendLog(WARN, args...)
}

func (l *LogTable) Warnf(format string, args ...interface{}) {

	l.AppendLogf(WARN, format, args...)
}

func (l *LogTable) Error(args ...interface{}) {

	l.AppendLog(ERROR, args...)
}

func (l *LogTable) Errorf(format string, args ...interface{}) {

	l.AppendLogf(ERROR, format, args...)
}

func (l *LogTable) appendLog(level Level, text string) {

	ts := time.Now().Format(l.timeFormat)
	color := levelColors[level]
	row := make([]*gig.Widget, 0, 3)
	row = append(row, gig.NewTextColored(color, ts).GetWidget())
	row = append(row, gig.NewTextColored(color, levelNames[level]).GetWidget())
	row = append(row, gig.NewTextColored(color, text).GetWidget())
	l.model.AddRow(nil, row)
	l.ScrollToBottom()
}
