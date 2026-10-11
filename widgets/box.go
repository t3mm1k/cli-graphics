package widgets

import (
	"cli-graphics/engine"

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

	for _, child := range b.children {
		x, y := child.GetCoords()
		w, h := child.GetSize()
		childCanvas := canvas.SubCanvas(
			x+style.Padding.Left, 
			y+style.Padding.Top, 
			w-style.Padding.Left-style.Padding.Right, 
			h-style.Padding.Top-style.Padding.Bottom,
		)
		
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
		
		childComp, ok := engine.Registry.GetComponent(child.GetId())
		var childStyle engine.Style
		if ok {
			if baseComp, isBase := childComp.(interface{ CurrentStyle() engine.Style }); isBase {
				childStyle = baseComp.CurrentStyle()
			}
		}

		if style.FlexDirection == "row" {
			if childStyle.FlexGrow > 0 {
				growCount++
			} else {
				cw, _ := child.GetSize()
				fixedSizeSum += cw
			}
		} else { // "col"
			if childStyle.FlexGrow > 0 {
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

		childComp, ok := engine.Registry.GetComponent(child.GetId())
		var childStyle engine.Style
		if ok {
			if baseComp, isBase := childComp.(interface{ CurrentStyle() engine.Style }); isBase {
				childStyle = baseComp.CurrentStyle()
			}
		}

		cw, ch := child.GetSize()

		if style.FlexDirection == "row" {
			// В режиме row высота всех детей растягивается под полную высоту контента контейнера
			targetH := contentH
			targetW := cw

			if childStyle.FlexGrow > 0 {
				targetW = growSize
				if remainder > 0 {
					targetW++
					remainder--
				}
			}
			child.SetSize(targetW, targetH)
		} else { // "col"
			// В режиме col ширина всех детей растягивается под полную ширину контента контейнера
			targetW := contentW
			targetH := ch

			if childStyle.FlexGrow > 0 {
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