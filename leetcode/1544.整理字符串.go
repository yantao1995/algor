package leetcode

/*
 * @lc app=leetcode.cn id=1544 lang=golang
 *
 * [1544] 整理字符串
 */

// @lc code=start
func makeGood(s string) string {
	bts := []byte(s)
	for {
		for i := 0; i < len(bts)-1; i++ {
			if (bts[i] > 96 && bts[i+1] == bts[i]-32) ||
				(bts[i] < 96 && bts[i+1] == bts[i]+32) {
				if i+2 < len(bts) {
					bts = append(bts[:i], bts[i+2:]...)
				} else {
					bts = bts[:i]
				}
				goto come
			}
		}
		break
	come:
	}
	return string(bts)
}

// @lc code=end
