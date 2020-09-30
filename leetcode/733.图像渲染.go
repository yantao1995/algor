package leetcode

/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 {
		return image
	}
	iLength, jLength := len(image), len(image[0])
	value := image[sr][sc]
	dfs733(image, sr, sc, &newColor, &iLength, &jLength, value, 1, 1) //右
	dfs733(image, sr, sc, &newColor, &iLength, &jLength, value, 2, 1) //左
	dfs733(image, sr, sc, &newColor, &iLength, &jLength, value, 3, 1) //上
	dfs733(image, sr, sc, &newColor, &iLength, &jLength, value, 4, 1) //下
	return image
}

func dfs733(image [][]int, i, j int, newColor, iLength, jLength *int, value, direct, count int) {
	if i < *iLength && i >= 0 && j < *jLength && j >= 0 &&
		(image[i][j] == value || count == 1) {
		if image[i][j] == *newColor && count != 1 { //防止爆栈，已遍历过的点
			return
		}
		image[i][j] = *newColor
		if direct != 2 {
			dfs733(image, i+1, j, newColor, iLength, jLength, value, 1, count+1)
		}
		if direct != 1 {
			dfs733(image, i-1, j, newColor, iLength, jLength, value, 2, count+1)
		}
		if direct != 4 {
			dfs733(image, i, j+1, newColor, iLength, jLength, value, 3, count+1)
		}
		if direct != 3 {
			dfs733(image, i, j-1, newColor, iLength, jLength, value, 4, count+1)
		}
	}
}

// @lc code=end
