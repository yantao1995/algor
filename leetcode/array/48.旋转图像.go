package leetcode

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	length := len(matrix)
	side := 0 //边界
	for count := length; count > 1; count -= 2 {
		tempSlice := make([]int, count)
		copy(tempSlice, matrix[side][0+side:length-side])          // 顶边复制转移
		tempSide := side                                           //临时边界
		for index := length - side - 1; index >= 0+side; index-- { // 上左
			matrix[side][index] = matrix[tempSide][side]
			tempSide++
		}
		tempSide = side
		for index := side; index < length-side; index++ { //左下
			matrix[tempSide][side] = matrix[length-side-1][index]
			tempSide++
		}
		for index := side; index < length-side; index++ { //下右
			matrix[length-side-1][index] = matrix[length-index-1][length-side-1]
		}
		tempSide = 0
		for index := side; index < length-side; index++ { //顶边复制到右
			matrix[index][length-side-1] = tempSlice[tempSide]
			tempSide++
		}
		side++
	}
}

// @lc code=end
