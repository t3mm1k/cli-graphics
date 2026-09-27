package engine

type Color struct {
	R, G, B   uint8
	IsDefault bool
}

func ColorDefault() Color {
	return Color{IsDefault: true}
}

func RGBToColor(r, g, b uint8) Color {
	return Color{r, g, b, false}
}

type Cell struct {
	R       rune
	FgColor Color
	BgColor Color
}

func NewCell(r rune) Cell {
	return Cell{
		R:       r,
		FgColor: ColorDefault(),
		BgColor: ColorDefault(),
	}
}

func NewCellColored(r rune, fg, bg Color) Cell {
	return Cell{
		R:       r,
		FgColor: fg,
		BgColor: bg,
	}
}

// Константы цветов, надо расширить)
var (
	ColorRed   = Color{R: 255, G: 0, B: 0, IsDefault: false}
	ColorGreen = Color{R: 0, G: 255, B: 0, IsDefault: false}
	ColorBlue  = Color{R: 0, G: 120, B: 255, IsDefault: false}
	ColorWhite = Color{R: 255, G: 255, B: 255, IsDefault: false}
	ColorBlack = Color{R: 0, G: 0, B: 0, IsDefault: false}
)

var (
	ColorDarkGray  = Color{R: 40, G: 40, B: 40, IsDefault: false}
	ColorDimGray   = Color{R: 110, G: 110, B: 110, IsDefault: false}
	ColorLightGray = Color{R: 190, G: 190, B: 190, IsDefault: false}
)

var (
	ColorNeonCyan  = Color{R: 0, G: 235, B: 215, IsDefault: false}
	ColorNeonPink  = Color{R: 255, G: 60, B: 170, IsDefault: false}
	ColorNeonGreen = Color{R: 50, G: 255, B: 120, IsDefault: false}
	ColorOrange    = Color{R: 255, G: 130, B: 0, IsDefault: false}
)

var (
	ColorPastelPurple = Color{R: 189, G: 147, B: 249, IsDefault: false}
	ColorPastelCyan   = Color{R: 139, G: 233, B: 253, IsDefault: false}
	ColorPastelGreen  = Color{R: 80, G: 250, B: 123, IsDefault: false}
)
