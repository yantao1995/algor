package leetcode

/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// @lc code=start
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) || len(A) < 2 {
		return false
	}
	count := 0
	index1, index2 := 0, 0
	for k := range A {
		if A[k] != B[k] {
			count++
			if count == 1 {
				index1 = k
			} else {
				index2 = k
			}
		}
	}
	if count > 2 || count == 1 {
		return false
	}
	if count == 2 && (A[index1] != B[index2] || A[index2] != B[index1]) {
		return false
	}
	if count == 0 {
		ht := map[byte]bool{}
		for k := range A {
			if ht[A[k]] {
				return true
			}
			ht[A[k]] = true
		}
		return false
	}
	return true
}

// @lc code=end
