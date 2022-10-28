package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=862 lang=golang
 *
 * [862] 和至少为 K 的最短子数组
 */

// @lc code=start
func shortestSubarray(nums []int, k int) int {
	result := -1
	prefix := make([]int, len(nums))
	stack := make([]int, 0, len(nums))
	prefix[0] = nums[0]
	if nums[0] > 0 {
		stack = append(stack, 0)
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] >= k {
			return 1
		}
		prefix[i] = nums[i] + prefix[i-1]

	}
	fmt.Println(prefix)

	return result
}

// @lc code=end

/*


常规暴力解法，直接累加 超时
func shortestSubarray(nums []int, k int) int {
	result := -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= k {
				if result == -1 || j-i+1 < result {
					result = j - i + 1
				}
				break
			}
		}
	}
	return result
}
*/
