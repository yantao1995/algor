package leetcode

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
// type MyQueue struct {
// 	QueueSlice []int
// 	Head       int //队首
// 	Tail       int //队尾
// 	Size       int
// }

// /** Initialize your data structure here. */
// func Constructor() MyQueue {
// 	return MyQueue{
// 		QueueSlice: []int{},
// 		Head:       0,
// 		Tail:       0,
// 		Size:       0,
// 	}
// }

// /** Push element x to the back of queue. */
// func (q *MyQueue) Push(x int) {
// 	if q.Head == 0 {
// 		if q.Size == len(q.QueueSlice) {
// 			s := []int{x}
// 			s = append(s, q.QueueSlice...)
// 			q.QueueSlice = s
// 		} else { //右移
// 			for i := 0; i < q.Size; i++ {
// 				q.QueueSlice[q.Tail-i+1] = q.QueueSlice[q.Tail-i]
// 			}
// 		}
// 		if q.Size != 0 {
// 			q.Tail++
// 		}
// 	} else {
// 		if q.Size != 0 {
// 			q.Head--
// 		}
// 	}
// 	q.QueueSlice[q.Head] = x
// 	q.Size++
// }

// /** Removes the element from in front of queue and returns that element. */
// func (q *MyQueue) Pop() int {
// 	var rtnVal int
// 	if q.Size > 0 {
// 		//fmt.Printf("%v\t", q.QueueSlice[q.Tail])
// 		rtnVal = q.QueueSlice[q.Tail]
// 		q.QueueSlice[q.Tail] = 0
// 		if q.Head != q.Tail {
// 			q.Tail--
// 		}
// 		q.Size--
// 		return rtnVal
// 	}
// 	return 0
// }

// /** Get the front element. */
// func (q *MyQueue) Peek() int {
// 	var rtnVal int
// 	if q.Size > 0 {
// 		//fmt.Printf("%v\t", q.QueueSlice[q.Tail])
// 		rtnVal = q.QueueSlice[q.Tail]
// 		// q.QueueSlice[q.Tail] = 0
// 		// if q.Head != q.Tail {
// 		// 	q.Tail--
// 		// }
// 		// q.Size--
// 		return rtnVal
// 	}
// 	return 0
// }

// /** Returns whether the queue is empty. */
// func (q *MyQueue) Empty() bool {
// 	if q.Size > 0 {
// 		return false
// 	}
// 	return true
// }

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
