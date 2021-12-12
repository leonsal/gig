package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

type WindowFlags C.ImGuiWindowFlags

const (
	WindowFlags_None                      WindowFlags = C.ImGuiWindowFlags_None
	WindowFlags_NoTitleBar                WindowFlags = C.ImGuiWindowFlags_NoTitleBar
	WindowFlags_NoResize                  WindowFlags = C.ImGuiWindowFlags_NoResize
	WindowFlags_NoMove                    WindowFlags = C.ImGuiWindowFlags_NoMove
	WindowFlags_NoScrollbar               WindowFlags = C.ImGuiWindowFlags_NoScrollbar
	WindowFlags_NoScrollWithMounse        WindowFlags = C.ImGuiWindowFlags_NoScrollWithMouse
	WindowFlags_NoCollapse                WindowFlags = C.ImGuiWindowFlags_NoCollapse
	WindowFlags_AlwaysAutoResize          WindowFlags = C.ImGuiWindowFlags_AlwaysAutoResize
	WindowFlags_NoBackground              WindowFlags = C.ImGuiWindowFlags_NoBackground
	WindowFlags_NoSavedSettings           WindowFlags = C.ImGuiWindowFlags_NoSavedSettings
	WindowFlags_NoMouseInputs             WindowFlags = C.ImGuiWindowFlags_NoMouseInputs
	WindowFlags_MenuBar                   WindowFlags = C.ImGuiWindowFlags_MenuBar
	WindowFlags_HorizontalScrollbar       WindowFlags = C.ImGuiWindowFlags_HorizontalScrollbar
	WindowFlags_NoFocusOnAppearing        WindowFlags = C.ImGuiWindowFlags_NoFocusOnAppearing
	WindowFlags_NoBringToFrontOnFocus     WindowFlags = C.ImGuiWindowFlags_NoBringToFrontOnFocus
	WindowFlags_AlwaysVerticalScrollbar   WindowFlags = C.ImGuiWindowFlags_AlwaysVerticalScrollbar
	WindowFlags_AlwaysHorizontalScrollbar WindowFlags = C.ImGuiWindowFlags_AlwaysHorizontalScrollbar
	WindowFlags_AlwaysUseWindowPadding    WindowFlags = C.ImGuiWindowFlags_AlwaysUseWindowPadding
	WindowFlags_NoNavInputs               WindowFlags = C.ImGuiWindowFlags_NoNavInputs
	WindowFlags_NoNavFocus                WindowFlags = C.ImGuiWindowFlags_NoNavFocus
	WindowFlags_UnsavedDocument           WindowFlags = C.ImGuiWindowFlags_UnsavedDocument
	WindowFlags_NoNav                     WindowFlags = C.ImGuiWindowFlags_NoNavInputs | C.ImGuiWindowFlags_NoNavFocus
	WindowFlags_NoDecoration              WindowFlags = C.ImGuiWindowFlags_NoTitleBar | C.ImGuiWindowFlags_NoResize | C.ImGuiWindowFlags_NoScrollbar | C.ImGuiWindowFlags_NoCollapse
	WindowFlags_NoInputs                  WindowFlags = C.ImGuiWindowFlags_NoMouseInputs | C.ImGuiWindowFlags_NoNavInputs | C.ImGuiWindowFlags_NoNavFocus
	WindowFlags_NavFlattened              WindowFlags = C.ImGuiWindowFlags_NavFlattened
	WindowFlags_ChildWindow               WindowFlags = C.ImGuiWindowFlags_ChildWindow
	WindowFlags_Tooltip                   WindowFlags = C.ImGuiWindowFlags_Tooltip
	WindowFlags_Popup                     WindowFlags = C.ImGuiWindowFlags_Popup
	WindowFlags_Modal                     WindowFlags = C.ImGuiWindowFlags_Modal
	WindowFlags_ChildMenu                 WindowFlags = C.ImGuiWindowFlags_ChildMenu
)

type InputTextFlags C.ImGuiInputTextFlags

const (
	InputTextFlags_None                InputTextFlags = C.ImGuiInputTextFlags_None
	InputTextFlags_CharsDecimal        InputTextFlags = C.ImGuiInputTextFlags_CharsDecimal
	InputTextFlags_CharsHexadecimal    InputTextFlags = C.ImGuiInputTextFlags_CharsHexadecimal
	InputTextFlags_CharsUppercase      InputTextFlags = C.ImGuiInputTextFlags_CharsUppercase
	InputTextFlags_CharsNoBlank        InputTextFlags = C.ImGuiInputTextFlags_CharsNoBlank
	InputTextFlags_AutoSelectAll       InputTextFlags = C.ImGuiInputTextFlags_AutoSelectAll
	InputTextFlags_EnterReturnsTrue    InputTextFlags = C.ImGuiInputTextFlags_EnterReturnsTrue
	InputTextFlags_CallbackCompletion  InputTextFlags = C.ImGuiInputTextFlags_CallbackCompletion
	InputTextFlags_CallbackHistory     InputTextFlags = C.ImGuiInputTextFlags_CallbackHistory
	InputTextFlags_CallbackAlways      InputTextFlags = C.ImGuiInputTextFlags_CallbackAlways
	InputTextFlags_CallbackCharFilter  InputTextFlags = C.ImGuiInputTextFlags_CallbackCharFilter
	InputTextFlags_AllowTabInput       InputTextFlags = C.ImGuiInputTextFlags_AllowTabInput
	InputTextFlags_CtrlEnterForNewLine InputTextFlags = C.ImGuiInputTextFlags_CtrlEnterForNewLine
	InputTextFlags_NoHorizontalScroll  InputTextFlags = C.ImGuiInputTextFlags_NoHorizontalScroll
	InputTextFlags_AlwaysOverwrite     InputTextFlags = C.ImGuiInputTextFlags_AlwaysOverwrite
	InputTextFlags_ReadOnly            InputTextFlags = C.ImGuiInputTextFlags_ReadOnly
	InputTextFlags_Password            InputTextFlags = C.ImGuiInputTextFlags_Password
	InputTextFlags_NoUndoRedo          InputTextFlags = C.ImGuiInputTextFlags_NoUndoRedo
	InputTextFlags_CharsScientific     InputTextFlags = C.ImGuiInputTextFlags_CharsScientific
	InputTextFlags_CallbackResize      InputTextFlags = C.ImGuiInputTextFlags_CallbackResize
	InputTextFlags_CallbackEdit        InputTextFlags = C.ImGuiInputTextFlags_CallbackEdit
	InputTextFlags_Multiline           InputTextFlags = C.ImGuiInputTextFlags_Multiline
	InputTextFlags_NoMarkEdited        InputTextFlags = C.ImGuiInputTextFlags_NoMarkEdited
	InputTextFlags_MergedItem          InputTextFlags = C.ImGuiInputTextFlags_MergedItem
)

