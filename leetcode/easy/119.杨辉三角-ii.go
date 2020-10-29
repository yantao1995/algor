package leetcode

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	rowIndex++
	row1 := []int{1}
	if rowIndex == 1 {
		return row1
	}
	temp := row1
	for i := 2; i <= rowIndex; i++ {
		newTemp := []int{1}
		for j := 0; j < len(temp)-1; j++ {
			newTemp = append(newTemp, temp[j]+temp[j+1])
		}
		newTemp = append(newTemp, 1)
		temp = newTemp
	}
	return temp
}

// @lc code=end
