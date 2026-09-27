package widgets

import (
	"cli-graphics/engine"
	"image"

	"github.com/google/uuid"
)

type Image struct {
	engine.BaseComponent
	img image.Image
}

func NewImage(x, y, w, h int, img image.Image) *Image {
	return &Image{
		BaseComponent: engine.NewBaseComponent(uuid.New(), x, y, w, h),
		img:           img,
	}
}

func (i *Image) SetImage(img image.Image) {
	i.img = img
}

func (i *Image) OnTick() {}

func (imgWidget *Image) Render(canvas *engine.Canvas) {
	if imgWidget.img == nil {
		return
	}

	w, h := imgWidget.GetSize()
	bounds := imgWidget.img.Bounds()
	imgW := bounds.Dx()
	imgH := bounds.Dy()

	targetPixelW := w
	targetPixelH := h * 2

	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			topImgX := bounds.Min.X + (col * imgW / targetPixelW)
			topImgY := bounds.Min.Y + ((row * 2) * imgH / targetPixelH)
			topColor := engine.ToEngineColor(imgWidget.img.At(topImgX, topImgY))

			botImgX := bounds.Min.X + (col * imgW / targetPixelW)
			botImgY := bounds.Min.Y + ((row*2 + 1) * imgH / targetPixelH)
			botColor := engine.ToEngineColor(imgWidget.img.At(botImgX, botImgY))

			cell := engine.NewCellColored('▀', topColor, botColor)
			canvas.SetCell(col, row, cell)
		}
	}
}