type TreeNodeFlags C.ImGuiTreeNodeFlags

const (
	TreeNodeFlags_None                 TreeNodeFlags = C.ImGuiTreeNodeFlags_None
	TreeNodeFlags_Selected             TreeNodeFlags = C.ImGuiTreeNodeFlags_Selected
	TreeNodeFlags_Framed               TreeNodeFlags = C.ImGuiTreeNodeFlags_Framed
	TreeNodeFlags_AllowItemOverlap     TreeNodeFlags = C.ImGuiTreeNodeFlags_AllowItemOverlap
	TreeNodeFlags_NoTreePushOnOpen     TreeNodeFlags = C.ImGuiTreeNodeFlags_NoTreePushOnOpen
	TreeNodeFlags_NoAutoOpenOnLog      TreeNodeFlags = C.ImGuiTreeNodeFlags_NoAutoOpenOnLog
	TreeNodeFlags_DefaultOpen          TreeNodeFlags = C.ImGuiTreeNodeFlags_DefaultOpen
	TreeNodeFlags_OpenOnDoubleClick    TreeNodeFlags = C.ImGuiTreeNodeFlags_OpenOnDoubleClick
	TreeNodeFlags_OpenOnArrow          TreeNodeFlags = C.ImGuiTreeNodeFlags_OpenOnArrow
	TreeNodeFlags_Leaf                 TreeNodeFlags = C.ImGuiTreeNodeFlags_Leaf
	TreeNodeFlags_Bullet               TreeNodeFlags = C.ImGuiTreeNodeFlags_Bullet
	TreeNodeFlags_FramePadding         TreeNodeFlags = C.ImGuiTreeNodeFlags_FramePadding
	TreeNodeFlags_SpanAvailWidth       TreeNodeFlags = C.ImGuiTreeNodeFlags_SpanAvailWidth
	TreeNodeFlags_SpanFullWidth        TreeNodeFlags = C.ImGuiTreeNodeFlags_SpanFullWidth
	TreeNodeFlags_NavLeftJumpsBackHere TreeNodeFlags = C.ImGuiTreeNodeFlags_NavLeftJumpsBackHere
	TreeNodeFlags_CollapsingHeader     TreeNodeFlags = C.ImGuiTreeNodeFlags_Framed | C.ImGuiTreeNodeFlags_NoTreePushOnOpen | C.ImGuiTreeNodeFlags_NoAutoOpenOnLog
)

type PopupFlags C.ImGuiPopupFlags

const (
	PopupFlags_None                    PopupFlags = C.ImGuiPopupFlags_None
	PopupFlags_MouseButtonLeft         PopupFlags = C.ImGuiPopupFlags_MouseButtonLeft
	PopupFlags_MouseButtonRight        PopupFlags = C.ImGuiPopupFlags_MouseButtonRight
	PopupFlags_MouseButtonMiddle       PopupFlags = C.ImGuiPopupFlags_MouseButtonMiddle
	PopupFlags_MouseButtonMask_        PopupFlags = C.ImGuiPopupFlags_MouseButtonMask_
	PopupFlags_MouseButtonDefault_     PopupFlags = C.ImGuiPopupFlags_MouseButtonDefault_
	PopupFlags_NoOpenOverExistingPopup PopupFlags = C.ImGuiPopupFlags_NoOpenOverExistingPopup
	PopupFlags_NoOpenOverItems         PopupFlags = C.ImGuiPopupFlags_NoOpenOverItems
	PopupFlags_AnyPopupId              PopupFlags = C.ImGuiPopupFlags_AnyPopupId
	PopupFlags_AnyPopupLevel           PopupFlags = C.ImGuiPopupFlags_AnyPopupLevel
	PopupFlags_AnyPopup                PopupFlags = C.ImGuiPopupFlags_AnyPopup
)

type SelectableFlags C.ImGuiSelectableFlags

const (
	SelectableFlags_None             SelectableFlags = C.ImGuiSelectableFlags_None
	SelectableFlags_DontClosePopups  SelectableFlags = C.ImGuiSelectableFlags_DontClosePopups
	SelectableFlags_SpanAllColumns   SelectableFlags = C.ImGuiSelectableFlags_SpanAllColumns
	SelectableFlags_AllowDoubleClick SelectableFlags = C.ImGuiSelectableFlags_AllowDoubleClick
	SelectableFlags_Disabled         SelectableFlags = C.ImGuiSelectableFlags_Disabled
	SelectableFlags_AllowItemOverlap SelectableFlags = C.ImGuiSelectableFlags_AllowItemOverlap
)

type ComboFlags C.ImGuiComboFlags

const (
	ComboFlags_None           ComboFlags = C.ImGuiComboFlags_None
	ComboFlags_PopupAlignLeft ComboFlags = C.ImGuiComboFlags_PopupAlignLeft
	ComboFlags_HeightSmall    ComboFlags = C.ImGuiComboFlags_HeightSmall
	ComboFlags_HeightRegular  ComboFlags = C.ImGuiComboFlags_HeightRegular
	ComboFlags_HeightLarge    ComboFlags = C.ImGuiComboFlags_HeightLarge
	ComboFlags_HeightLargest  ComboFlags = C.ImGuiComboFlags_HeightLargest
	ComboFlags_NoArrowButton  ComboFlags = C.ImGuiComboFlags_NoArrowButton
	ComboFlags_NoPreview      ComboFlags = C.ImGuiComboFlags_NoPreview
)

type TabBarFlags C.ImGuiTabBarFlags

