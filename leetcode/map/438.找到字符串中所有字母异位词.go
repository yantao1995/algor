package leetcode

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	result := []int{}
	pMap := map[byte]int{}
	for k := range p {
		pMap[p[k]]++
	}
	iMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if i < len(p) {
			iMap[s[i]]++
			if i < len(p)-1 {
				continue
			}
		} else {
			iMap[s[i]]++
			iMap[s[i-len(p)]]--
		}
		for k := range pMap {
			if pMap[k] != iMap[k] {
				goto lab
			}
		}
		result = append(result, i-len(p)+1)
	lab:
	}
	return result
}

// @lc code=end
