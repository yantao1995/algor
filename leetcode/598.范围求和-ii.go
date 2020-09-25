package leetcode

/*
 * @lc app=leetcode.cn id=598 lang=golang
 *
 * [598] 范围求和 II
 */

// @lc code=start
func maxCount(m int, n int, ops [][]int) int {
	data := [][]int{}
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		data = append(data, temp)
	}
	for k := range ops {
		for i := 0; i < ops[k][0]; i++ {
			for j := 0; j < ops[k][1]; j++ {
				data[i][j]++
			}
		}
	}
	ht := map[int]int{}
	for k := range data {
		for m := range data[k] {
			ht[data[k][m]]++
		}
	}
	max := 0
	for k := range ht {
		if k > max {
			max = k
		}
	}
	return ht[max]
}

// @lc code=end
