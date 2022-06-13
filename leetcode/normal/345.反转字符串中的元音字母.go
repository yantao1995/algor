package leetcode

/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	byts := []byte(s)
	m := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	prev, tail := 0, len(byts)-1
	for prev < tail {
		if !m[byts[prev]] {
			prev++
			continue
		}
		if !m[byts[tail]] {
			tail--
			continue
		}
		byts[prev], byts[tail] = byts[tail], byts[prev]
		prev++
		tail--
	}
	return string(byts)
}

// @lc code=end
