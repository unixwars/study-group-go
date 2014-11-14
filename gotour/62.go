package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}


func (im Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (im Image) At(x, y int) color.Color {
    v1 := uint8(x+y)
    v2 := uint8(x-y)
    return color.RGBA{v1, v2, 255, 255}   
}


func main() {
    m := Image{}
    pic.ShowImage(m)
}
