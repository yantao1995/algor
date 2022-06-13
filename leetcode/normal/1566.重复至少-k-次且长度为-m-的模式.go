package leetcode

/*
 * @lc app=leetcode.cn id=1566 lang=golang
 *
 * [1566] 重复至少 K 次且长度为 M 的模式
 */

// @lc code=start
func containsPattern(arr []int, m int, k int) bool {
	for i := range arr {
		count := 1
		for j := i + m; j+m <= len(arr); j += m {
			temp := m - 1
			for temp >= 0 {
				if arr[j+temp] != arr[i+temp] {
					goto come
				}
				temp--
			}
			count++
		}
	come:
		if count >= k {
			return true
		}
	}
	return false
}

// @lc code=end
