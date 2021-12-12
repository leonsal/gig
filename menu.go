package gig

import "github.com/leonsal/gig/imgui"

type MainMenuBar struct {
	Widget
}

// NewMainMenuBar creates and returns a pointer to a new MainMenuBar Widget.
func NewMainMenuBar() *MainMenuBar {

	m := new(MainMenuBar)
	m.Init(m)
	m.SetRender(func() {
		if imgui.BeginMainMenuBar() {
			m.RenderChildren()
			imgui.EndMainMenuBar()
		}
	})
	return m
}

// Menu widget is a vertical list of Menu items
type Menu struct {
	Widget
	label   imgui.CString
	enabled bool
}

// NewMenu creates and returns a pointer to a new Menu widget with the specified label
func NewMenu(label string) *Menu {

	m := new(Menu)
	m.Init(m)
	m.label.SetString(label)
	m.enabled = true
	m.SetRender(func() {
		if imgui.BeginMenuCS(m.label, m.enabled) {
			m.RenderChildren()
			imgui.EndMenu()
		}
	})
	return m
}

// Save saves this menu pointer into the specified destination.
// Returns the pointer to this menu to allow chaining.
func (m *Menu) Save(dst **Menu) *Menu {

	*dst = m
	return m
}

// SetEnabled sets the enabled state of this Menu.
// Returns the pointer to this menu to allow chaining.
func (m *Menu) SetEnabled(enabled bool) *Menu {

	m.enabled = enabled
	return m
}

type MenuItem struct {
	Widget
	label    imgui.CString
	shortcut imgui.CString
	selected bool
	enabled  bool
	onclick  func(*MenuItem)
}

// NewMenuItem creates and returns a pointer to a new MenuItem Widget with the specified label
func NewMenuItem(label string) *MenuItem {

	mi := new(MenuItem)
	mi.Init(mi)
	mi.label.SetString(label)
	mi.selected = false
	mi.enabled = true
	mi.SetRender(func() {
		if imgui.MenuItemCS(mi.label, mi.shortcut, mi.selected, mi.enabled) {
			if mi.onclick != nil {
				mi.onclick(mi)
			}
		}
	})
	return mi
}

// Save saves this menu item pointer into the specified destination.
// Returns the pointer to this menu item to allow chaining.
func (mi *MenuItem) Save(dst **MenuItem) *MenuItem {

	*dst = mi
	return mi
}

// SetShortcut sets the shortcut string for this MenuItem.
// Returns the pointer to this menu item to allow chaining.
func (mi *MenuItem) SetShortcut(s string) *MenuItem {

	mi.shortcut.SetString(s)
	return mi
}

// SetSelected sets the selected states of this MenuItem.
// Returns the pointer to this menu item to allow chaining.
func (mi *MenuItem) SetSelected(selected bool) *MenuItem {

	mi.selected = selected
	return mi
}

// Selected returns the selected state of this MenuItem.
func (mi *MenuItem) Selected() bool {

	return mi.selected
}

// SetEnabled sets the enabled state of this MenuItem.
// Returns the pointer to this menu item to allow chaining.
func (mi *MenuItem) SetEnabled(enabled bool) *MenuItem {

	mi.enabled = enabled
	return mi
}

// Enabled returns the enabled state of this MenuItem.
func (mi *MenuItem) Enabled() bool {

	return mi.enabled
}

// SetOnClick sets the callback for when the MenuItem is clicked.
// Returns the pointer to this menu item to allow chaining.
func (mi *MenuItem) SetOnClick(f func(*MenuItem)) *MenuItem {

	mi.onclick = f
	return mi
}