const (
	TabBarFlags_None                         TabBarFlags = C.ImGuiTabBarFlags_None
	TabBarFlags_Reorderable                  TabBarFlags = C.ImGuiTabBarFlags_Reorderable
	TabBarFlags_AutoSelectNewTabs            TabBarFlags = C.ImGuiTabBarFlags_AutoSelectNewTabs
	TabBarFlags_TabListPopupButton           TabBarFlags = C.ImGuiTabBarFlags_TabListPopupButton
	TabBarFlags_NoCloseWithMiddleMouseButton TabBarFlags = C.ImGuiTabBarFlags_NoCloseWithMiddleMouseButton
	TabBarFlags_NoTabListScrollingButtons    TabBarFlags = C.ImGuiTabBarFlags_NoTabListScrollingButtons
	TabBarFlags_NoTooltip                    TabBarFlags = C.ImGuiTabBarFlags_NoTooltip
	TabBarFlags_FittingPolicyResizeDown      TabBarFlags = C.ImGuiTabBarFlags_FittingPolicyResizeDown
	TabBarFlags_FittingPolicyScroll          TabBarFlags = C.ImGuiTabBarFlags_FittingPolicyScroll
)

type TabItemFlags C.ImGuiTabItemFlags

const (
	TabItemFlags_None                         TabItemFlags = C.ImGuiTabItemFlags_None
	TabItemFlags_UnsavedDocument              TabItemFlags = C.ImGuiTabItemFlags_UnsavedDocument
	TabItemFlags_SetSelected                  TabItemFlags = C.ImGuiTabItemFlags_SetSelected
	TabItemFlags_NoCloseWithMiddleMouseButton TabItemFlags = C.ImGuiTabItemFlags_NoCloseWithMiddleMouseButton
	TabItemFlags_NoPushId                     TabItemFlags = C.ImGuiTabItemFlags_NoPushId
	TabItemFlags_NoTooltip                    TabItemFlags = C.ImGuiTabItemFlags_NoTooltip
	TabItemFlags_NoReorder                    TabItemFlags = C.ImGuiTabItemFlags_NoReorder
	TabItemFlags_Leading                      TabItemFlags = C.ImGuiTabItemFlags_Leading
	TabItemFlags_Trailing                     TabItemFlags = C.ImGuiTabItemFlags_Trailing
)

type TableFlags C.ImGuiTableFlags

const (
	TableFlags_None                       TableFlags = C.ImGuiTableFlags_None
	TableFlags_Resizable                  TableFlags = C.ImGuiTableFlags_Resizable
	TableFlags_Reorderable                TableFlags = C.ImGuiTableFlags_Reorderable
	TableFlags_Hideable                   TableFlags = C.ImGuiTableFlags_Hideable
	TableFlags_Sortable                   TableFlags = C.ImGuiTableFlags_Sortable
	TableFlags_NoSavedSettings            TableFlags = C.ImGuiTableFlags_NoSavedSettings
	TableFlags_ContextMenuInBody          TableFlags = C.ImGuiTableFlags_ContextMenuInBody
	TableFlags_RowBg                      TableFlags = C.ImGuiTableFlags_RowBg
	TableFlags_BordersInnerH              TableFlags = C.ImGuiTableFlags_BordersInnerH
	TableFlags_BordersOuterH              TableFlags = C.ImGuiTableFlags_BordersOuterH
	TableFlags_BordersInnerV              TableFlags = C.ImGuiTableFlags_BordersInnerV
	TableFlags_BordersOuterV              TableFlags = C.ImGuiTableFlags_BordersOuterV
	TableFlags_BordersH                   TableFlags = C.ImGuiTableFlags_BordersH
	TableFlags_BordersV                   TableFlags = C.ImGuiTableFlags_BordersV
	TableFlags_BordersInner               TableFlags = C.ImGuiTableFlags_BordersInner
	TableFlags_BordersOuter               TableFlags = C.ImGuiTableFlags_BordersOuter
	TableFlags_Borders                    TableFlags = C.ImGuiTableFlags_Borders
	TableFlags_NoBordersInBody            TableFlags = C.ImGuiTableFlags_NoBordersInBody
	TableFlags_NoBordersInBodyUntilResize TableFlags = C.ImGuiTableFlags_NoBordersInBodyUntilResize
	TableFlags_SizingFixedFit             TableFlags = C.ImGuiTableFlags_SizingFixedFit
	TableFlags_SizingFixedSame            TableFlags = C.ImGuiTableFlags_SizingFixedSame
	TableFlags_SizingStretchProp          TableFlags = C.ImGuiTableFlags_SizingStretchProp
	TableFlags_SizingStretchSame          TableFlags = C.ImGuiTableFlags_SizingStretchSame
	TableFlags_NoHostExtendX              TableFlags = C.ImGuiTableFlags_NoHostExtendX
	TableFlags_NoHostExtendY              TableFlags = C.ImGuiTableFlags_NoHostExtendY
	TableFlags_NoKeepColumnsVisible       TableFlags = C.ImGuiTableFlags_NoKeepColumnsVisible
	TableFlags_PreciseWidths              TableFlags = C.ImGuiTableFlags_PreciseWidths
	TableFlags_NoClip                     TableFlags = C.ImGuiTableFlags_NoClip
	TableFlags_PadOuterX                  TableFlags = C.ImGuiTableFlags_PadOuterX
	TableFlags_NoPadOuterX                TableFlags = C.ImGuiTableFlags_NoPadOuterX
	TableFlags_NoPadInnerX                TableFlags = C.ImGuiTableFlags_NoPadInnerX
	TableFlags_ScrollX                    TableFlags = C.ImGuiTableFlags_ScrollX
	TableFlags_ScrollY                    TableFlags = C.ImGuiTableFlags_ScrollY
	TableFlags_SortMulti                  TableFlags = C.ImGuiTableFlags_SortMulti
	TableFlags_SortTristate               TableFlags = C.ImGuiTableFlags_SortTristate
)

type TableColumnFlags C.ImGuiTableColumnFlags

