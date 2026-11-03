package widgets

import (
	"github.com/t3mm1k/cli-graphics/engine"
	"github.com/t3mm1k/cli-graphics/utils"
	"unicode/utf8"

	"github.com/google/uuid"
)

// ==================== НОВЫЙ КОД (ДОКУМЕНТИРОВАНИЕ ИИ) ====================

// Label представляет собой базовый графический TUI-виджет для отображения 
// статичного однострочного текста с возможностью стилизации, выравнивания и рамок.
type Label struct {
	engine.BaseComponent
	text string // Хранит отображаемую текстовую строку
}

// NewLabel инициализирует и регистрирует новый экземпляр текстовой метки.
// Параметры x, y задают начальные координаты, а флаг border определяет наличие рамки.
// Ширина и высота вычисляются автоматически в зависимости от длины текста и рамок.
func NewLabel(x, y int, border bool, text string, classes ...string) *Label {
	// Подсчет количества рун гарантирует корректную ширину для UTF-8 строк (например, кириллицы)
	width := utf8.RuneCountInString(text)
	// Корректировка размеров виджета: если есть рамка, добавляется по 2 ячейки к ширине и высоте
	width, height := utils.GetActuallySize(border, width, 1)

	id := uuid.New()
	lbl := &Label{
		BaseComponent: engine.NewBaseComponent(id, x, y, width, height),
		text:          text,
	}

	// Определение стилей по умолчанию на базе флага наличия рамки
	defaultClasses := "border-none fg:white text-left"
	if border {
		defaultClasses = "border-single border:dim-gray fg:white text-left"
	}
	lbl.InitStyle(defaultClasses, classes...)

	// Добавление компонента в глобальный реестр ядра движка для хит-тестинга и управления
	engine.Registry.AddComponent(lbl)

	return lbl
}

// SetText динамически обновляет текстовое содержимое метки и автоматически 
// пересчитывает габариты компонента, чтобы избежать обрезания текста.
func (l *Label) SetText(text string) {
	l.text = text
	width := utf8.RuneCountInString(text)
	border := l.Style.Border != engine.BorderNone
	width, height := utils.GetActuallySize(border, width, 1)
	l.SetSize(width, height)
}

// Render выполняет отрисовку текстовой метки на предоставленной проекции холста (Canvas).
// Метод учитывает текущие состояния стилей, внутренние отступы (padding) и типы выравнивания.
func (l *Label) Render(canvas *engine.Canvas) {
	style := l.CurrentStyle()
	w, h := l.GetSize()

	// Смещение текста относительно краев. Если рамка есть, текст сдвигается внутрь на 1 ячейку
	offset := 0
	if style.Border != engine.BorderNone {
		canvas.DrawRect(0, 0, w, h, style.Border, style.BorderFg, style.Bg)
		offset = 1
	}

	// Вычисление стартовой горизонтальной координаты отрисовки строки с учетом Padding и Align
	textX := offset + style.Padding.Left
	switch style.Align {
	case engine.TextAlignCenter:
		// Выравнивание по центру: добавляем к базовому смещению половину свободного пространства контейнера
		textX = offset + style.Padding.Left + (w-offset*2-style.Padding.Left-style.Padding.Right)/2
	case engine.TextAlignRight:
		// Выравнивание по правому краю: прижимаем текст к правой внутренней границе контейнера
		textX = w - offset - style.Padding.Right
	}
	textY := offset + style.Padding.Top

	// Посимвольный вывод строки на холст с применением ANSI-стилей текста и цветов
	canvas.DrawStringAligned(textX, textY, l.text, style.Fg, style.TextStyle, style.Align)
}

// OnTick выполняет логику компонента в рамках каждого игрового/системного тика.
// Для Label данный метод является пустым, так как компонент статичен.
func (l *Label) OnTick() {}

// ==================== КОНЕЦ НОВОГО КОДА ====================
