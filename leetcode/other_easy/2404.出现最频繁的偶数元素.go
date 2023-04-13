package leetcode

/*
 * @lc app=leetcode.cn id=2404 lang=golang
 *
 * [2404] 出现最频繁的偶数元素
 */

// @lc code=start
func mostFrequentEven(nums []int) int {
	val, count := -1, 0
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 0 {
			m[nums[i]]++
			if m[nums[i]] > count ||
				(m[nums[i]] == count && (val == -1 || val > nums[i])) {
				count = m[nums[i]]
				val = nums[i]
			}
		}
	}
	return val
}

// @lc code=end

/*
	用map记录个数，然后val和count分别记录当前的值和个数
*/
