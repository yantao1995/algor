package design_mode

/*
	装饰模式使用对象组合的方式动态改变或增加对象行为。
	Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。
	使用匿名组合，在装饰器中不必显式定义转调原对象方法。
*/

//装饰器接口
type ComponentDecorator interface {
	Calc() int
}

//装饰器实现类
type ComponentDecoratorImpl struct{}

func (*ComponentDecoratorImpl) Calc() int {
	return 0
}

//乘法装饰器实现
type MulDecoratorImpl struct {
	ComponentDecorator
	num int
}

func NewMulDecoratorImpl(c ComponentDecorator, num int) ComponentDecorator {
	return &MulDecoratorImpl{c, num}
}

func (d *MulDecoratorImpl) Calc() int {
	return d.ComponentDecorator.Calc() * d.num
}

//加法装饰器实现
type AddDecoratorImpl struct {
	ComponentDecorator
	num int
}

func NewAddDecoratorImpl(c ComponentDecorator, num int) ComponentDecorator {
	return &AddDecoratorImpl{c, num}
}

func (d *AddDecoratorImpl) Calc() int {
	return d.ComponentDecorator.Calc() + d.num
}
