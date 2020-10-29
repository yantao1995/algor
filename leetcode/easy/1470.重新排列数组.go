package leetcode

/*
 * @lc app=leetcode.cn id=1470 lang=golang
 *
 * [1470] 重新排列数组
 */

// @lc code=start
func shuffle(nums []int, n int) []int {
	prev, next := []int{}, []int{}
	for k := 0; k < n; k++ {
		prev = append(prev, nums[k])
		next = append(next, nums[n+k])
	}
	index := 0
	indexP, indexN := 0, 0
	for index < len(nums) {
		nums[index] = prev[indexP]
		indexP++
		nums[index+1] = next[indexN]
		indexN++
		index += 2
	}
	return nums
}

// @lc code=end
