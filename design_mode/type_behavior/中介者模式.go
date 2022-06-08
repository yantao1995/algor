package design_mode

import (
	"fmt"
	"strings"
)

/*
	中介者模式封装对象之间互交，使依赖变的简单，并且使复杂互交简单化，封装在中介者中。
	例子中的中介者使用单例模式生成中介者。
	中介者的change使用switch判断类型。
*/

//需要交互的对象
type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadData() {
	c.Data = "music,image"
	fmt.Println("CDDriver ReadData:", c.Data)
	GetMediatorInstance().change(c)
}

//需要交互的对象
type CPU struct {
	Video, Sound string
}

func (c *CPU) Process(data string) {
	slice := strings.Split(data, ",")
	c.Video = slice[0]
	c.Sound = slice[1]
	fmt.Println("CPU Process:", *c)
	GetMediatorInstance().change(c)
}

type Mediator struct {
	CD  *CDDriver
	CPU *CPU
}

//中介对象
var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

//中介实现
func (m *Mediator) change(i interface{}) {
	switch ty := i.(type) {
	case *CDDriver:
		m.CPU.Process(ty.Data)
	}
}
