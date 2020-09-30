package leetcode

import "fmt"

//测试
func TestAlgor() {
	str := [][]int{
		{0, 0, 0},
		{0, 1, 0},
	}
	fmt.Println(floodFill(str, 1, 0, 2))
}
