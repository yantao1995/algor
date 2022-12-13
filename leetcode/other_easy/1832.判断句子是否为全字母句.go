package leetcode

/*
 * @lc app=leetcode.cn id=1832 lang=golang
 *
 * [1832] 判断句子是否为全字母句
 */

// @lc code=start
func checkIfPangram(sentence string) bool {
	m := map[byte]struct{}{}
	for i := 0; i < len(sentence); i++ {
		m[sentence[i]] = struct{}{}
	}
	return len(m) == 26
}

// @lc code=end

/*
	直接用map数有多少个即可
*/
