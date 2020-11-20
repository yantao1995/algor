package leetcode

import (
	"fmt"
)

//测试
func TestAlgor() {

	fmt.Println()
	fmt.Println("______________")
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1))
}

/////////////////////////////////////////////////////
type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
