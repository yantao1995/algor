package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */

// @lc code=start
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	result := 1
	lastIndex := 0
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > pairs[lastIndex][1] {
			lastIndex = i
			result++
		}
	}

	return result
}

// @lc code=end

/*
	因为当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面
	假设描述为 (起点,终点)
	贪心，每次贪剩下的最小的终点
	所以可以以数对终点为标准进行升序排列，这样每次选取的数对，都能
	保证下一个选取的数对有最低的起点。数对链更紧凑，中间空隙越小，
	越能容纳更多的数对。
	记录上一个选取，然后遍历，遇到可以选取的就直接选进来即可
*/
