package leetcode

/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
type MyCircularDeque struct {
	data             []int
	headPtr, tailPtr int
	length           int
	isFirst          bool
}

// func Constructor(k int) MyCircularDeque {
// 	return MyCircularDeque{
// 		data:    make([]int, k),
// 		headPtr: 0,
// 		tailPtr: 0,
// 		length:  0,
// 		isFirst: true,
// 	}
// }

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.headPtr] = value
	this.length++
	this.headPtr--
	if this.headPtr == -1 {
		this.headPtr = len(this.data) - 1
	}
	if this.isFirst {
		this.tailPtr++
		this.isFirst = false
	}
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.tailPtr] = value
	this.length++
	this.tailPtr++
	if this.tailPtr == len(this.data) {
		this.tailPtr = 0
	}
	if this.isFirst {
		this.headPtr = len(this.data) - 1
		this.isFirst = false
	}
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.headPtr++
	this.length--
	if this.headPtr == len(this.data) {
		this.headPtr = 0
	}
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tailPtr--
	this.length--
	if this.tailPtr == -1 {
		this.tailPtr = len(this.data) - 1
	}
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	if this.headPtr+1 == len(this.data) {
		return this.data[0]
	}
	return this.data[this.headPtr+1]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	if this.tailPtr-1 == -1 {
		return this.data[len(this.data)-1]
	}
	return this.data[this.tailPtr-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.length == len(this.data)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

/*
	和设计循环队列差不多，需要注意第一个元素入栈时额外处理 headPtr 和 tailPtr
*/