const (
	TableColumnFlags_None                 TableColumnFlags = C.ImGuiTableColumnFlags_None
	TableColumnFlags_DefaultHide          TableColumnFlags = C.ImGuiTableColumnFlags_DefaultHide
	TableColumnFlags_DefaultSort          TableColumnFlags = C.ImGuiTableColumnFlags_DefaultSort
	TableColumnFlags_WidthStretch         TableColumnFlags = C.ImGuiTableColumnFlags_WidthStretch
	TableColumnFlags_WidthFixed           TableColumnFlags = C.ImGuiTableColumnFlags_WidthFixed
	TableColumnFlags_NoResize             TableColumnFlags = C.ImGuiTableColumnFlags_NoResize
	TableColumnFlags_NoReorder            TableColumnFlags = C.ImGuiTableColumnFlags_NoReorder
	TableColumnFlags_NoHide               TableColumnFlags = C.ImGuiTableColumnFlags_NoHide
	TableColumnFlags_NoClip               TableColumnFlags = C.ImGuiTableColumnFlags_NoClip
	TableColumnFlags_NoSort               TableColumnFlags = C.ImGuiTableColumnFlags_NoSort
	TableColumnFlags_NoSortAscending      TableColumnFlags = C.ImGuiTableColumnFlags_NoSortAscending
	TableColumnFlags_NoSortDescending     TableColumnFlags = C.ImGuiTableColumnFlags_NoSortDescending
	TableColumnFlags_NoHeaderWidth        TableColumnFlags = C.ImGuiTableColumnFlags_NoHeaderWidth
	TableColumnFlags_PreferSortAscending  TableColumnFlags = C.ImGuiTableColumnFlags_PreferSortAscending
	TableColumnFlags_PreferSortDescending TableColumnFlags = C.ImGuiTableColumnFlags_PreferSortDescending
	TableColumnFlags_IndentEnable         TableColumnFlags = C.ImGuiTableColumnFlags_IndentEnable
	TableColumnFlags_IndentDisable        TableColumnFlags = C.ImGuiTableColumnFlags_IndentDisable
	TableColumnFlags_IsEnabled            TableColumnFlags = C.ImGuiTableColumnFlags_IsEnabled
	TableColumnFlags_IsVisible            TableColumnFlags = C.ImGuiTableColumnFlags_IsVisible
	TableColumnFlags_IsSorted             TableColumnFlags = C.ImGuiTableColumnFlags_IsSorted
	TableColumnFlags_IsHovered            TableColumnFlags = C.ImGuiTableColumnFlags_IsHovered
)

type TableRowFlags C.ImGuiTableRowFlags

const (
	TableRowFlags_None    TableRowFlags = C.ImGuiTableRowFlags_None
	TableRowFlags_Headers TableRowFlags = C.ImGuiTableRowFlags_Headers
)

type TableBgTarget C.ImGuiTableBgTarget

const (
	TableBgTarget_None   TableBgTarget = C.ImGuiTableBgTarget_None
	TableBgTarget_RowBg0 TableBgTarget = C.ImGuiTableBgTarget_RowBg0
	TableBgTarget_RowBg1 TableBgTarget = C.ImGuiTableBgTarget_RowBg1
	TableBgTarget_CellBg TableBgTarget = C.ImGuiTableBgTarget_CellBg
)

type FocusedFlags C.ImGuiFocusedFlags

const (
	FocusedFlags_None                FocusedFlags = C.ImGuiFocusedFlags_None
	FocusedFlags_ChildWindows        FocusedFlags = C.ImGuiFocusedFlags_ChildWindows
	FocusedFlags_RootWindow          FocusedFlags = C.ImGuiFocusedFlags_RootWindow
	FocusedFlags_AnyWindow           FocusedFlags = C.ImGuiFocusedFlags_AnyWindow
	FocusedFlags_RootAndChildWindows FocusedFlags = C.ImGuiFocusedFlags_RootAndChildWindows
)

type HoveredFlags C.ImGuiHoveredFlags

const (
	HoveredFlags_None                         HoveredFlags = C.ImGuiHoveredFlags_None
	HoveredFlags_ChildWindows                 HoveredFlags = C.ImGuiHoveredFlags_ChildWindows
	HoveredFlags_RootWindow                   HoveredFlags = C.ImGuiHoveredFlags_RootWindow
	HoveredFlags_AnyWindow                    HoveredFlags = C.ImGuiHoveredFlags_AnyWindow
	HoveredFlags_AllowWhenBlockedByPopup      HoveredFlags = C.ImGuiHoveredFlags_AllowWhenBlockedByPopup
	HoveredFlags_AllowWhenBlockedByActiveItem HoveredFlags = C.ImGuiHoveredFlags_AllowWhenBlockedByActiveItem
	HoveredFlags_AllowWhenOverlapped          HoveredFlags = C.ImGuiHoveredFlags_AllowWhenOverlapped
	HoveredFlags_AllowWhenDisabled            HoveredFlags = C.ImGuiHoveredFlags_AllowWhenDisabled
	HoveredFlags_RectOnly                     HoveredFlags = C.ImGuiHoveredFlags_RectOnly
	HoveredFlags_RootAndChildWindows          HoveredFlags = C.ImGuiHoveredFlags_RootAndChildWindows
)

type DragDropFlags C.ImGuiDragDropFlags

const (
	DragDropFlags_None                     DragDropFlags = C.ImGuiDragDropFlags_None
	DragDropFlags_SourceNoPreviewTooltip   DragDropFlags = C.ImGuiDragDropFlags_SourceNoPreviewTooltip
	DragDropFlags_SourceNoDisableHover     DragDropFlags = C.ImGuiDragDropFlags_SourceNoDisableHover
	DragDropFlags_SourceNoHoldToOpenOthers DragDropFlags = C.ImGuiDragDropFlags_SourceNoHoldToOpenOthers
	DragDropFlags_SourceAllowNullID        DragDropFlags = C.ImGuiDragDropFlags_SourceAllowNullID
	DragDropFlags_SourceExtern             DragDropFlags = C.ImGuiDragDropFlags_SourceExtern
	DragDropFlags_SourceAutoExpirePayload  DragDropFlags = C.ImGuiDragDropFlags_SourceAutoExpirePayload
	DragDropFlags_AcceptBeforeDelivery     DragDropFlags = C.ImGuiDragDropFlags_AcceptBeforeDelivery
	DragDropFlags_AcceptNoDrawDefaultRect  DragDropFlags = C.ImGuiDragDropFlags_AcceptNoDrawDefaultRect
	DragDropFlags_AcceptNoPreviewTooltip   DragDropFlags = C.ImGuiDragDropFlags_AcceptNoPreviewTooltip
	DragDropFlags_AcceptPeekOnly           DragDropFlags = C.ImGuiDragDropFlags_AcceptPeekOnly
)

type DataType C.ImGuiDataType

