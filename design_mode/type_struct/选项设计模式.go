package design_mode

type Option struct {
	id   int
	name string
}

type OptionFunc func(*Option)

//初始化1
func InitOptions(op *Option, of ...OptionFunc) {
	for _, v := range of {
		v(op)
	}
}

//初始化2
func NewOption(of ...OptionFunc) *Option {
	op := &Option{}
	for _, v := range of {
		v(op)
	}
	return op
}

func IdOpFunc(id int) OptionFunc {
	return func(ops *Option) {
		ops.id = id
	}
}

func NameOpFunc(name string) OptionFunc {
	return func(o *Option) {
		o.name = name
	}
}

/*
有时候一个函数会有很多参数，为了方便函数的使用，
会给希望给一些参数设定默认值，调用时只需要传与默认值不同的参数即可
*/
