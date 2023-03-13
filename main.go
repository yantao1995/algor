package main

import (
	"C"
	"image"
	"image/color"
	"os"

	"github.com/fogleman/gg"
)

func main() {

}

//export CutAndFill
func CutAndFill(inPath, outPath string, width, height int) {
	in, err := os.Open(inPath)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(in)
	if err != nil {
		panic(err)
	}

	dc := gg.NewContext(width, height)
	dc.SetColor(color.Transparent)
	dc.Clear()
	dc.DrawImage(img, 0, 0)
	dc.SavePNG(outPath)
}