const (
	DataType_S8     DataType = C.ImGuiDataType_S8
	DataType_U8     DataType = C.ImGuiDataType_U8
	DataType_S16    DataType = C.ImGuiDataType_S16
	DataType_U16    DataType = C.ImGuiDataType_U16
	DataType_S32    DataType = C.ImGuiDataType_S32
	DataType_U32    DataType = C.ImGuiDataType_U32
	DataType_S64    DataType = C.ImGuiDataType_S64
	DataType_U64    DataType = C.ImGuiDataType_U64
	DataType_Float  DataType = C.ImGuiDataType_Float
	DataType_Double DataType = C.ImGuiDataType_Double
	DataType_COUNT  DataType = C.ImGuiDataType_COUNT
)

type Dir C.ImGuiDir

const (
	Dir_None  Dir = C.ImGuiDir_None
	Dir_Left  Dir = C.ImGuiDir_Left
	Dir_Right Dir = C.ImGuiDir_Right
	Dir_Up    Dir = C.ImGuiDir_Up
	Dir_Down  Dir = C.ImGuiDir_Down
)

type SortDirection C.ImGuiSortDirection

const (
	SortDirection_None       SortDirection = C.ImGuiSortDirection_None
	SortDirection_Ascending  SortDirection = C.ImGuiSortDirection_Ascending
	SortDirection_Descending SortDirection = C.ImGuiSortDirection_Descending
)

type Key C.ImGuiKey

const (
	Key_Tab         Key = C.ImGuiKey_Tab
	Key_LeftArrow   Key = C.ImGuiKey_LeftArrow
	Key_RightArrow  Key = C.ImGuiKey_RightArrow
	Key_UpArrow     Key = C.ImGuiKey_UpArrow
	Key_DownArrow   Key = C.ImGuiKey_DownArrow
	Key_PageUp      Key = C.ImGuiKey_PageUp
	Key_PageDown    Key = C.ImGuiKey_PageDown
	Key_Home        Key = C.ImGuiKey_Home
	Key_End         Key = C.ImGuiKey_End
	Key_Insert      Key = C.ImGuiKey_Insert
	Key_Delete      Key = C.ImGuiKey_Delete
	Key_Backspace   Key = C.ImGuiKey_Backspace
	Key_Space       Key = C.ImGuiKey_Space
	Key_Enter       Key = C.ImGuiKey_Enter
	Key_Escape      Key = C.ImGuiKey_Escape
	Key_KeyPadEnter Key = C.ImGuiKey_KeyPadEnter
	Key_A           Key = C.ImGuiKey_A
	Key_C           Key = C.ImGuiKey_C
	Key_V           Key = C.ImGuiKey_V
	Key_X           Key = C.ImGuiKey_X
	Key_Y           Key = C.ImGuiKey_Y
	Key_Z           Key = C.ImGuiKey_Z
)

type KeyModFlags C.ImGuiKeyModFlags

const (
	KeyModFlags_None  KeyModFlags = C.ImGuiKeyModFlags_None
	KeyModFlags_Ctrl  KeyModFlags = C.ImGuiKeyModFlags_Ctrl
	KeyModFlags_Shift KeyModFlags = C.ImGuiKeyModFlags_Shift
	KeyModFlags_Alt   KeyModFlags = C.ImGuiKeyModFlags_Alt
	KeyModFlags_Super KeyModFlags = C.ImGuiKeyModFlags_Super
)

type NavInput C.ImGuiNavInput

const (
	NavInput_Activate    NavInput = C.ImGuiNavInput_Activate
	NavInput_Cancel      NavInput = C.ImGuiNavInput_Cancel
	NavInput_Input       NavInput = C.ImGuiNavInput_Input
	NavInput_Menu        NavInput = C.ImGuiNavInput_Menu
	NavInput_DpadLeft    NavInput = C.ImGuiNavInput_DpadLeft
	NavInput_DpadRight   NavInput = C.ImGuiNavInput_DpadRight
	NavInput_DpadUp      NavInput = C.ImGuiNavInput_DpadUp
	NavInput_DpadDown    NavInput = C.ImGuiNavInput_DpadDown
	NavInput_LStickLeft  NavInput = C.ImGuiNavInput_LStickLeft
	NavInput_LStickRight NavInput = C.ImGuiNavInput_LStickRight
	NavInput_LStickUp    NavInput = C.ImGuiNavInput_LStickUp
	NavInput_LStickDown  NavInput = C.ImGuiNavInput_LStickDown
	NavInput_FocusPrev   NavInput = C.ImGuiNavInput_FocusPrev
	NavInput_FocusNext   NavInput = C.ImGuiNavInput_FocusNext
	NavInput_TweakSlow   NavInput = C.ImGuiNavInput_TweakSlow
	NavInput_TweakFast   NavInput = C.ImGuiNavInput_TweakFast
	NavInput_KeyLeft_    NavInput = C.ImGuiNavInput_KeyLeft_
	NavInput_KeyRight_   NavInput = C.ImGuiNavInput_KeyRight_
	NavInput_KeyUp_      NavInput = C.ImGuiNavInput_KeyUp_
	NavInput_KeyDown_    NavInput = C.ImGuiNavInput_KeyDown_
)

type ConfigFlags C.ImGuiConfigFlags

const (
	ConfigFlags_None                 ConfigFlags = C.ImGuiConfigFlags_None
	ConfigFlags_NavEnableKeyboard    ConfigFlags = C.ImGuiConfigFlags_NavEnableKeyboard
	ConfigFlags_NavEnableGamepad     ConfigFlags = C.ImGuiConfigFlags_NavEnableGamepad
	ConfigFlags_NavEnableSetMousePos ConfigFlags = C.ImGuiConfigFlags_NavEnableSetMousePos
	ConfigFlags_NavNoCaptureKeyboard ConfigFlags = C.ImGuiConfigFlags_NavNoCaptureKeyboard
	ConfigFlags_NoMouse              ConfigFlags = C.ImGuiConfigFlags_NoMouse
	ConfigFlags_NoMouseCursorChange  ConfigFlags = C.ImGuiConfigFlags_NoMouseCursorChange
	ConfigFlags_IsSRGB               ConfigFlags = C.ImGuiConfigFlags_IsSRGB
	ConfigFlags_IsTouchScreen        ConfigFlags = C.ImGuiConfigFlags_IsTouchScreen
)

type BackendFlags C.ImGuiBackendFlags

