package leetcode

/*
 * @lc app=leetcode.cn id=1512 lang=golang
 *
 * [1512] 好数对的数目
 */

// @lc code=start
func numIdenticalPairs(nums []int) int {
	result := 0
	ht := map[int][]int{}
	for k := range nums {
		ht[nums[k]] = append(ht[nums[k]], k)
	}
	for k := range ht {
		temp := len(ht[k]) - 1
		for temp > 0 {
			result += temp
			temp--
		}
	}
	return result
}

// @lc code=end
