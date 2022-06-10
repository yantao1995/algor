package design_mode

import "fmt"

/*
	状态模式用于分离状态和行为。
*/

//链式调用接口
type Week interface {
	Today()
	Next(*DayContext)
}

// 链式调用接口实现
type DayContext struct {
	today Week
}

//初始化为周日
func NewDayContext() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

//简单实现  周日 -> 周一 -> 周日 ....
//实现一周得写7个

//周日实现
type Sunday struct{}

func (*Sunday) Today() {
	fmt.Println("Sunday")
}

func (*Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}

//周一实现
type Monday struct{}

func (*Monday) Today() {
	fmt.Println("Monday")
}

func (*Monday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}