const (
	BackendFlags_None                 BackendFlags = C.ImGuiBackendFlags_None
	BackendFlags_HasGamepad           BackendFlags = C.ImGuiBackendFlags_HasGamepad
	BackendFlags_HasMouseCursors      BackendFlags = C.ImGuiBackendFlags_HasMouseCursors
	BackendFlags_HasSetMousePos       BackendFlags = C.ImGuiBackendFlags_HasSetMousePos
	BackendFlags_RendererHasVtxOffset BackendFlags = C.ImGuiBackendFlags_RendererHasVtxOffset
)

type Col C.ImGuiCol

const (
	Col_Text                  Col = C.ImGuiCol_Text
	Col_TextDisabled          Col = C.ImGuiCol_TextDisabled
	Col_WindowBg              Col = C.ImGuiCol_WindowBg
	Col_ChildBg               Col = C.ImGuiCol_ChildBg
	Col_PopupBg               Col = C.ImGuiCol_PopupBg
	Col_Border                Col = C.ImGuiCol_Border
	Col_BorderShadow          Col = C.ImGuiCol_BorderShadow
	Col_FrameBg               Col = C.ImGuiCol_FrameBg
	Col_FrameBgHovered        Col = C.ImGuiCol_FrameBgHovered
	Col_FrameBgActive         Col = C.ImGuiCol_FrameBgActive
	Col_TitleBg               Col = C.ImGuiCol_TitleBg
	Col_TitleBgActive         Col = C.ImGuiCol_TitleBgActive
	Col_TitleBgCollapsed      Col = C.ImGuiCol_TitleBgCollapsed
	Col_MenuBarBg             Col = C.ImGuiCol_MenuBarBg
	Col_ScrollbarBg           Col = C.ImGuiCol_ScrollbarBg
	Col_ScrollbarGrab         Col = C.ImGuiCol_ScrollbarGrab
	Col_ScrollbarGrabHovered  Col = C.ImGuiCol_ScrollbarGrabHovered
	Col_ScrollbarGrabActive   Col = C.ImGuiCol_ScrollbarGrabActive
	Col_CheckMark             Col = C.ImGuiCol_CheckMark
	Col_SliderGrab            Col = C.ImGuiCol_SliderGrab
	Col_SliderGrabActive      Col = C.ImGuiCol_SliderGrabActive
	Col_Button                Col = C.ImGuiCol_Button
	Col_ButtonHovered         Col = C.ImGuiCol_ButtonHovered
	Col_ButtonActive          Col = C.ImGuiCol_ButtonActive
	Col_Header                Col = C.ImGuiCol_Header
	Col_HeaderHovered         Col = C.ImGuiCol_HeaderHovered
	Col_HeaderActive          Col = C.ImGuiCol_HeaderActive
	Col_Separator             Col = C.ImGuiCol_Separator
	Col_SeparatorHovered      Col = C.ImGuiCol_SeparatorHovered
	Col_SeparatorActive       Col = C.ImGuiCol_SeparatorActive
	Col_ResizeGrip            Col = C.ImGuiCol_ResizeGrip
	Col_ResizeGripHovered     Col = C.ImGuiCol_ResizeGripHovered
	Col_ResizeGripActive      Col = C.ImGuiCol_ResizeGripActive
	Col_Tab                   Col = C.ImGuiCol_Tab
	Col_TabHovered            Col = C.ImGuiCol_TabHovered
	Col_TabActive             Col = C.ImGuiCol_TabActive
	Col_TabUnfocused          Col = C.ImGuiCol_TabUnfocused
	Col_TabUnfocusedActive    Col = C.ImGuiCol_TabUnfocusedActive
	Col_PlotLines             Col = C.ImGuiCol_PlotLines
	Col_PlotLinesHovered      Col = C.ImGuiCol_PlotLinesHovered
	Col_PlotHistogram         Col = C.ImGuiCol_PlotHistogram
	Col_PlotHistogramHovered  Col = C.ImGuiCol_PlotHistogramHovered
	Col_TableHeaderBg         Col = C.ImGuiCol_TableHeaderBg
	Col_TableBorderStrong     Col = C.ImGuiCol_TableBorderStrong
	Col_TableBorderLight      Col = C.ImGuiCol_TableBorderLight
	Col_TableRowBg            Col = C.ImGuiCol_TableRowBg
	Col_TableRowBgAlt         Col = C.ImGuiCol_TableRowBgAlt
	Col_TextSelectedBg        Col = C.ImGuiCol_TextSelectedBg
	Col_DragDropTarget        Col = C.ImGuiCol_DragDropTarget
	Col_NavHighlight          Col = C.ImGuiCol_NavHighlight
	Col_NavWindowingHighlight Col = C.ImGuiCol_NavWindowingHighlight
	Col_NavWindowingDimBg     Col = C.ImGuiCol_NavWindowingDimBg
	Col_ModalWindowDimBg      Col = C.ImGuiCol_ModalWindowDimBg
)

type StyleVar C.ImGuiStyleVar

const (
	StyleVar_Alpha               StyleVar = C.ImGuiStyleVar_Alpha
	StyleVar_WindowPadding       StyleVar = C.ImGuiStyleVar_WindowPadding
	StyleVar_WindowRounding      StyleVar = C.ImGuiStyleVar_WindowRounding
	StyleVar_WindowBorderSize    StyleVar = C.ImGuiStyleVar_WindowBorderSize
	StyleVar_WindowMinSize       StyleVar = C.ImGuiStyleVar_WindowMinSize
	StyleVar_WindowTitleAlign    StyleVar = C.ImGuiStyleVar_WindowTitleAlign
	StyleVar_ChildRounding       StyleVar = C.ImGuiStyleVar_ChildRounding
	StyleVar_ChildBorderSize     StyleVar = C.ImGuiStyleVar_ChildBorderSize
	StyleVar_PopupRounding       StyleVar = C.ImGuiStyleVar_PopupRounding
	StyleVar_PopupBorderSize     StyleVar = C.ImGuiStyleVar_PopupBorderSize
	StyleVar_FramePadding        StyleVar = C.ImGuiStyleVar_FramePadding
	StyleVar_FrameRounding       StyleVar = C.ImGuiStyleVar_FrameRounding
	StyleVar_FrameBorderSize     StyleVar = C.ImGuiStyleVar_FrameBorderSize
	StyleVar_ItemSpacing         StyleVar = C.ImGuiStyleVar_ItemSpacing
	StyleVar_ItemInnerSpacing    StyleVar = C.ImGuiStyleVar_ItemInnerSpacing
	StyleVar_IndentSpacing       StyleVar = C.ImGuiStyleVar_IndentSpacing
	StyleVar_CellPadding         StyleVar = C.ImGuiStyleVar_CellPadding
	StyleVar_ScrollbarSize       StyleVar = C.ImGuiStyleVar_ScrollbarSize
	StyleVar_ScrollbarRounding   StyleVar = C.ImGuiStyleVar_ScrollbarRounding
	StyleVar_GrabMinSize         StyleVar = C.ImGuiStyleVar_GrabMinSize
	StyleVar_GrabRounding        StyleVar = C.ImGuiStyleVar_GrabRounding
	StyleVar_TabRounding         StyleVar = C.ImGuiStyleVar_TabRounding
	StyleVar_ButtonTextAlign     StyleVar = C.ImGuiStyleVar_ButtonTextAlign
	StyleVar_SelectableTextAlign StyleVar = C.ImGuiStyleVar_SelectableTextAlign
)

