package design_mode

import "fmt"

/*
	API 为facade 模块的外观接口，大部分代码使用此接口简化对facade类的访问。
	facade模块同时暴露了a和b 两个Module 的NewXXX和interface，其它代码如果需要使用细节功能时可以直接调用。
*/

//a接口
type AModuleAPI interface {
	TestA() string
}

//b接口
type BModuleAPI interface {
	TestB() string
}

//a接口实现
type aModuleImpl struct{}

func (a *aModuleImpl) TestA() string {
	return "a module test"
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

//b接口实现
type bModuleImpl struct{}

func (b *bModuleImpl) TestB() string {
	return "b module test"
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

//对外api
type API interface {
	Test() string
}

//对外api实现
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func NewApi() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

func (api *apiImpl) Test() string {
	a := api.a.TestA()
	b := api.b.TestB()
	return fmt.Sprint(a, " - ", b)
}
