package design_mode

import "fmt"

/*
	命令模式本质是把某个对象的方法调用封装到对象中，方便传递、存储、调用。
	示例中把主板单中的启动(start)方法和重启(reboot)方法封装为命令对象，再传递到主机(box)对象中。于两个按钮进行绑定：
	第一个机箱(box1)设置按钮1(button1) 为开机按钮2(button2)为重启。
	第二个机箱(box1)设置按钮2(button2) 为开机按钮1(button1)为重启。

*/

//命令执行实现类
type MainBoard struct{}

//开机方法
func (*MainBoard) Start() {
	fmt.Println("main board start.")
}

//重启方法
func (*MainBoard) Reboot() {
	fmt.Println("main board reboot.")
}

//按钮实现类
type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}

//命令执行接口
type Command interface {
	Execute()
}

//开机命令执行封装类
type StartCommand struct {
	mb *MainBoard
}

func NewStartCommand(mb *MainBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

//吃奶粉钱命令执行封装类
type RebootCommand struct {
	mb *MainBoard
}

func NewRebootCommand(mb *MainBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}
