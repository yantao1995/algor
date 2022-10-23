package leetcode

/*
 * @lc app=leetcode.cn id=1768 lang=golang
 *
 * [1768] 交替合并字符串
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	bts := make([]byte, 0, len(word1)+len(word2))
	for i, j := 0, 0; i < len(word1) || j < len(word2); i, j = i+1, j+1 {
		if i < len(word1) {
			bts = append(bts, word1[i])
		}
		if j < len(word2) {
			bts = append(bts, word2[j])
		}
	}
	return string(bts)
}

// @lc code=end

/*
	直接模拟整个操作即可
*/
