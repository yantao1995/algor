package lcpr

/*
 * @lc app=leetcode.cn id=2951 lang=golang
 * @lcpr version=30202
 *
 * [2951] 找出峰值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findPeaks(mountain []int) []int {
	result := []int{}
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			result = append(result, i)
			i++
		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [2,4,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,4,3,8,5]\n
// @lcpr case=end

*/
