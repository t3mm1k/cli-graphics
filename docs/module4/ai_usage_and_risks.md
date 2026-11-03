**Пример промпта:**
"Напиши модульные тесты на языке Go с использованием стандартного пакета `testing` и библиотеки `testify` для функции `ParseSGRMouse(data []byte) (*SGRMouseData, bool)`. Тесты должны быть написаны в виде Table-driven tests (табличных тестов). Обязательно включи:

1. Позитивные сценарии: клик левой кнопкой мыши, отпускание кнопки, движение мыши с зажатой кнопкой, прокрутка колесика (WheelUp/WheelDown).
2. Граничные и негативные сценарии: пустой слайс байт, некорректный префикс (не `\x1b[<`), отсутствие символов окончания `M` или `m`, неполные данные (передано меньше 3 координат через точку с запятой), а также проверку корректности вычитания 1 из координат X и Y (так как терминал передает 1-based, а функция преобразует в 0-based)."

**Полученный код:**
```
package utils

import (
	"testing"

	"://github.com"
)

// ==================== НОВЫЙ КОД (СГЕНЕРИРОВАН ИИ) ====================
func TestParseSGRMouse_Comprehensive(t *testing.T) {
	tests := []struct {
		name       string
		input      []byte
		wantValid  bool
		wantBtn    MouseBtn
		wantAction MouseAction
		wantX      int
		wantY      int
	}{
		// --- ПОЗИТИВНЫЕ СЦЕНАРИИ ---
		{
			name:       "Левый клик (нажатие) в точке 10,20 -> преобразуется в 9,19",
			input:      []byte("\x1b[<0;10;20M"),
			wantValid:  true,
			wantBtn:    MouseBtnLeft,
			wantAction: MouseActionPress,
			wantX:      9,
			wantY:      19,
		},
		{
			name:       "Левый клик (отпускание) в точке 10,20",
			input:      []byte("\x1b[<0;10;20m"),
			wantValid:  true,
			wantBtn:    MouseBtnLeft,
			wantAction: MouseActionRelease,
			wantX:      9,
			wantY:      19,
		},
		{
			name:       "Движение мыши с зажатой левой кнопкой (код 32)",
			input:      []byte("\x1b[<32;15;25M"),
			wantValid:  true,
			wantBtn:    MouseBtnLeft,
			wantAction: MouseActionMove,
			wantX:      14,
			wantY:      24,
		},
		{
			name:       "Колесико мыши вверх (Wheel Up, код 64)",
			input:      []byte("\x1b[<64;5;5M"),
			wantValid:  true,
			wantBtn:    MouseBtnWheelUp,
			wantAction: MouseActionPress,
			wantX:      4,
			wantY:      4,
		},
		{
			name:       "Колесико мыши вниз (Wheel Down, код 65)",
			input:      []byte("\x1b[<65;5;5M"),
			wantValid:  true,
			wantBtn:    MouseBtnWheelDown,
			wantAction: MouseActionPress,
			wantX:      4,
			wantY:      4,
		},

		// --- НЕГАТИВНЫЕ И ГРАНИЧНЫЕ СЦЕНАРИИ ---
		{
			name:      "Граничный случай: пустой слайс байт",
			input:     []byte{},
			wantValid: false,
		},
		{
			name:      "Негативный случай: неверный ANSI префикс",
			input:     []byte("\x1b[A12345M"),
			wantValid: false,
		},
		{
			name:      "Негативный случай: отсутствует финальная буква M/m",
			input:     []byte("\x1b[<0;10;20"),
			wantValid: false,
		},
		{
			name:      "Негативный случай: неполные координаты (нет Y)",
			input:     []byte("\x1b[<0;10M"),
			wantValid: false,
		},
		{
			name:      "Негативный случай: нечисловые данные в координатах",
			input:     []byte("\x1b[<0;abc;20M"),
			wantValid: false,
		},
		{
			name:       "Граничный случай: координаты равны 0 (проверка защиты от ухода в минус)",
			input:      []byte("\x1b[<0;0;0M"),
			wantValid:  true,
			wantBtn:    MouseBtnLeft,
			wantAction: MouseActionPress,
			wantX:      0, // Функция не должна вычитать 1, если координата уже 0
			wantY:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, ok := ParseSGRMouse(tt.input)
			if !tt.wantValid {
				assert.False(t, ok)
				assert.Nil(t, res)
			} else {
				assert.True(t, ok)
				assert.NotNil(t, res)
				assert.Equal(t, tt.wantBtn, res.Button)
				assert.Equal(t, tt.wantAction, res.Action)
				assert.Equal(t, tt.wantX, res.X)
				assert.Equal(t, tt.wantY, res.Y)
			}
		})
	}
}
// ==================== КОНЕЦ НОВОГО КОДА ====================
```

**Проблемы использования ИИ:**
1. Модель склонна галлюцинировать версиями репозиториев в URL-адресах вида ://github.com, генерируя несуществующие эндпоинты, из-за чего менеджер зависимостей Go (go mod tidy) выдавал ошибки сборки. Любые внешние ссылки и пути импорта пакетов требуют обязательной ручной валидации.
2. ИИ не способен качественно выполнять абстрактные или верхнеуровневые команды (например, "напиши тесты для мыши"). Без жесткого указания формата (Table-driven tests), используемых библиотек (stretchr/testify/assert) и явного перечисления граничных случаев (обработка пустых слайсов, нулевых координат, некорректных префиксов), модель генерирует поверхностные и избыточные тест-кейсы, которые пропускают критические логические ошибки в ядре. Качественный результат напрямую зависит от детализации и архитектурной точности промпта.
3. Главным ограничением моделей является лимит окна контекста и изоляция от среды выполнения. ИИ не «видит» проект целиком в реальном времени.