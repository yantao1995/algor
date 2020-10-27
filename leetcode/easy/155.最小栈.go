package easy

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
// type MinStack struct {
// 	StackSlice []int
// 	Size       int
// 	TopPointer int   //栈顶
// 	MinSlice   []int //最小元素
// }

// /** initialize your data structure here. */
// func Constructor() MinStack {
// 	return MinStack{
// 		StackSlice: []int{},
// 		Size:       0,
// 		TopPointer: 0,
// 		MinSlice:   []int{},
// 	}
// }

// func (this *MinStack) Push(x int) {
// 	if this.Size > 0 {
// 		this.TopPointer++
// 	}
// 	if this.Size < len(this.StackSlice) { //复用
// 		this.StackSlice[this.TopPointer] = x
// 	} else {
// 		this.StackSlice = append(this.StackSlice, x)
// 	}
// 	this.Size++
// 	//最小栈判断入栈
// 	if len(this.MinSlice) > 0 {
// 		if x <= this.MinSlice[len(this.MinSlice)-1] {
// 			this.MinSlice = append(this.MinSlice, x)
// 		}
// 	} else {
// 		this.MinSlice = append(this.MinSlice, x)
// 	}
// }

// func (this *MinStack) Pop() {
// 	temp := 999
// 	if this.Size > 0 {
// 		//fmt.Printf("%v\t", this.StackSlice[this.TopPointer])
// 		temp = this.StackSlice[this.TopPointer]
// 		this.StackSlice[this.TopPointer] = 0
// 		if this.TopPointer != 0 {
// 			this.TopPointer--
// 		}
// 		this.Size--
// 	}
// 	//最小栈判断出栈
// 	if temp == this.MinSlice[len(this.MinSlice)-1] {
// 		this.MinSlice = this.MinSlice[:len(this.MinSlice)-1]
// 	}
// }

// func (this *MinStack) Top() int {
// 	if this.Size > 0 {
// 		return this.StackSlice[this.TopPointer]
// 	}
// 	return 0
// }

// func (this *MinStack) GetMin() int {
// 	if len(this.MinSlice) > 0 {
// 		return this.MinSlice[len(this.MinSlice)-1]
// 	}
// 	return 0
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.TopPointer();
 * param_4 := obj.GetMin();
 */
// @lc code=end
