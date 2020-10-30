package leetcode

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	ht := map[int]int{}
	ht2 := map[int]bool{}
	for k := range nums {
		ht[nums[k]]++
	}
	for i := 0; i < len(nums)-1; i++ { //字典序
		for j := 0; j < len(nums)-1-i; j++ {
			if ht[nums[j]] > ht[nums[j+1]] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if !ht2[nums[i]] && len(ht2) < k {
			ht2[nums[i]] = true
		}
	}
	result := []int{}
	for m := range ht2 {
		result = append(result, m)
	}
	return result
}

// @lc code=end
