package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"image"
	"image/color"

	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/fogleman/gg"
)

func main() {
	// str := `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`
	// result := CutAndFill(str, 500, 500)
	// fmt.Println(result)
}

/*
	图片剪裁和填充,传入base64,宽，高，返回指定大小的图片的base64
*/
//export CutAndFill
func CutAndFill(in *C.char, width, height int) *C.char {
	inBase64 := C.GoString(in)
	bts, err := base64.StdEncoding.DecodeString(inBase64)
	if err != nil {
		panic("base64 Decode image:" + err.Error())
	}
	inBuf := bytes.NewBuffer(bts)
	img, _, err := image.Decode(inBuf)
	if err != nil {
		panic("image Decode image:" + err.Error())
	}
	dc := gg.NewContext(width, height)
	dc.SetColor(color.Transparent)
	dc.Clear()
	dc.DrawImage(img, 0, 0)
	bf := &bytes.Buffer{}
	err = png.Encode(bf, dc.Image())
	if err != nil {
		panic("image Encode image:" + err.Error())
	}
	out := base64.StdEncoding.EncodeToString(bf.Bytes())
	return C.CString(out)
}
