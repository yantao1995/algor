package design_mode

import (
	"strconv"
	"strings"
)

/*
	解释器模式定义一套语言文法，并设计该语言解释器，使用户能使用特定文法控制解释器行为。
	解释器模式的意义在于，它分离多种复杂功能的实现，每个功能只需关注自身的解释。
	对于调用者不用关心内部的解释器的工作，只需要用简单的方式组合命令就可以。
*/

//定义结点接口
type Node interface {
	Interpret() int
}

//结点值对象
type ValNode struct {
	val int
}

func (n *ValNode) Interpret() int {
	return n.val
}

//结点加法运算
type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

//结点减法运算
type MinNode struct {
	left, right Node
}

func (n *MinNode) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

// 解释器对象
type Parser struct {
	exp   []string
	index int
	prev  Node
}

//获取当前结点
func (p *Parser) newValNode() Node {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{val: v}
}

//结点加法运算
func (p *Parser) newAddNode() Node {
	p.index++
	return &AddNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

//结点减法运算
func (p *Parser) newMinNode() Node {
	p.index++
	return &MinNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

//解释执行
func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")
	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddNode()
		case "-":
			p.prev = p.newMinNode()
		default:
			p.prev = p.newValNode()
		}
	}
}

func (p *Parser) Result() Node {
	return p.prev
}
