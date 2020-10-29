package leetcode

/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	prev, tail := 0, 0
	nums := []byte(s)
	if len(s) < k {
		prev, tail = 0, len(s)-1
		for prev < tail {
			nums[prev], nums[tail] = nums[tail], nums[prev]
			prev++
			tail--
		}
		return string(nums)
	}

	for i := 0; i < len(nums); i++ {
		if i%k == 0 && (i/k)%2 == 0 {
			prev, tail = i, i+k-1
			if i+k-1 >= len(nums) {
				tail = len(nums) - 1
			}
			for prev < tail {
				nums[prev], nums[tail] = nums[tail], nums[prev]
				prev++
				tail--
			}
		}
	}
	return string(nums)
}

// @lc code=end
