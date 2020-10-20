package leetcode

/*
 * @lc app=leetcode.cn id=1160 lang=golang
 *
 * [1160] 拼写单词
 */

// @lc code=start
func countCharacters(words []string, chars string) int {
	ht, tempHt := map[byte]int{}, map[byte]int{}
	for k := range chars {
		ht[chars[k]]++
	}
	count := 0
	for _, v := range words {
		tempHt = map[byte]int{}
		for k := range v {
			tempHt[v[k]]++
		}
		for k := range tempHt {
			if tempHt[k] > ht[k] {
				goto come
			}
		}
		count += len(v)
	come:
	}
	return count
}

// @lc code=end
