package easy

/*
 * @lc app=leetcode.cn id=1417 lang=golang
 *
 * [1417] 重新格式化字符串
 */

// @lc code=start
func reformat(s string) string {
	nums, chars := []byte{}, []byte{}
	for k := range s {
		if s[k] >= '0' && s[k] <= '9' {
			nums = append(nums, s[k])
		} else {
			chars = append(chars, s[k])
		}
	}
	maxN := true
	if len(nums)-len(chars) > 1 || len(nums)-len(chars) < -1 {
		return ""
	}
	if len(chars) < len(nums) {
		maxN = false
	}
	bts := make([]byte, len(nums)+len(chars))
	indexN, indexC := 0, 0
	for k := range bts {
		if (k%2 == 0 && !maxN) || (k%2 == 1 && maxN) {
			bts[k] = nums[indexN]
			indexN++
		} else {
			bts[k] = chars[indexC]
			indexC++
		}
	}
	return string(bts)
}

// @lc code=end
