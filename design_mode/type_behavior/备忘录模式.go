package design_mode

import "fmt"

/*
	备忘录模式用于保存程序内部状态到外部，又不希望暴露内部状态的情形。
	程序内部状态使用interface{}给外部进行存储，从而不暴露程序实现细节。
	备忘录模式同时可以离线保存内部状态，如保存到数据库，文件等。
*/

//存储对象
type Memento interface{}

//实现对象
type Game struct {
	hp, mp int
}

//内部对象
type gameMemento struct {
	hp, mp int
}

func (g *Game) Play(mpDelta, hpDelta int) {
	g.mp += mpDelta
	g.hp += hpDelta
}

func (g *Game) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *Game) Load(m Memento) {
	gm := m.(*gameMemento)
	g.hp = gm.hp
	g.mp = gm.mp
}

func (g *Game) Status() {
	fmt.Println("current hp:", g.hp, " mp:", g.mp)
}
