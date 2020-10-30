package leetcode

import "fmt"

//测试
func TestAlgor() {

	num := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	fmt.Println(topKFrequent(num, 10))

}

/////////////////////////////////////////////////////
type ListNode struct {
	Val  int
	Next *ListNode
}
