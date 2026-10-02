package engine

import "image/color"

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

type TextStyle uint8

const (
    TextStyleDefault TextStyle = iota
    TextStyleBold
    TextStyleDim
    TextStyleItalic
    TextStyleUnderline
    TextStyleBlink
    TextStyleReverse
    TextStyleHidden
    TextStyleStrikethrough
)

type Cell struct {
    R       rune
    FgColor Color
    BgColor Color
    Style   TextStyle
}

func NewCell(r rune) Cell {
    return Cell{
        R:       r,
        FgColor: ColorDefault(),
        BgColor: ColorDefault(),
        Style:   TextStyleDefault,
    }
}

func NewCellColored(r rune, fg, bg Color) Cell {
    return Cell{
        R:       r,
        FgColor: fg,
        BgColor: bg,
        Style:   TextStyleDefault,
    }
}

func ToEngineColor(c color.Color) Color {
	r, g, b, _ := c.RGBA()
	return Color{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8)}
}

type ColorRegistry struct {
	baseColors map[string]Color
}

func (cr *ColorRegistry) Get(name string) (Color, bool) {
	col, ok := cr.baseColors[name]
	return col, ok
}

func (cr *ColorRegistry) Set(name string, c Color) Color {
	cr.baseColors[name] = c
	return c
}