type ButtonFlags C.ImGuiButtonFlags

const (
	ButtonFlags_None                ButtonFlags = C.ImGuiButtonFlags_None
	ButtonFlags_MouseButtonLeft     ButtonFlags = C.ImGuiButtonFlags_MouseButtonLeft
	ButtonFlags_MouseButtonRight    ButtonFlags = C.ImGuiButtonFlags_MouseButtonRight
	ButtonFlags_MouseButtonMiddle   ButtonFlags = C.ImGuiButtonFlags_MouseButtonMiddle
	ButtonFlags_MouseButtonMask_    ButtonFlags = C.ImGuiButtonFlags_MouseButtonMask_
	ButtonFlags_MouseButtonDefault_ ButtonFlags = C.ImGuiButtonFlags_MouseButtonDefault_
)

type ColorEditFlags C.ImGuiColorEditFlags

const (
	ColorEditFlags_None             ColorEditFlags = C.ImGuiColorEditFlags_None
	ColorEditFlags_NoAlpha          ColorEditFlags = C.ImGuiColorEditFlags_NoAlpha
	ColorEditFlags_NoPicker         ColorEditFlags = C.ImGuiColorEditFlags_NoPicker
	ColorEditFlags_NoOptions        ColorEditFlags = C.ImGuiColorEditFlags_NoOptions
	ColorEditFlags_NoSmallPreview   ColorEditFlags = C.ImGuiColorEditFlags_NoSmallPreview
	ColorEditFlags_NoInputs         ColorEditFlags = C.ImGuiColorEditFlags_NoInputs
	ColorEditFlags_NoTooltip        ColorEditFlags = C.ImGuiColorEditFlags_NoTooltip
	ColorEditFlags_NoLabel          ColorEditFlags = C.ImGuiColorEditFlags_NoLabel
	ColorEditFlags_NoSidePreview    ColorEditFlags = C.ImGuiColorEditFlags_NoSidePreview
	ColorEditFlags_NoDragDrop       ColorEditFlags = C.ImGuiColorEditFlags_NoDragDrop
	ColorEditFlags_NoBorder         ColorEditFlags = C.ImGuiColorEditFlags_NoBorder
	ColorEditFlags_AlphaBar         ColorEditFlags = C.ImGuiColorEditFlags_AlphaBar
	ColorEditFlags_AlphaPreview     ColorEditFlags = C.ImGuiColorEditFlags_AlphaPreview
	ColorEditFlags_AlphaPreviewHalf ColorEditFlags = C.ImGuiColorEditFlags_AlphaPreviewHalf
	ColorEditFlags_HDR              ColorEditFlags = C.ImGuiColorEditFlags_HDR
	ColorEditFlags_DisplayRGB       ColorEditFlags = C.ImGuiColorEditFlags_DisplayRGB
	ColorEditFlags_DisplayHSV       ColorEditFlags = C.ImGuiColorEditFlags_DisplayHSV
	ColorEditFlags_DisplayHex       ColorEditFlags = C.ImGuiColorEditFlags_DisplayHex
	ColorEditFlags_Uint8            ColorEditFlags = C.ImGuiColorEditFlags_Uint8
	ColorEditFlags_Float            ColorEditFlags = C.ImGuiColorEditFlags_Float
	ColorEditFlags_PickerHueBar     ColorEditFlags = C.ImGuiColorEditFlags_PickerHueBar
	ColorEditFlags_PickerHueWheel   ColorEditFlags = C.ImGuiColorEditFlags_PickerHueWheel
	ColorEditFlags_InputRGB         ColorEditFlags = C.ImGuiColorEditFlags_InputRGB
	ColorEditFlags_InputHSV         ColorEditFlags = C.ImGuiColorEditFlags_InputHSV
)

type SliderFlags C.ImGuiSliderFlags

const (
	SliderFlags_None            SliderFlags = C.ImGuiSliderFlags_None
	SliderFlags_AlwaysClamp     SliderFlags = C.ImGuiSliderFlags_AlwaysClamp
	SliderFlags_Logarithmic     SliderFlags = C.ImGuiSliderFlags_Logarithmic
	SliderFlags_NoRoundToFormat SliderFlags = C.ImGuiSliderFlags_NoRoundToFormat
	SliderFlags_NoInput         SliderFlags = C.ImGuiSliderFlags_NoInput
	SliderFlags_InvalidMask_    SliderFlags = C.ImGuiSliderFlags_InvalidMask_
)

type MouseButton C.ImGuiMouseButton

const (
	MouseButton_Left   MouseButton = C.ImGuiMouseButton_Left
	MouseButton_Right  MouseButton = C.ImGuiMouseButton_Right
	MouseButton_Middle MouseButton = C.ImGuiMouseButton_Middle
	MouseButton_COUNT  MouseButton = C.ImGuiMouseButton_COUNT
)

type MouseCursor C.ImGuiMouseCursor

