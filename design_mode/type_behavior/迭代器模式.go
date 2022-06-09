package design_mode

import "fmt"

/*
	送代器模式用于使用相同方式送代不同类型集合或者隐藏集合类型的具体实现。
	可以使用送代器模式使遍历同时应用送代策略，如请求新对象、过滤、处理对象等。
*/

//迭代器接口
type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

//目标对象的结构抽象索引
type Number struct {
	start, end int
}

func NewNumber(start, end int) *Number {
	return &Number{
		start: start,
		end:   end,
	}
}

func (n *Number) Iterator() Iterator {
	return &NumberIterator{
		number: n,
		next:   n.start,
	}
}

//迭代器实现类
type NumberIterator struct {
	number *Number
	next   int
}

func (i *NumberIterator) First() {
	i.next = i.number.start
}

func (i *NumberIterator) IsDone() bool {
	return i.next > i.number.end
}

func (i *NumberIterator) Next() interface{} {
	if !i.IsDone() {
		next := i.next
		i.next++
		return next
	}
	return nil
}

//迭代器聚合对象
type Aggregate interface {
	Iterator() Iterator
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		fmt.Println(i.Next())
	}
}
