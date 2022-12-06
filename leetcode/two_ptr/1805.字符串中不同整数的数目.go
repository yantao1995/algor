package leetcode

/*
 * @lc app=leetcode.cn id=1805 lang=golang
 *
 * [1805] 字符串中不同整数的数目
 */

// @lc code=start
func numDifferentIntegers(word string) int {
	m := map[string]struct{}{}
	for i, j := 0, 0; i < len(word) && j < len(word); {
		if word[i] >= '0' && word[i] <= '9' {
			for j = i; j < len(word) && word[j] >= '0' && word[j] <= '9'; {
				j++
			}
			for i != j && word[i] == '0' {
				i++
			}
			m[word[i:j]] = struct{}{}
			i = j
		} else {
			i++
		}
	}
	return len(m)
}

// @lc code=end

/*
	双指针，定位数字，然后清除数字开头的所有0，然后用map存储，返回个数即可
*/
