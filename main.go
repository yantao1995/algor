package main

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// // open "test.jpg"
	// file, err := os.Open("test.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// os.Remove("test_resized.jpg")
	// out, err := os.Create("test_resized.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer out.Close()

	var ic ImageControl = ImageControl{}
	bts, err := os.ReadFile("test.jpg")
	if err != nil {
		panic(err)
	}
	src, _, err := image.Decode(bytes.NewReader(bts))
	if err != nil {
		panic(err)
	}
	tag, err := ic.ImageCopy(src, 0, 0, 200, 200)
	if err != nil {
		panic(err)
	}
	ic.SaveImage("test_resized.jpg", tag)

}

func ImageCut() {

}

type ImageControl struct {
}

func (this *ImageControl) Trimming(before_filename string, after_filename string, x, y, w, h int) {
	src, err := this.LoadImage(before_filename)
	if err != nil {
		log.Println("load image fail..")
	}

	img, err := this.ImageCopy(src, x, y, w, h)
	if err != nil {
		log.Println("image copy fail...")
	}
	saveErr := this.SaveImage(after_filename, img)
	if saveErr != nil {
		log.Println("save image fail..")
	}
}

func (this *ImageControl) LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

func (this *ImageControl) SaveImage(p string, src image.Image) error {
	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return err
	}
	defer f.Close()
	ext := filepath.Ext(p)

	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {

		err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})

	} else if strings.EqualFold(ext, ".png") {
		err = png.Encode(f, src)
	} else if strings.EqualFold(ext, ".gif") {
		err = gif.Encode(f, src, &gif.Options{NumColors: 256})
	}
	return err
}

func (this *ImageControl) ImageCopy(src image.Image, x, y, w, h int) (image.Image, error) {

	var subImg image.Image

	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	} else {

		return subImg, errors.New("图片解码失败")
	}

	return subImg, nil
}
func (this *ImageControl) GetImageBytes(p string, src image.Image) error {
	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	ext := filepath.Ext(p)

	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {

		err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})

	} else if strings.EqualFold(ext, ".png") {
		err = png.Encode(f, src)
	} else if strings.EqualFold(ext, ".gif") {
		err = gif.Encode(f, src, &gif.Options{NumColors: 256})
	}
	return err
}