var BaseColors = ColorRegistry{
	baseColors: map[string]Color{
		// ==========================================
		// 1. Монохром, темнота и градации серого
		// ==========================================
		"black":       {R: 0, G: 0, B: 0, IsDefault: false},
		"true-black":  {R: 0, G: 0, B: 0, IsDefault: false},
		"charcoal":    {R: 26, G: 26, B: 26, IsDefault: false},
		"dark-gray":   {R: 40, G: 40, B: 40, IsDefault: false},
		"gray":        {R: 110, G: 110, B: 110, IsDefault: false},
		"dim-gray":    {R: 110, G: 110, B: 110, IsDefault: false},
		"mid-gray":    {R: 128, G: 128, B: 128, IsDefault: false},
		"light-gray":  {R: 190, G: 190, B: 190, IsDefault: false},
		"silver":      {R: 192, G: 192, B: 192, IsDefault: false},
		"slate":       {R: 100, G: 116, B: 139, IsDefault: false},
		"zinc":        {R: 113, G: 113, B: 122, IsDefault: false},
		"neutral":     {R: 115, G: 115, B: 115, IsDefault: false},
		"stone":       {R: 120, G: 113, B: 108, IsDefault: false},
		"white":       {R: 255, G: 255, B: 255, IsDefault: false},
		"smoke":       {R: 245, G: 245, B: 245, IsDefault: false},
		"ghost-white": {R: 248, G: 248, B: 255, IsDefault: false},
		// ==========================================
		// 2. Красные и винные оттенки
		// ==========================================
		"red":       {R: 255, G: 0, B: 0, IsDefault: false},
		"crimson":   {R: 220, G: 20, B: 60, IsDefault: false},
		"scarlet":   {R: 255, G: 36, B: 0, IsDefault: false},
		"ruby":      {R: 224, G: 17, B: 95, IsDefault: false},
		"maroon":    {R: 128, G: 0, B: 0, IsDefault: false},
		"burgundy":  {R: 144, G: 0, B: 32, IsDefault: false},
		"wine":      {R: 114, G: 47, B: 55, IsDefault: false},
		"brick":     {R: 178, G: 34, B: 34, IsDefault: false},
		"cherry":    {R: 222, G: 49, B: 99, IsDefault: false},
		// ==========================================
		// 3. Оранжевые, коричневые и земляные
		// ==========================================
		"orange":    {R: 255, G: 130, B: 0, IsDefault: false},
		"tangerine": {R: 242, G: 133, B: 0, IsDefault: false},
		"coral":     {R: 255, G: 127, B: 80, IsDefault: false},
		"peach":     {R: 255, G: 218, B: 185, IsDefault: false},
		"apricot":   {R: 251, G: 206, B: 177, IsDefault: false},
		"rust":      {R: 183, G: 65, B: 14, IsDefault: false},
		"bronze":    {R: 205, G: 127, B: 50, IsDefault: false},
		"copper":    {R: 184, G: 115, B: 51, IsDefault: false},
		"brown":     {R: 139, G: 69, B: 19, IsDefault: false},
		"chocolate": {R: 123, G: 63, B: 0, IsDefault: false},
		"coffee":    {R: 111, G: 78, B: 55, IsDefault: false},
		"caramel":   {R: 198, G: 142, B: 23, IsDefault: false},
		"sand":      {R: 194, G: 178, B: 128, IsDefault: false},
		// ==========================================
		// 4. Жёлтые и золотые
		// ==========================================
		"yellow":  {R: 255, G: 255, B: 0, IsDefault: false},
		"gold":    {R: 255, G: 215, B: 0, IsDefault: false},
		"amber":   {R: 245, G: 158, B: 11, IsDefault: false},
		"lemon":   {R: 255, G: 247, B: 0, IsDefault: false},
		"canary":  {R: 255, G: 239, B: 0, IsDefault: false},
		"mustard": {R: 255, G: 219, B: 88, IsDefault: false},
		"cream":   {R: 255, G: 253, B: 208, IsDefault: false},
		"vanilla": {R: 243, G: 229, B: 171, IsDefault: false},
		// ==========================================
		// 5. Зелёные и природные
		// ==========================================
		"green":      {R: 0, G: 255, B: 0, IsDefault: false},
		"lime":       {R: 132, G: 204, B: 22, IsDefault: false},
		"emerald":    {R: 16, G: 185, B: 129, IsDefault: false},
		"jade":       {R: 0, G: 168, B: 107, IsDefault: false},
		"mint":       {R: 167, G: 243, B: 208, IsDefault: false},
		"sage":       {R: 158, G: 170, B: 144, IsDefault: false},
		"forest":     {R: 34, G: 139, B: 34, IsDefault: false},
		"pine":       {R: 1, G: 121, B: 111, IsDefault: false},
		"olive":      {R: 128, G: 128, B: 0, IsDefault: false},
		"moss":       {R: 138, G: 154, B: 91, IsDefault: false},
		"sea-green":  {R: 46, G: 139, B: 87, IsDefault: false},
		"chartreuse": {R: 127, G: 255, B: 0, IsDefault: false},
		// ==========================================
		// 6. Циановые, бирюзовые и морские
		// ==========================================
		"cyan":      {R: 0, G: 255, B: 255, IsDefault: false},
		"aqua":      {R: 0, G: 255, B: 255, IsDefault: false},
		"teal":      {R: 20, G: 184, B: 166, IsDefault: false},
		"dark-teal": {R: 17, G: 94, B: 89, IsDefault: false},
		"turquoise": {R: 64, G: 224, B: 208, IsDefault: false},
		"ocean":     {R: 0, G: 105, B: 148, IsDefault: false},
		// ==========================================
		// 7. Синие и глубокие оттенки
		// ==========================================
		"blue":       {R: 0, G: 120, B: 255, IsDefault: false},
		"sky":        {R: 14, G: 165, B: 233, IsDefault: false},
		"sky-blue":   {R: 135, G: 206, B: 235, IsDefault: false},
		"azure":      {R: 0, G: 127, B: 255, IsDefault: false},
		"cerulean":   {R: 0, G: 123, B: 167, IsDefault: false},
		"cobalt":     {R: 0, G: 71, B: 171, IsDefault: false},
		"royal-blue": {R: 65, G: 105, B: 225, IsDefault: false},
		"sapphire":   {R: 15, G: 82, B: 186, IsDefault: false},
		"navy":       {R: 10, G: 25, B: 70, IsDefault: false},
		"midnight":   {R: 25, G: 25, B: 112, IsDefault: false},
		"steel-blue": {R: 70, G: 130, B: 180, IsDefault: false},
		"denim":      {R: 21, G: 96, B: 189, IsDefault: false},
		"ice-blue":   {R: 200, G: 230, B: 255, IsDefault: false},
		// ==========================================
		// 8. Фиолетовые, пурпурные и индиго
		// ==========================================
		"indigo":      {R: 99, G: 102, B: 241, IsDefault: false},
		"violet":      {R: 139, G: 92, B: 246, IsDefault: false},
		"purple":      {R: 168, G: 85, B: 247, IsDefault: false},
		"deep-purple": {R: 88, G: 28, B: 135, IsDefault: false},
		"amethyst":    {R: 153, G: 102, B: 204, IsDefault: false},
		"lavender":    {R: 230, G: 230, B: 250, IsDefault: false},
		"lilac":       {R: 200, G: 162, B: 200, IsDefault: false},
		"plum":        {R: 142, G: 69, B: 133, IsDefault: false},
		"eggplant":    {R: 97, G: 64, B: 81, IsDefault: false},
		"periwinkle":  {R: 204, G: 204, B: 255, IsDefault: false},
		// ==========================================
		// 9. Розовые, маджента и фуксия
		// ==========================================
		"pink":     {R: 236, G: 72, B: 153, IsDefault: false},
		"hot-pink": {R: 255, G: 105, B: 180, IsDefault: false},
		"rose":     {R: 244, G: 63, B: 94, IsDefault: false},
		"magenta":  {R: 255, G: 0, B: 255, IsDefault: false},
		"fuchsia":  {R: 217, G: 70, B: 239, IsDefault: false},
		"salmon":   {R: 250, G: 128, B: 114, IsDefault: false},
		"flamingo": {R: 252, G: 142, B: 172, IsDefault: false},
		"blush":    {R: 222, G: 93, B: 131, IsDefault: false},
		// ==========================================
		// 10. Неоновые (яркие акценты)
		// ==========================================
		"neon-cyan":   {R: 0, G: 235, B: 215, IsDefault: false},
		"neon-pink":   {R: 255, G: 60, B: 170, IsDefault: false},
		"neon-green":  {R: 50, G: 255, B: 120, IsDefault: false},
		"neon-yellow": {R: 255, G: 255, B: 51, IsDefault: false},
		"neon-purple": {R: 180, G: 4, B: 255, IsDefault: false},
		"neon-orange": {R: 255, G: 95, B: 31, IsDefault: false},
		"neon-blue":   {R: 31, G: 81, B: 255, IsDefault: false},
		// ==========================================
		// 11. Пастельные (мягкие оттенки)
		// ==========================================
		"pastel-purple": {R: 189, G: 147, B: 249, IsDefault: false},
		"pastel-cyan":   {R: 139, G: 233, B: 253, IsDefault: false},
		"pastel-green":  {R: 80, G: 250, B: 123, IsDefault: false},
		"pastel-pink":   {R: 255, G: 184, B: 212, IsDefault: false},
		"pastel-yellow": {R: 241, G: 250, B: 140, IsDefault: false},
		"pastel-orange": {R: 255, G: 184, B: 108, IsDefault: false},
		"pastel-blue":   {R: 174, G: 198, B: 207, IsDefault: false},
		"pastel-red":    {R: 255, G: 105, B: 97, IsDefault: false},
	},
}