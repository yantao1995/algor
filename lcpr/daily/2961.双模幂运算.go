package lcpr

/*
 * @lc app=leetcode.cn id=2961 lang=golang
 * @lcpr version=30204
 *
 * [2961] 双模幂运算
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func getGoodIndices(variables [][]int, target int) []int {
	result := make([]int, 0, len(variables))
	var pow func(x, y, mod int) int
	pow = func(x, y, mod int) int {
		if y == 1 {
			return x % mod
		}
		return x * pow(x, y-1, mod) % mod
	}
	for k, v := range variables {
		if pow(pow(v[0], v[1], 10), v[2], v[3]) == target {
			result = append(result, k)
		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [[2,3,3,10],[3,3,3,1],[6,1,1,4]]\n2\n
// @lcpr case=end

// @lcpr case=start
// [[39,3,1000,1000]]\n17\n
// @lcpr case=end

*/

/*

指数运算需要注意是否越界，所以每步都取模即可

*/
