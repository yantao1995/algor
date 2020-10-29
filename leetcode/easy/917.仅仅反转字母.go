package leetcode

/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
func reverseOnlyLetters(S string) string {
	prev, last := 0, len(S)-1
	byts := []byte(S)
	for prev < last {
		if byts[prev] < 65 || (byts[prev] > 90 && byts[prev] < 97) || byts[prev] > 122 {
			prev++
			continue
		}
		if byts[last] < 65 || (byts[last] > 90 && byts[last] < 97) || byts[last] > 122 {
			last--
			continue
		}
		byts[prev], byts[last] = byts[last], byts[prev]
		prev++
		last--
	}
	return string(byts)
}

// @lc code=end
