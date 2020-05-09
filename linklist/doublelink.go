package linklist

import (
	"fmt"
)

type Node struct {
	Data linkType
	Prev *Node
	Next *Node
}

type DoubleLink struct {
	Size int
	Head *Node
	Tail *Node
}

func InitDoubleLink() *DoubleLink {
	return &DoubleLink{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
}

//AppendOnTail 添加在末尾
func (dl *DoubleLink) AppendOnTail(data linkType) *DoubleLink {
	node := &Node{
		Data: data,
		Prev: nil,
		Next: nil,
	}
	if dl.Size == 0 {
		dl.Head = node
		dl.Tail = node
	} else {
		node.Prev = dl.Tail
		dl.Tail.Next = node
		dl.Tail = node
	}
	dl.Size++
	return dl
}

//AppendOnHead 添加在头节点
func (dl *DoubleLink) AppendOnHead(data linkType) *DoubleLink {
	node := &Node{
		Data: data,
		Prev: nil,
		Next: nil,
	}
	if dl.Size == 0 {
		dl.Head = node
		dl.Tail = node
	} else {
		dl.Head.Prev = node
		node.Next = dl.Head
		dl.Head = node
	}
	dl.Size++
	return dl
}

//AddNodeByIndex 按位置添加节点,在第几个结点之后添加 从0开始  ,0表示头部
func (dl *DoubleLink) AddNodeByIndex(index int, data linkType) *DoubleLink {
	newnode := &Node{
		Data: data,
		Prev: nil,
		Next: nil,
	}
	nodes := dl.Head
	if index > dl.Size || index < 0 {
		fmt.Println("索引越界")
		return dl
	}
	if index == 0 { //头处理
		dl.AppendOnHead(data)
	} else if index == dl.Size { //尾处理 当前遍历到最后一个
		dl.AppendOnTail(data)
	} else {
		i := 1
		for nodes != nil {
			if i == index {
				//中间节点
				newnode.Next = nodes.Next
				nodes.Next = newnode
				newnode.Prev = nodes.Next.Prev
				nodes.Next.Prev = newnode
				dl.Size++
				break
			}
			nodes = nodes.Next
		}
	}
	return dl
}

//DelNodeByIndex 按位置删除节点  从1开始
func (dl *DoubleLink) DelNodeByIndex(index int) *DoubleLink {
	nodes := dl.Head
	if index > dl.Size || index < 1 {
		fmt.Println("索引越界")
		return dl
	}
	for i := 0; i < dl.Size; i++ {
		if i+1 == index {
			if i == 0 { //头处理
				nodes = nodes.Next
				nodes.Prev = nil
				dl.Head = nodes
			} else if i == dl.Size-1 { //尾处理 当前遍历到最后一个
				nodes = nodes.Prev
				nodes.Next = nil
				dl.Tail = nodes
			} else { //中间节点
				nodes.Next.Prev = nodes.Prev
				nodes.Prev.Next = nodes.Next
			}
			dl.Size--
			break
		}
		nodes = nodes.Next
	}
	return dl
}

//PrintFromHead 头结点输出
func (dl *DoubleLink) PrintFromHead() {
	head := dl.Head
	for head != nil {
		fmt.Printf("%s ", head.Data)
		head = head.Next
	}
	fmt.Println("\ndouble link size = ", dl.Size)
}

//PrintFromTail 尾结点输出
func (dl *DoubleLink) PrintFromTail() {
	tail := dl.Tail
	for tail != nil {
		fmt.Printf("%s ", tail.Data)
		tail = tail.Prev
	}
	fmt.Println("\ndouble link size = ", dl.Size)
}

//TestDoubleLink 测试
func TestDoubleLink() {
	dbLink := InitDoubleLink()
	dbLink.AppendOnTail("1").AppendOnTail("2").AppendOnTail("3").AppendOnTail("4")
	dbLink.AppendOnHead("5").AppendOnHead("6")
	dbLink.PrintFromHead()
	dbLink.DelNodeByIndex(1)
	dbLink.PrintFromHead()
	dbLink.DelNodeByIndex(5)
	dbLink.PrintFromHead()
	dbLink.DelNodeByIndex(4)
	dbLink.PrintFromHead()
	dbLink.DelNodeByIndex(2)
	dbLink.PrintFromHead()
	dbLink.AddNodeByIndex(2, "A")
	dbLink.PrintFromHead()
	dbLink.AddNodeByIndex(0, "B")
	dbLink.PrintFromHead()
	dbLink.AddNodeByIndex(1, "c")
	dbLink.PrintFromHead()
	dbLink.AddNodeByIndex(1, "d")
	dbLink.PrintFromHead()
}
