package leetcode

/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */

// @lc code=start
// type MyLinkedList struct {
// 	head, tail *link707         //头尾结点
// 	idx        map[int]*link707 //结点索引
// }
// type link707 struct {
// 	val  int
// 	next *link707
// }

// func Constructor() MyLinkedList {
// 	return MyLinkedList{
// 		head: nil,
// 		tail: nil,
// 		idx:  map[int]*link707{},
// 	}
// }

// func (this *MyLinkedList) Get(index int) int {
// 	if lk, ok := this.idx[index]; ok {
// 		return lk.val
// 	}
// 	return -1
// }

// func (this *MyLinkedList) AddAtHead(val int) {
// 	this.AddAtIndex(0, val)
// }

// func (this *MyLinkedList) AddAtTail(val int) {
// 	this.AddAtIndex(len(this.idx), val)
// }

// func (this *MyLinkedList) AddAtIndex(index int, val int) {
// 	if index > len(this.idx) {
// 		return
// 	}
// 	node := &link707{
// 		val:  val,
// 		next: nil,
// 	}
// 	if lk, ok := this.idx[index]; ok {
// 		if index <= 0 {
// 			index = 0
// 			if this.head != nil {
// 				node.next = this.head.next
// 			} else {
// 				this.tail = node
// 			}
// 			this.head = node
// 		} else {
// 			this.idx[index-1].next = node
// 			node.next = lk
// 		}
// 		for i := len(this.idx); i > index; i-- {
// 			this.idx[i] = this.idx[i-1]
// 		}
// 		this.idx[index] = node
// 	} else {
// 		if this.head == nil {
// 			this.head = node
// 			this.tail = node
// 		} else {
// 			this.tail.next = node
// 			this.tail = node
// 		}
// 		this.idx[len(this.idx)] = node
// 	}
// }

// func (this *MyLinkedList) DeleteAtIndex(index int) {
// 	if lk, ok := this.idx[index]; ok {
// 		if this.tail == lk {
// 			if index == 0 {
// 				this.tail = nil
// 				this.head = nil
// 			} else {
// 				this.tail = this.idx[index-1]
// 				this.tail.next = nil
// 			}
// 		} else {
// 			if index == 0 {
// 				this.head = this.head.next
// 			} else {
// 				this.idx[index-1].next = lk.next
// 			}
// 			_, ok2 := this.idx[index+1]
// 			for ok2 {
// 				this.idx[index] = this.idx[index+1]
// 				index++
// 				_, ok2 = this.idx[index+1]
// 			}
// 		}
// 		delete(this.idx, index)
// 	}
// }

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end

/*
	this 记录 head 和 tail 方便头插和尾插
	用map 记录每个节点的索引值，方便按索引查找
	添加和删除就只需要判断一下是头，尾还是中间节点
*/
