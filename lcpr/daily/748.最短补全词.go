package lcpr

import "strings"

/*
 * @lc app=leetcode.cn id=748 lang=golang
 * @lcpr version=30204
 *
 * [748] 最短补全词
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func shortestCompletingWord(licensePlate string, words []string) string {
	source := [26]int{}
	licensePlate = strings.ToLower(licensePlate)
	for i := range licensePlate {
		if licensePlate[i] >= 'a' && licensePlate[i] <= 'z' {
			source[licensePlate[i]-'a']++
		}
	}
	sourceIndex := []int{}
	for idx, vv := range source {
		if vv > 0 {
			sourceIndex = append(sourceIndex, idx)
		}
	}
	countByte := [26]int{}
	result, maxCount := "", 0
	for _, w := range words {
		for _, v := range w {
			countByte[v-'a']++
		}
		count := 0
		for _, k := range sourceIndex {
			if source[k] > 0 {
				if source[k] <= countByte[k] {
					count++
				}
			}
		}
		if count > maxCount {
			result = w
			maxCount = count
		}
		if count == maxCount {
			if len(w) < len(result) {
				result = w
			}
		}
		for i := range countByte {
			countByte[i] = 0
		}
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// "1s3 PSt"\n["step", "steps", "stripe", "stepple"]\n
// @lcpr case=end

// @lcpr case=start
// "1s3 456"\n["looks", "pest", "stew", "show"]\n
// @lcpr case=end

*/

/*
	26长度的[26]byte记录每个字节出现的次数
	sourceIndex记录索引，这样不用每次都遍历26个字节
	然后对应数个数即可
*/
