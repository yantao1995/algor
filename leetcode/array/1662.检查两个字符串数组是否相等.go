package leetcode

/*
 * @lc app=leetcode.cn id=1662 lang=golang
 *
 * [1662] 检查两个字符串数组是否相等
 */

// @lc code=start
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	i, j, w1, w2 := 0, 0, 0, 0
	for w1 < len(word1) && w2 < len(word2) {
		if word1[w1][i] != word2[w2][j] {
			return false
		}
		i++
		if i == len(word1[w1]) {
			w1++
			i = 0
		}
		j++
		if j == len(word2[w2]) {
			w2++
			j = 0
		}
	}
	return w1 == len(word1) && w2 == len(word2) && i == 0 && j == 0
}

// @lc code=end

/*
	索引 w1 , w2 分别记录当前的 word1 和 word2 的索引
		i , j 分别记录 word1[w1] 和 word2[w2] 的索引

	如果不相等，直接返回false,
	相等则i,j索引都自增，如果到达单词边界，直接置0，并且跳转到下一个单词
	最后如果单词都到了末尾并且i,j都置0那说明相等
*/
