package design_mode

import "fmt"

/*
	定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符合开闭原则。
*/

//多态接口
type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

//实现对象
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

func NewPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardId,
			Money:  money,
		},
		strategy: strategy,
	}
}

//实现接口
func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

//cash实现
type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Println("pay by cash : ", ctx.Money, ctx.Name)
}

//bank实现
type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Println("pay :", ctx.Money, "to", ctx.Name, ctx.CardID)
}
