package leetcode

/*
 * @lc app=leetcode.cn id=622 lang=golang
 *
 * [622] 设计循环队列
 */

// @lc code=start
type MyCircularQueue struct {
	queue            []int //内存空间
	headPtr, tailPtr int   //头尾指针
	length           int   //当前存储长度
}

// func Constructor(k int) MyCircularQueue {
// 	return MyCircularQueue{
// 		queue:   make([]int, k),
// 		headPtr: 0,
// 		tailPtr: -1,
// 		length:  0,
// 	}
// }

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.tailPtr++
	if this.tailPtr == len(this.queue) {
		this.tailPtr = 0
	}
	this.queue[this.tailPtr] = value
	this.length++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.headPtr++
	if this.headPtr == len(this.queue) {
		this.headPtr = 0
	}
	this.length--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.headPtr]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.tailPtr]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.length == len(this.queue)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end

/*
第一种：
	queue 保存数据
	headPtr,tailPtr 分别指向队列的头和尾
	length记录当前队列的长度
	headPtr 在出队列时后移
	tailPtr 先后移再入队，那么需要在队列为空时，先初始化为-1
			队列为空时，headPtr领先tailPtr 1个位置


第二种：

tailPtr 在入队列时后移，在获取队尾数据时
因为队尾是悬空的，所以需要获取前一个位置的数据
队列为空时，headPtr与tailPtr处于同一个位置

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue:   make([]int, k),
		headPtr: 0,
		tailPtr: 0, /的话可以获取队尾需要前移一位
		length:  0,
	}
}


// func (this *MyCircularQueue) EnQueue(value int) bool {
// 	if this.IsFull() {
// 		return false
// 	}
// 	this.queue[this.tailPtr] = value
// 	this.tailPtr++
// 	if this.tailPtr == len(this.queue) {
// 		this.tailPtr = 0
// 	}
// 	this.length++
// 	return true
// }

// func (this *MyCircularQueue) Rear() int {
// 	if this.IsEmpty() {
// 		return -1
// 	}
// 	if this.tailPtr == 0 {
// 		return this.queue[len(this.queue)-1]
// 	}
// 	return this.queue[this.tailPtr-1]
// }


*/
