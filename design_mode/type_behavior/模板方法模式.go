package design_mode

import "fmt"

/*
	模版方法模式使用继承机制，把通用步骤和通用方法放到父类中，把具体实现延迟到子类中实现。使得实现符合开闭原则。
	如实例代码中通用步骤在父类中实现（准备、下载、保存、收尾）下载和保存的具体实现留到子类中，并且提供 保存方法的默认实现。
	因为父类需要调用子类方法，所以子类需要匿名组合父类的同时，父类需要持有子类的引用。
*/

//下载工具接口
type Downloader interface {
	Download(uri string)
}

//下载工具模板
type template struct {
	implement
	uri string
}

// 下载需要实现的接口
type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Println("pre  download")
	t.implement.download()
	t.implement.save()
	fmt.Println("finish download")
}

func (t *template) save() {
	fmt.Println("parent save")
}

//http实现类
type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Println("http  download")
}

func (d *HTTPDownloader) save() {
	fmt.Println("http  save")
}

//  ftp 实现类
type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Println("ftp  download")
}
