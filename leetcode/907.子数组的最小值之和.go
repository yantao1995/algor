package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=907 lang=golang
 *
 * [907] 子数组的最小值之和
 */

// @lc code=start
func sumSubarrayMins(arr []int) int {
	total := 0
	m := map[[2]int]int{}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var binarySearch func(i, j int) int
	binarySearch = func(i, j int) int {
		if m[[2]int{i, j}] == 0 {
			if i == j {
				m[[2]int{i, j}] = arr[i]
			} else if i+1 == j {
				m[[2]int{i, j}] = min(binarySearch(i, i), binarySearch(j, j))
			} else {
				for mid := i + 1; mid < j; mid++ {
					m[[2]int{i, j}] = min(binarySearch(i, mid), binarySearch(mid, j))
				}
			}
		}
		return m[[2]int{i, j}]
	}
	m[[2]int{0, len(arr) - 1}] = binarySearch(0, len(arr)-1)
	fmt.Println(m)
	mod := int(1e9 + 7)
	for k := range m {
		total += m[k]
		total %= mod
	}
	return total
}

// @lc code=end

/*


暴力模拟 超时
func sumSubarrayMins(arr []int) int {
	getMin := func(i, j int) int {
		result := arr[i]
		for i++; i <= j; i++ {
			if arr[i] < result {
				result = arr[i]
			}
		}
		return result
	}
	mod := int(1e9 + 7)
	total := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			total += getMin(i, j)
			total %= mod
		}
	}
	return total
}
*/
