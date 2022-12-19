package leetcode

/*
 * @lc app=leetcode.cn id=1703 lang=golang
 *
 * [1703] 得到连续 K 个 1 的最少相邻交换次数
 */

// @lc code=start
func minMoves(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	count := -1
	g, sum := []int{}, []int{0}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			g = append(g, i-count)
			sum = append(sum, sum[len(sum)-1]+g[len(g)-1])
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	result := 2 << 31
	for i := 0; i+k <= len(g); i++ {
		mid := (i + i + k - 1) / 2
		result = min(result, (2*(mid-i)-k+1)*g[mid]+(sum[i+k]-sum[mid+1])-(sum[mid]-sum[i]))
	}
	return result
}

// @lc code=end

/*
	参考官方题解，贪心加数学
*/
