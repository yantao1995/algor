package linklist

import (
	"fmt"
)

type Node struct {
	Data string
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
func (dl *DoubleLink) AppendOnTail(data string) *DoubleLink {
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
func (dl *DoubleLink) AppendOnHead(data string) *DoubleLink {
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

//DelNodeByIndex 按位置删除节点
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

func (dl *DoubleLink) Print() {
	head := dl.Head
	for head != nil {
		fmt.Printf("%s ", head.Data)
		head = head.Next
	}
	fmt.Println("\ndouble link size = ", dl.Size)
}

//TestDoubleLink 测试
func TestDoubleLink() {
	dbLink := InitDoubleLink()
	dbLink.AppendOnTail("1").AppendOnTail("2").AppendOnTail("3").AppendOnTail("4")
	dbLink.AppendOnHead("5").AppendOnHead("6")
	dbLink.Print()
	dbLink.DelNodeByIndex(1)
	dbLink.Print()
	dbLink.DelNodeByIndex(5)
	dbLink.Print()
	dbLink.DelNodeByIndex(4)
	dbLink.Print()
	dbLink.DelNodeByIndex(2)
	dbLink.Print()
}
