package leetcode

/*
 * @lc app=leetcode.cn id=961 lang=golang
 *
 * [961] 重复 N 次的元素
 */

// @lc code=start
func repeatedNTimes(A []int) int {
	ht := map[int]int{}
	for k := range A {
		ht[A[k]]++
	}
	for k := range ht {
		if ht[k] == len(A)/2 {
			return k
		}
	}
	return 0
}

// @lc code=end
