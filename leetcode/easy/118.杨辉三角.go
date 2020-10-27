package easy

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	datas := [][]int{}
	if numRows == 0 {
		return datas
	}
	row1 := []int{1}
	datas = append(datas, row1)
	if numRows == 1 {
		return datas
	}
	temp := row1
	for i := 2; i <= numRows; i++ {
		newTemp := []int{1}
		for j := 0; j < len(temp)-1; j++ {
			newTemp = append(newTemp, temp[j]+temp[j+1])
		}
		newTemp = append(newTemp, 1)
		temp = newTemp
		datas = append(datas, newTemp)
	}
	return datas
}

// @lc code=end
