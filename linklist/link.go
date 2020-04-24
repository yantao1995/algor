package linklist

import (
	"fmt"
)

//Link one way
type Link struct {
	Next *Link
	Data string
}

//Init ...
func Init(data string) *Link {
	return &Link{
		Next: nil,
		Data: data,
	}
}

//AppendByIndexTail 按位置添加之后
func (l *Link) AppendByIndexTail(i int, data string) bool {
	index := 1
	flag := false
	for l != nil {
		if index == i {
			node := &Link{
				Next: l.Next,
				Data: data,
			}
			l.Next = node
			flag = true
		}
		l = l.Next
		index++
	}
	return flag
}

//AppendByDateEqualsTail 按数据是否相等添加在之后
func (l *Link) AppendByDateEqualsTail(origin, data string) bool {
	flag := false
	for l != nil {
		if l.Data == origin {
			node := &Link{
				Next: l.Next,
				Data: data,
			}
			l.Next = node
			l = l.Next //结点之后添加，防止重复添加
			flag = true
		}
		l = l.Next
	}
	return flag
}

//AppendOnHeadPoint 头 指针接收器
func (l *Link) AppendOnHeadPoint(data string) *Link {
	return &Link{
		Next: l,
		Data: data,
	}
}

//AppendOnHead 头
func AppendOnHead(l **Link, data string) {
	head := Link{
		Next: *l,
		Data: data,
	}
	*l = &head
}

//AppendOnTailPoint 尾 指针接收器
func (l *Link) AppendOnTailPoint(data string) *Link {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &Link{nil, data}
	return l
}

//AppendOnTail 尾
func AppendOnTail(l *Link, data string) {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &Link{nil, data}
}

//Print 输出
func (l *Link) Print() {
	for l != nil {
		fmt.Printf("%s ", l.Data)
		l = l.Next
	}
	fmt.Println()
}

//ReverseOrderByHeadInsert 逆序  头插法
func (l *Link) ReverseOrderByHeadInsert() *Link {
	tail := new(Link)
	for l != nil {
		tail = tail.AppendOnHeadPoint(l.Data)
		//AppendOnHead(&tail, l.Data)
		l = l.Next
	}
	return tail
}

//ReverseOrder 直接逆序 ----
func ReverseOrder(l *Link) *Link {
	prev, next := new(Link), new(Link)
	for l != nil {
		next = l.Next
		l.Next = prev
		prev = l
		l = next
	}
	return prev
}

//LinkTest 测试
func LinkTest() {
	head := Init("0")
	head.AppendOnTailPoint("1").AppendOnTailPoint("2") //指针接收  尾插
	AppendOnHead(&head, "3")                           //普通   头插
	AppendOnHead(&head, "4")
	head = head.AppendOnHeadPoint("5").AppendOnHeadPoint("6").AppendOnHeadPoint("7") //指针接收 头插
	AppendOnTail(head, "8")                                                          //普通 尾插
	AppendOnTail(head, "9")
	fmt.Println(head.AppendByIndexTail(10, "10"))
	head.Print()
	fmt.Println(head.AppendByDateEqualsTail("10", "11"))
	head.Print()
	head = head.ReverseOrderByHeadInsert() //指针接收 头插法逆序
	head.Print()
	head = ReverseOrder(head)
	head.Print()
}
