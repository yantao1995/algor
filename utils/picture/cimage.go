package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"image"
	"image/color"
	"image/png"

	"bytes"
	"encoding/base64"
	"encoding/binary"

	"github.com/fogleman/gg"
)

func main() {
	str := `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`
	result := CutAndFillYuv(C.CString(str), 102, 126)
	fmt.Println(C.GoString(result))
}

/*
	图片剪裁和填充,传入base64,宽，高，返回指定大小的YUV图片的base64
*/
//export CutAndFillYuv
func CutAndFillYuv(in *C.char, width, height int) *C.char {

	in2 := CutAndFill(in, width, height)
	inBase64 := C.GoString(in2)
	bts, err := base64.StdEncoding.DecodeString(inBase64)
	if err != nil {
		panic("base64 Decode image:" + err.Error())
	}
	inBuf := bytes.NewBuffer(bts)
	img, _, err := image.Decode(inBuf)
	if err != nil {
		panic("image Decode image:" + err.Error())
	}
	//
	// of, err := os.Create("out.yuv")
	// if err != nil {
	// 	panic(err)
	// }
	// defer of.Close()
	// w := bufio.NewWriter(of)
	// image 转换成 yuv
	buf := new(bytes.Buffer)
	yuv := make([]uint8, 4, 4)
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x += 2 {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := img.At(x+1, y).RGBA()
			y1, u, v := color.RGBToYCbCr(uint8(r1), uint8(g1), uint8(b1))
			y2, _, _ := color.RGBToYCbCr(uint8(r2), uint8(g2), uint8(b2))
			yuv[0] = u
			yuv[1] = y1
			yuv[2] = v
			yuv[3] = y2
			err = binary.Write(buf, binary.LittleEndian, yuv)
			if err != nil {
				panic(err)
			}
		}
	}
	//
	// _, err = w.Write(buf.Bytes())
	// if err != nil {
	// 	panic(err)
	// }
	out := base64.StdEncoding.EncodeToString(buf.Bytes())
	return C.CString(out)
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
