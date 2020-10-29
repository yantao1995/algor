package leetcode

/*
 * @lc app=leetcode.cn id=804 lang=golang
 *
 * [804] 唯一摩尔斯密码词
 */

// @lc code=start
func uniqueMorseRepresentations(words []string) int {
	ms := []string{".-", "-...", "-.-.", "-..", ".", "..-.",
		"--.", "....", "..", ".---", "-.-", ".-..", "--", "-.",
		"---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-",
		".--", "-..-", "-.--", "--.."}
	ht := map[string]bool{}
	for _, v1 := range words {
		tempStr := ""
		for _, v2 := range v1 {
			tempStr += ms[v2-'a']
		}
		ht[tempStr] = true
	}
	return len(ht)
}

// @lc code=end
