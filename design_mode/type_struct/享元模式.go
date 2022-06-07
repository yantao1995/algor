package design_mode

import "fmt"

/*
	享元模式从对象中剥离出不发生改变且多个实例需要的重复数据，
	独立出一个享元，使多个对象共享，从而节省内存以及减少对象数量。
*/

//享元对象
type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(data string) *ImageFlyweight {
	return &ImageFlyweight{data}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

//享元工厂
type ImageFlyweightFactory struct {
	m map[string]*ImageFlyweight
}

var imageFactory *ImageFlyweightFactory

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{make(map[string]*ImageFlyweight)}
	}
	return imageFactory
}

func (f *ImageFlyweightFactory) Get(data string) *ImageFlyweight {
	image := f.m[data]
	if image == nil {
		image = NewImageFlyweight(data)
		f.m[data] = image
	}
	return image
}

//需要享元的对象
type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(data string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(data)
	return &ImageViewer{image}
}

func (i *ImageViewer) Display() {
	fmt.Println("display: ", i.Data())
}
