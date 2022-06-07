package design_mode

/*
  代理模式用于延迟处理操作或者在进行实际操作前后进行其它处理。
  代理模式的常见用法有
	虚代理
	COW代理
	远程代理
	保护代理
	Cache 代理
	防火墙代理
	同步代理
	智能指引
	 。。。
*/

//原对象实现方法
type Subject interface {
	Do() string
}

type SubJectImpl struct{}

func (SubJectImpl) Do() string {
	return "sub impl"
}

//代理对象
type Proxy struct {
	origin SubJectImpl
}

func (p Proxy) Do() (res string) {
	//代理前
	res += "proxy prev"
	//调用对象方法
	res += p.origin.Do()
	//代理后
	res += "proxy after"
	return
}
