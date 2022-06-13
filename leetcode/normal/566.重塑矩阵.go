package leetcode

/*
 * @lc app=leetcode.cn id=566 lang=golang
 *
 * [566] 重塑矩阵
 */

// @lc code=start
func matrixReshape(nums [][]int, r int, c int) [][]int {
	checkLen := 0
	for k := range nums {
		checkLen += len(nums[k])
	}
	if checkLen < r*c {
		return nums
	}
	result := [][]int{}
	for i := 0; i < r; i++ {
		temp := make([]int, c)
		result = append(result, temp)
	}
	iIndex, jIndex := 0, 0
	for k := range nums {
		for m := range nums[k] {
			result[iIndex][jIndex] = nums[k][m]
			jIndex++
			if jIndex%c == 0 {
				iIndex++
				jIndex = 0
			}
		}
	}
	return result
}

// @lc code=end
