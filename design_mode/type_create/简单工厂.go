package design_mode

/*
	go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
	NewXXX 函数返回接口时就是简单工厂模式，只有API 接口和NewAPI函数为包外可见，封装了实现细节。
*/
type SimpleFactoryApi interface {
	Say(name string) string
}

func NewSimpleFactoryApi(tp int) SimpleFactoryApi {
	if tp == 1 {
		return &object1{}
	} else if tp == 2 {
		return &object2{}
	}
	return nil
}

type object1 struct{}
type object2 struct{}

func (*object1) Say(name string) string {
	return "object1 Hello: " + name
}

func (*object2) Say(name string) string {
	return "object2 Hello: " + name
}
