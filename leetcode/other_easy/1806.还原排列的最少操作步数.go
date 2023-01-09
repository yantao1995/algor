package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=1806 lang=golang
 *
 * [1806] 还原排列的最少操作步数
 */

// @lc code=start
func reinitializePermutation(n int) int {
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}
	equals := func(b []int) bool {
		for i := 1; i < len(b); i++ {
			if b[i]-b[i-1] != 1 {
				return false
			}
		}
		return true
	}
	new := make([]int, n)
	result := 0
	for !equals(perm) || result == 0 {
		new = make([]int, n)
		for i := 0; i < n; i++ {
			new[i] = perm[i/2]
			if i&1 == 1 {
				new[i] = perm[n/2+(i-1)/2]
			}
		}
		perm = new
		result++
		fmt.Println(perm)
	}
	return result
}

// @lc code=end

/*
直接模拟整个过程即可
*/
