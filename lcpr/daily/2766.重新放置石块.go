package lcpr

import "sort"

/*
 * @lc app=leetcode.cn id=2766 lang=golang
 * @lcpr version=30204
 *
 * [2766] 重新放置石块
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	for k, v := range moveFrom {
		delete(m, v)
		m[moveTo[k]] = true
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

// @lc code=end

/*
// @lcpr case=start
// [1,6,7,8]\n[1,7,2]\n[2,9,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,3,3]\n[1,3]\n[2,2]\n
// @lcpr case=end

*/

/*
按步骤操作即可
*/
