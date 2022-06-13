package leetcode

/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找常用字符
 */

// @lc code=start
func commonChars(A []string) []string {
	result := []string{}
	htAll := []map[byte]int{}
	for k := range A {
		ht := map[byte]int{}
		for m := range A[k] {
			ht[A[k][m]]++
		}
		htAll = append(htAll, ht)
	}
	for k := range htAll[0] {
		exists := true
		for i := 1; i < len(htAll); i++ {
			if val, ok := htAll[i][k]; !ok {
				exists = false
				break
			} else {
				if val < htAll[0][k] {
					htAll[0][k] = val
				}
			}
		}
		if exists && htAll[0][k] != 0 {
			for mm := 0; mm < htAll[0][k]; mm++ {
				result = append(result, string([]byte{k}))
			}
		}
	}

	return result
}

// @lc code=end
