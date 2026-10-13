package widgets

import (
	"github.com/t3mm1k/cli-graphics/engine"

	"github.com/google/uuid"
)

type Box struct {
	engine.BaseComponent
	children []engine.Component
	focused  bool
}

const DefaultBoxClasses = "border-single border:dim-gray bg:default"

func (b *Box) IsFocused() bool {
	return b.focused
}

func NewBox(x, y, w, h int, classes ...string) *Box {
	id := uuid.New()

	box := &Box{
		BaseComponent: engine.NewBaseComponent(id, x, y, w, h),
	}

	box.InitStyle(DefaultBoxClasses, classes...)

	engine.Registry.AddComponent(box)
	return box
}

func (b *Box) Render(canvas *engine.Canvas) {
	style := b.CurrentStyle()
	w, h := b.GetSize()
	
	canvas.DrawRect(0, 0, w, h, style.Border, style.BorderFg, style.Bg)

	if style.IsFlex {
		b.computeFlexLayout(style)
	}

	offset := 1
	if style.Border == engine.BorderNone {
		offset = 0
	}

	currentX := offset + style.Padding.Left
	currentY := offset + style.Padding.Top

	for _, child := range b.children {
		if child == nil {
			continue
		}

		var x, y int
		cw, ch := child.GetSize()

		if style.IsFlex {
			x = currentX
			y = currentY

			if setter, ok := child.(interface{ SetCoords(x, y int) }); ok {
				setter.SetCoords(x, y)
			}

			if style.FlexDirection == "row" {
				currentX += cw
			} else { // "col"
				currentY += ch
			}
		} else {
			// Проверяем наличие методов GetX/GetY у компонента безопасно
			if getterX, ok := child.(interface{ GetX() int }); ok {
				if getterY, ok := child.(interface{ GetY() int }); ok {
					x = offset + style.Padding.Left + getterX.GetX()
					y = offset + style.Padding.Top + getterY.GetY()
				}
			} else {
				// Если методов GetX/GetY нет, возвращаем стандартный рабочий вариант ядра
				cx, cy := child.GetCoords()
				x = offset + style.Padding.Left + cx
				y = offset + style.Padding.Top + cy
			}
		}

		childCanvas := canvas.SubCanvas(x, y, cw, ch)
		
		child.Render(childCanvas)
	}
}

func (b *Box) OnTick() {
	for _, child := range b.children {
		child.OnTick()
	}
}

func (b *Box) AddChild(child engine.Component) {
	if child == nil {
		engine.Log.Warn("AddChild: attempt to add nil child, ignored", "err", engine.NilChildError)
		return
	}
	if child.GetId() == b.GetId() {
		panic(engine.CyclicDependencyError)
	}

	b.children = append(b.children, child)
	engine.Registry.SetParent(child.GetId(), b.GetId())
}

func (b *Box) SetFocus(focused bool) {
	b.focused = focused
}

func (b *Box) HandleKey(key string) bool {

	focusedId := engine.FocusManagerInstance.GetFocused()
	if focusedId == uuid.Nil {
		return false
	}

	comp, exists := engine.Registry.GetComponent(focusedId)
	if !exists {
		return false
	}

	if f, ok := comp.(engine.Focusable); ok {
		return f.HandleKey(key)
	}

	return false
}

func (b *Box) Children() []engine.Component {
	return b.children
}

func (b *Box) computeFlexLayout(style engine.Style) {
	if !style.IsFlex || len(b.children) == 0 {
		return
	}

	w, h := b.GetSize()
	offset := 1
	if style.Border == engine.BorderNone {
		offset = 0
	}

	contentW := w - (offset * 2) - style.Padding.Left - style.Padding.Right
	contentH := h - (offset * 2) - style.Padding.Top - style.Padding.Bottom

	if contentW <= 0 || contentH <= 0 {
		return
	}

	var fixedSizeSum int
	var growCount int

	for _, child := range b.children {
		if child == nil {
			continue
		}
		
		var isFlex1 bool
		if baseComp, ok := child.(interface{ GetId() uuid.UUID }); ok {
			if comp, exists := engine.Registry.GetComponent(baseComp.GetId()); exists {
				if reflectComp, hasStyle := comp.(interface{ CurrentStyle() engine.Style }); hasStyle {
					isFlex1 = reflectComp.CurrentStyle().FlexGrow > 0
					
					if directComp, isDirect := comp.(*Button); isDirect {
						isFlex1 = directComp.Style.FlexGrow > 0
					} else if directBox, isBox := comp.(*Box); isBox {
						isFlex1 = directBox.Style.FlexGrow > 0
					}
				}
			}
		}

		if style.FlexDirection == "row" {
			if isFlex1 {
				growCount++
			} else {
				cw, _ := child.GetSize()
				fixedSizeSum += cw
			}
		} else { // "col"
			if isFlex1 {
				growCount++
			} else {
				_, ch := child.GetSize()
				fixedSizeSum += ch
			}
		}
	}

	var growSize int
	var remainder int

	if style.FlexDirection == "row" {
		freeSpace := contentW - fixedSizeSum
		if freeSpace < 0 {
			freeSpace = 0
		}
		if growCount > 0 {
			growSize = freeSpace / growCount
			remainder = freeSpace % growCount
		}
	} else { // "col"
		freeSpace := contentH - fixedSizeSum
		if freeSpace < 0 {
			freeSpace = 0
		}
		if growCount > 0 {
			growSize = freeSpace / growCount
			remainder = freeSpace % growCount
		}
	}

	for _, child := range b.children {
		if child == nil {
			continue
		}

		var isFlex1 bool
		if directComp, isDirect := child.(*Button); isDirect {
			isFlex1 = directComp.Style.FlexGrow > 0
		} else if directBox, isBox := child.(*Box); isBox {
			isFlex1 = directBox.Style.FlexGrow > 0
		}

		cw, ch := child.GetSize()

		if style.FlexDirection == "row" {
			targetH := contentH
			targetW := cw

			if isFlex1 {
				targetW = growSize
				if remainder > 0 {
					targetW++
					remainder--
				}
			}
			child.SetSize(targetW, targetH)
		} else { // "col"
			targetW := contentW
			targetH := ch

			if isFlex1 {
				targetH = growSize
				if remainder > 0 {
					targetH++
					remainder--
				}
			}
			child.SetSize(targetW, targetH)
		}
	}
}