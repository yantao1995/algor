package leetcode

/*
 * @lc app=leetcode.cn id=341 lang=golang
 *
 * [341] 扁平化嵌套列表迭代器
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	stack []*NestedInteger          //尾为栈顶
	top   func() *NestedInteger     //栈顶
	pop   func() *NestedInteger     //出栈
	push  func(elem *NestedInteger) //入栈
}

// func Constructor(nestedList []*NestedInteger) *NestedIterator {
// 	nest := &NestedIterator{
// 		stack: make([]*NestedInteger, 0, len(nestedList)),
// 	}
// 	nest.pop = func() *NestedInteger {
// 		elem := nest.stack[len(nest.stack)-1]
// 		nest.stack = nest.stack[:len(nest.stack)-1]
// 		return elem
// 	}
// 	nest.push = func(elem *NestedInteger) {
// 		nest.stack = append(nest.stack, elem)
// 	}
// 	nest.top = func() *NestedInteger {
// 		return nest.stack[len(nest.stack)-1]
// 	}
// 	for i := len(nestedList) - 1; i >= 0; i-- {
// 		nest.stack = append(nest.stack, nestedList[i])
// 	}
// 	return nest
// }

func (this *NestedIterator) Next() int {
	for len(this.stack) > 0 {
		elem := this.pop()
		if elem.IsInteger() {
			return elem.GetInteger()
		}
		list := elem.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			this.push(list[i])
		}
	}
	return 0
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		elem := this.top()
		if elem.IsInteger() {
			return true
		}
		this.pop()
		list := elem.GetList()
		if len(list) == 0 {
			continue
		}
		for i := len(list) - 1; i >= 0; i-- {
			this.push(list[i])
		}
	}
	return false
}

// @lc code=end

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return true
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

/*
	因为按顺序输出嵌套列表，所以可以使用栈来输出
	stack定义栈，pop,push,top分别定义栈操作
	因为存在 [[]] 这种空嵌套模式，所以在 hasNext 判断时，不能
	只判断stack是否有元素存在，因为空结构在调用Next时会出现返回0的情况，
	而正常情况下，是不应该返回0的，所以在hasNext时，如果当前不是Integer的话，
	就需要循环出栈入栈，直到栈顶是Integer为止
*/