const (
	MouseCursor_None       MouseCursor = C.ImGuiMouseCursor_None
	MouseCursor_Arrow      MouseCursor = C.ImGuiMouseCursor_Arrow
	MouseCursor_TextInput  MouseCursor = C.ImGuiMouseCursor_TextInput
	MouseCursor_ResizeAll  MouseCursor = C.ImGuiMouseCursor_ResizeAll
	MouseCursor_ResizeNS   MouseCursor = C.ImGuiMouseCursor_ResizeNS
	MouseCursor_ResizeEW   MouseCursor = C.ImGuiMouseCursor_ResizeEW
	MouseCursor_ResizeNESW MouseCursor = C.ImGuiMouseCursor_ResizeNESW
	MouseCursor_ResizeNWSE MouseCursor = C.ImGuiMouseCursor_ResizeNWSE
	MouseCursor_Hand       MouseCursor = C.ImGuiMouseCursor_Hand
	MouseCursor_NotAllowed MouseCursor = C.ImGuiMouseCursor_NotAllowed
	MouseCursor_COUNT      MouseCursor = C.ImGuiMouseCursor_COUNT
)

type Cond C.ImGuiCond

const (
	Cond_None         Cond = C.ImGuiCond_None
	Cond_Always       Cond = C.ImGuiCond_Always
	Cond_Once         Cond = C.ImGuiCond_Once
	Cond_FirstUseEver Cond = C.ImGuiCond_FirstUseEver
	Cond_Appearing    Cond = C.ImGuiCond_Appearing
)

type ViewportFlags C.ImGuiViewportFlags

const (
	ViewportFlags_Node              = C.ImGuiViewportFlags_None
	ViewportFlags_IsPlatformWindow  = C.ImGuiViewportFlags_IsPlatformWindow
	ViewportFlags_IsPlatformMonitor = C.ImGuiViewportFlags_IsPlatformMonitor
	ViewportFlags_OwnedByApp        = C.ImGuiViewportFlags_OwnedByApp
)

type TableColumnSortSpecs struct {
	ColumnUserID ID            // User id of the column (if specified by a TableSetupColumn() call)
	ColumnIndex  int           // Index of the columnt
	SortOrder    int           // Index within parent ImGuiTableSortSpecs
	Direction    SortDirection // SortDirection_Ascending or SortDirection_Descending
}

type TableSortSpecs struct {
	Specs      []TableColumnSortSpecs
	SpecsDirty bool // Set to true when specs have changed since last time! Use this to sort again, then clear the flag.
}

type InputTextCallbackData struct {
	c *C.ImGuiInputTextCallbackData
}

func (cbd InputTextCallbackData) EventFlag() InputTextFlags {

	return InputTextFlags(cbd.c.EventFlag)
}

func (cbd InputTextCallbackData) Flags() InputTextFlags {

	return InputTextFlags(cbd.c.Flags)
}

func (cbd InputTextCallbackData) EventChar() Wchar {

	return Wchar(cbd.c.EventChar)
}

func (cbd InputTextCallbackData) EventKey() Key {

	return Key(cbd.c.EventKey)
}

func (cbd InputTextCallbackData) tBuf() string {

	return C.GoString(cbd.c.Buf)
}

func (cbd InputTextCallbackData) BufTextLen() int {

	return int(cbd.c.BufTextLen)
}

func (cbd InputTextCallbackData) BufSize() int {

	return int(cbd.c.BufSize)
}

func (cbd InputTextCallbackData) BufDirty() bool {

	return bool(cbd.c.BufDirty)
}

func (cbd InputTextCallbackData) CursorPos() int {

	return int(cbd.c.CursorPos)
}

func (cbd InputTextCallbackData) SelectionStart() int {

	return int(cbd.c.SelectionStart)
}

func (cbd InputTextCallbackData) SelectionEnd() int {

	return int(cbd.c.SelectionEnd)
}

func (cbd InputTextCallbackData) DeleteChars(pos, bytesCount int) {

	C.ImGuiInputTextCallbackData_DeleteChars(cbd.c, C.int(pos), C.int(bytesCount))
}

func (cbd InputTextCallbackData) InsertChars(pos int, text string) {

	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.ImGuiInputTextCallbackData_InsertChars(cbd.c, C.int(pos), ctext, nil)
}

func (cbd *InputTextCallbackData) SelectAll() {

	C.ImGuiInputTextCallbackData_SelectAll(cbd.c)
}

func (cbd *InputTextCallbackData) ClearSelection() {

	C.ImGuiInputTextCallbackData_ClearSelection(cbd.c)
}

func (cbd *InputTextCallbackData) HasSelection() bool {

	return bool(C.ImGuiInputTextCallbackData_HasSelection(cbd.c))
}

type ListClipper C.ImGuiListClipper

func NewListClipper() *ListClipper {

	return (*ListClipper)(C.ImGuiListClipper_ImGuiListClipper())
}

func (lc *ListClipper) Destroy() {

	C.ImGuiListClipper_destroy((*C.ImGuiListClipper)(lc))
}

func (lc *ListClipper) Begin(itemsCount int, itemsHeight float32) {

	C.ImGuiListClipper_Begin((*C.ImGuiListClipper)(lc), C.int(itemsCount), C.float(itemsHeight))
}

func (lc *ListClipper) End() {

	C.ImGuiListClipper_End((*C.ImGuiListClipper)(lc))
}

func (lc *ListClipper) Step() bool {

	return bool(C.ImGuiListClipper_Step((*C.ImGuiListClipper)(lc)))
}

type InputTextCallback func(InputTextCallbackData) int

type Context C.ImGuiContext

type TextureID C.ImTextureID

type ID C.ImGuiID

type Payload C.ImGuiPayload

type Storage C.ImGuiStorage

//type DrawList C.ImDrawList

type DrawListSharedData C.ImDrawListSharedData

type U32 C.ImU32

type Wchar C.ImWchar

type Color C.ImColor

type Vec2 struct {
	X float32
	Y float32
}

func (v *Vec2) ImVec2() C.ImVec2 {

	return C.ImVec2{C.float(v.X), C.float(v.Y)}
}

func (v *Vec2) Set(src *C.ImVec2) {

	v.X = float32(src.x)
	v.Y = float32(src.y)
}

func (v C.ImVec2) Vec2() Vec2 {

	return Vec2{float32(v.x), float32(v.y)}
}

type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (v *Vec4) ImVec4() C.ImVec4 {

	return C.ImVec4{C.float(v.X), C.float(v.Y), C.float(v.Z), C.float(v.W)}
}

func (v C.ImVec4) Vec4() Vec4 {

	return Vec4{float32(v.x), float32(v.y), float32(v.z), float32(v.w)}
}
