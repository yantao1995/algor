package leetcode

/*
 * @lc app=leetcode.cn id=862 lang=golang
 *
 * [862] 和至少为 K 的最短子数组
 */

// @lc code=start
func shortestSubarray(nums []int, k int) int {
	result := -1
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	queue := []int{}
	for i := 0; i <= len(nums); i++ {
		for len(queue) > 0 && prefix[i]-prefix[queue[0]] >= k {
			if result == -1 || i-queue[0] < result {
				result = i - queue[0]
			}
			queue = queue[1:]
		}
		for len(queue) > 0 && prefix[queue[len(queue)-1]] >= prefix[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	return result
}

// @lc code=end

/*

参考题解：https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/solutions/1926644/by-haonshi-r7zi/
前缀和 + 双端队列

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
