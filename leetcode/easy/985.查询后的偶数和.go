package leetcode

/*
 * @lc app=leetcode.cn id=985 lang=golang
 *
 * [985] 查询后的偶数和
 */

// @lc code=start
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	total := 0
	result := []int{}
	for k := range A {
		if A[k]%2 == 0 {
			total += A[k]
		}
	}
	for k := range queries {
		if A[queries[k][1]]%2 == 0 {
			total -= A[queries[k][1]]
		}
		A[queries[k][1]] += queries[k][0]
		if A[queries[k][1]]%2 == 0 {
			total += A[queries[k][1]]
		}
		result = append(result, total)
	}
	return result
}

// @lc code=end
