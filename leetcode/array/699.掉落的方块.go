package leetcode

/*
 * @lc app=leetcode.cn id=699 lang=golang
 *
 * [699] 掉落的方块
 */

// @lc code=start
func fallingSquares(positions [][]int) []int {
	result := make([]int, len(positions))
	real := make([]int, len(positions))
	canOverlap := func(i, j int) bool {
		return !((positions[j][0] >= positions[i][0]+positions[i][1]) ||
			(positions[j][0]+positions[j][1] <= positions[i][0]))
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(positions); i++ {
		real[i] = positions[i][1]
		maxHeight := 0
		for j := i - 1; j >= 0; j-- {
			if canOverlap(i, j) {
				maxHeight = max(maxHeight, real[j])
			}
		}
		real[i] += maxHeight
		result[i] = real[i]
		if i > 0 && result[i-1] > real[i] {
			result[i] = result[i-1]
		}
	}
	return result
}

// @lc code=end

/*
	纯暴力解法
	real 记录当当前方块重叠的真实高度
	result 记录当前方块重叠后的整体最大高度
	canOverlap 记录当前方块是否可以重叠
	每一个方块都和前面的方块比较,如果可以重叠,则记录当前的高度，
	最后遍历完成后选择最高的高度作为当前的高度
*/
