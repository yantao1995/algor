package design_mode

/*
	工厂方法模式使用子类的方式延迟生成对象到子类中实现。
	Go中不存在继承 所以使用匿名组合来实现
*/

type Operator interface {
	SetA(a int)
	SetB(b int)
	GetResult() int
}

type OperatorFactory interface {
	CreateOperator() Operator
}

//接口实现基类
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

//加法工厂
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) CreateOperator() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//具体实现类  加法
type PlusOperator struct {
	*OperatorBase
}

// GetResult implements Operator
func (p PlusOperator) GetResult() int {
	return p.a + p.b
}

//减法工厂
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) CreateOperator() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//具体实现类 减法
type MinusOperator struct {
	*OperatorBase
}

// GetResult implements Operator
func (m MinusOperator) GetResult() int {
	return m.a - m.b
}
