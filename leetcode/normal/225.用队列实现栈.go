package leetcode

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
// type MyStack struct {
// 	StackSlice []int
// 	Size       int
// 	TopPointer int //栈顶
// }

// /** Initialize your data structure here. */
// func Constructor() MyStack {
// 	return MyStack{
// 		StackSlice: []int{},
// 		Size:       0,
// 		TopPointer: 0,
// 	}
// }

// /** Push element x onto stack. */
// func (s *MyStack) Push(x int) {
// 	if s.Size > 0 {
// 		s.TopPointer++
// 	}
// 	if s.Size < len(s.StackSlice) { //复用
// 		s.StackSlice[s.TopPointer] = x
// 	} else {
// 		s.StackSlice = append(s.StackSlice, x)
// 	}
// 	s.Size++
// }

// /** Removes the element on top of the stack and returns that element. */
// func (s *MyStack) Pop() int {
// 	var rtnVal int
// 	if s.Size > 0 {
// 		//fmt.Printf("%v\t", s.StackSlice[s.TopPointer])
// 		rtnVal = s.StackSlice[s.TopPointer]
// 		s.StackSlice[s.TopPointer] = 0
// 		if s.TopPointer != 0 {
// 			s.TopPointer--
// 		}
// 		s.Size--
// 		return rtnVal
// 	}
// 	return 0
// }

// /** Get the top element. */
// func (s *MyStack) Top() int {
// 	if s.Size > 0 {
// 		return s.StackSlice[s.TopPointer]
// 	}
// 	return 0
// }

// /** Returns whether the stack is empty. */
// func (s *MyStack) Empty() bool {
// 	if s.Size > 0 {
// 		return false
// 	}
// 	return true
// }

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.TopPointer();
 * param_4 := obj.Empty();
 */
// @lc code=end
