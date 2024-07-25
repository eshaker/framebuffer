// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package framebuffer

import (
	"image"
	"image/color"
)

type BGRA struct {
	Pix    []byte
	Rect   image.Rectangle
	Stride int
}

func (i *BGRA) Bounds() image.Rectangle { return i.Rect }
func (i *BGRA) ColorModel() color.Model { return color.RGBAModel }

func (i *BGRA) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(i.Rect)) {
		return color.RGBA{}
	}

	n := i.PixOffset(x, y)
	//	pix := i.Pix[n:]
	return color.RGBA{i.Pix[n+2], i.Pix[n+1], i.Pix[n+0], i.Pix[n+3]}
}

func (i *BGRA) Set(x, y int, c color.Color) {
	i.SetRGBA(x, y, color.RGBAModel.Convert(c).(color.RGBA))
}

func (i *BGRA) SetRGBA(x, y int, c color.RGBA) {
	if !(image.Point{x, y}.In(i.Rect)) {
		return
	}

	n := i.PixOffset(x, y)
	//pix := i.Pix[n:]

	i.Pix[n+0] = c.B
	i.Pix[n+1] = c.G
	i.Pix[n+2] = c.R
	i.Pix[n+3] = c.A
}

func (i *BGRA) PixOffset(x, y int) int {
	return (y-i.Rect.Min.Y)*i.Stride + (x-i.Rect.Min.X)*4
}
