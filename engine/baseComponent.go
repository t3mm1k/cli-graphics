package engine

import "github.com/google/uuid"

type BaseComponent struct {
	id   uuid.UUID
	x, y int
	w, h int

	Style         Style
	FocusedStyle  Style
	DisabledStyle Style
	ActiveStyle   Style
	SelectedStyle Style
	HoverStyle    Style

	isDisabled bool
	isActive   bool
	isSelected bool
	isHovered  bool
}

func NewBaseComponent(id uuid.UUID, x, y, w, h int) BaseComponent {
	return BaseComponent{
		id: id,
		x:  x, y: y, w: w, h: h,
	}
}

func (b *BaseComponent) GetSize() (w, h int) {
	return b.w, b.h
}

func (b *BaseComponent) SetSize(w, h int) {
	if w < 0 {
		w = 0
	}
	if h < 0 {
		h = 0
	}
	b.w, b.h = w, h
}

func (b *BaseComponent) SetCoords(x, y int) {
	b.x, b.y = x, y
}

func (b *BaseComponent) GetId() uuid.UUID {
	return b.id
}

func (b *BaseComponent) GetCoords() (x, y int) {
	return b.x, b.y
}

func (b *BaseComponent) InitStyle(defaultClasses string, classes ...string) {
	allClasses := make([]string, 0, 1+len(classes))
	if defaultClasses != "" {
		allClasses = append(allClasses, defaultClasses)
	}
	allClasses = append(allClasses, classes...)
	b.Style, b.FocusedStyle, b.DisabledStyle, b.ActiveStyle, b.SelectedStyle, b.HoverStyle = ParseStyles(DefaultStyle(), allClasses)
}

func (b *BaseComponent) CurrentStyle() Style {
	if b.isDisabled {
		return b.DisabledStyle
	}
	if b.isActive {
		return b.ActiveStyle
	}
	if b.isHovered {
		return b.HoverStyle
	}
	if FocusManagerInstance != nil && FocusManagerInstance.GetFocused() == b.id {
		return b.FocusedStyle
	}
	if b.isSelected {
		return b.SelectedStyle
	}
	return b.Style
}

func (b *BaseComponent) SetDisabled(disabled bool) { b.isDisabled = disabled }
func (b *BaseComponent) IsDisabled() bool          { return b.isDisabled }
func (b *BaseComponent) SetActive(active bool)     { b.isActive = active }
func (b *BaseComponent) IsActive() bool            { return b.isActive }
func (b *BaseComponent) SetSelected(selected bool) { b.isSelected = selected }
func (b *BaseComponent) IsSelected() bool          { return b.isSelected }
func (b *BaseComponent) SetHovered(hovered bool)   { b.isHovered = hovered }
func (b *BaseComponent) IsHovered() bool           { return b.isHovered }
