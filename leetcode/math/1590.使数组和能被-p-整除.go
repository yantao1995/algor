package leetcode

/*
 * @lc app=leetcode.cn id=1590 lang=golang
 *
 * [1590] 使数组和能被 P 整除
 */

// @lc code=start
func minSubarray(nums []int, p int) int {
	mod := 0
	for i := 0; i < len(nums); i++ {
		mod = (mod + nums[i]) % p
	}
	if mod == 0 {
		return 0
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	result := len(nums)
	cur := 0
	lastModMap := map[int]int{0: -1} //同余的最大索引位置
	for i := 0; i < len(nums); i++ {
		cur = (cur + nums[i]) % p
		target := (cur - mod + p) % p
		if idx, ok := lastModMap[target]; ok {
			result = min(result, i-idx)
		}
		lastModMap[cur] = i
	}
	if result == len(nums) {
		return -1
	}
	return result
}

// @lc code=end

/*

优化为O(n)参考官方的同余定理与公式转化



超时
Time Limit Exceeded
141/142 cases passed (N/A)

func minSubarray(nums []int, p int) int {
	sum := 0
	nums = append([]int{0}, nums...)
	memo := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		memo[i] = memo[i-1] + nums[i]
		sum += nums[i]
	}
	if sum%p == 0 {
		return 0
	}
	for length := 1; length < len(nums)-1; length++ {
		for j := 0; j+length < len(nums); j++ {
			if (sum-(memo[j+length]-memo[j]))%p == 0 {
				return length
			}
		}
	}
	return -1
}
*/
