package leetcode

/*
 * @lc app=leetcode.cn id=1754 lang=golang
 *
 * [1754] 构造字典序最大的合并字符串
 */

// @lc code=start
func largestMerge(word1 string, word2 string) string {
	check := func(i, j int) bool {
		for i < len(word1) && j < len(word2) {
			if word1[i] > word2[j] {
				return true
			} else if word1[i] < word2[j] {
				return false
			}
			i++
			j++
		}
		return i < len(word1)
	}
	result := make([]byte, 0, len(word1)+len(word2))
	for i, j := 0, 0; i < len(word1) || j < len(word2); {
		if j >= len(word2) ||
			(i < len(word1) && j < len(word2) &&
				(word1[i] > word2[j] || (word1[i] == word2[j] && check(i, j)))) {
			result = append(result, word1[i])
			i++
		} else {
			result = append(result, word2[j])
			j++
		}
	}
	return string(result)
}

// @lc code=end

/*
	word1和word2只要有一个取完了，就直接取另一个
	如果都没取完：
		word1[i] > word2[j] ，那么就取第一个
		word1[i] < word2[j] , 取第二个
		word1[i] == word2[j] ，那么依次向后比较，看那个大，直接取大的那个
		，比较完都没有就取剩下的还没取完的，如果两个都取完了，就取任意个
*/
