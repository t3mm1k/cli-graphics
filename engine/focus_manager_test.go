package engine

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Тестовый компонент, реализующий Focusable
type mockFocusable struct {
	BaseComponent
	focused bool
	handled []string
}

func newMockFocusable(id uuid.UUID) *mockFocusable {
	return &mockFocusable{
		BaseComponent: NewBaseComponent(id, 0, 0, 10, 1),
	}
}

func (m *mockFocusable) SetFocus(f bool) {
	m.focused = f
}

func (m *mockFocusable) IsFocused() bool {
	return m.focused
}

func (m *mockFocusable) HandleKey(key string) bool {
	m.handled = append(m.handled, key)
	return true
}

func (m *mockFocusable) Render(canvas *Canvas) {}
func (m *mockFocusable) OnTick()                {}

func TestFocusManager_Empty(t *testing.T) {
	fm := &FocusManager{current: -1}

	assert.Equal(t, uuid.Nil, fm.GetFocused())

	// Проверяем, что вызовы на пустом списке не вызывают панику
	assert.NotPanics(t, func() {
		fm.FocusNext()
		fm.FocusPrev()
	})
	assert.Equal(t, uuid.Nil, fm.GetFocused())
}

func TestFocusManager_Cycle(t *testing.T) {
	fm := &FocusManager{current: -1}

	id1 := uuid.New()
	id2 := uuid.New()
	id3 := uuid.New()

	comp1 := newMockFocusable(id1)
	comp2 := newMockFocusable(id2)
	comp3 := newMockFocusable(id3)

	Registry.AddComponent(comp1)
	Registry.AddComponent(comp2)
	Registry.AddComponent(comp3)
	defer func() {
		Registry.RemoveComponent(id1)
		Registry.RemoveComponent(id2)
		Registry.RemoveComponent(id3)
	}()

	fm.Register(id1)
	fm.Register(id2)
	fm.Register(id3)

	// 1. Первый FocusNext должен выбрать первый элемент (id1)
	fm.FocusNext()
	assert.Equal(t, id1, fm.GetFocused())
	assert.True(t, comp1.IsFocused())
	assert.False(t, comp2.IsFocused())
	assert.False(t, comp3.IsFocused())

	// 2. Второй FocusNext -> id2
	fm.FocusNext()
	assert.Equal(t, id2, fm.GetFocused())
	assert.False(t, comp1.IsFocused())
	assert.True(t, comp2.IsFocused())
	assert.False(t, comp3.IsFocused())

	// 3. Третий FocusNext -> id3
	fm.FocusNext()
	assert.Equal(t, id3, fm.GetFocused())
	assert.False(t, comp1.IsFocused())
	assert.False(t, comp2.IsFocused())
	assert.True(t, comp3.IsFocused())

	// 4. Четвертый FocusNext -> зацикливается обратно на id1!
	fm.FocusNext()
	assert.Equal(t, id1, fm.GetFocused())
	assert.True(t, comp1.IsFocused())
	assert.False(t, comp2.IsFocused())
	assert.False(t, comp3.IsFocused())

	// 5. FocusPrev -> идет назад на id3
	fm.FocusPrev()
	assert.Equal(t, id3, fm.GetFocused())
	assert.False(t, comp1.IsFocused())
	assert.True(t, comp3.IsFocused())

	// 6. FocusPrev -> на id2
	fm.FocusPrev()
	assert.Equal(t, id2, fm.GetFocused())
	assert.True(t, comp2.IsFocused())
}

func TestFocusManager_SetFocused(t *testing.T) {
	fm := &FocusManager{current: -1}

	id1 := uuid.New()
	id2 := uuid.New()

	comp1 := newMockFocusable(id1)
	comp2 := newMockFocusable(id2)

	Registry.AddComponent(comp1)
	Registry.AddComponent(comp2)
	defer func() {
		Registry.RemoveComponent(id1)
		Registry.RemoveComponent(id2)
	}()

	fm.Register(id1)
	fm.Register(id2)

	fm.SetFocused(id2)
	assert.Equal(t, id2, fm.GetFocused())
	assert.False(t, comp1.IsFocused())
	assert.True(t, comp2.IsFocused())

	fm.SetFocused(id1)
	assert.Equal(t, id1, fm.GetFocused())
	assert.True(t, comp1.IsFocused())
	assert.False(t, comp2.IsFocused())
}
