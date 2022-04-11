package leetcode

/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */

// @lc code=start
type MyLinkedList struct {
	head, tail *link707         //头尾结点
	idx        map[int]*link707 //结点索引
}
type link707 struct {
	val  int
	next *link707
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
		idx:  map[int]*link707{},
	}
}

func (this *MyLinkedList) Get(index int) int {
	if lk, ok := this.idx[index]; ok {
		return lk.val
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {

}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &link707{
		val:  val,
		next: nil,
	}
	this.tail.next = node
	this.tail = node
	this.idx[len(this.idx)] = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if lk, ok := this.idx[index]; ok {
		if index == 0 {
			this.head = this.head.next
		} else {
			this.idx[index-1].next = lk.next
			for _, ok2 := this.idx[index+1]; ok2; index++ {
				this.idx[index] = this.idx[index+1]
			}
			delete(this.idx, index)
		}
	}
}

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
