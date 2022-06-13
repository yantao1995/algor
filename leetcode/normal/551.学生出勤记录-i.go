package leetcode

/*
 * @lc app=leetcode.cn id=551 lang=golang
 *
 * [551] 学生出勤记录 I
 */

// @lc code=start
func checkRecord(s string) bool {
	ht := map[string]int{}
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			ht["A"]++
		}
		if i+2 < len(s) {
			if s[i:i+3] == "LLL" {
				ht["LLL"]++
			}
		}
		if ht["A"] == 2 || ht["LLL"] == 1 {
			return false
		}
	}
	return true
}

// @lc code=end
