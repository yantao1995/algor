package lcpr

/*
 * @lc app=leetcode.cn id=661 lang=golang
 * @lcpr version=30204
 *
 * [661] 图片平滑器
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func imageSmoother(img [][]int) [][]int {
	result := make([][]int, len(img))
	for k := range result {
		result[k] = make([]int, len(img[k]))
	}
	getAvg := func(lenI, lenJ int, i, j int) int {
		count := 0 //至少为1
		total := 0
		for ii := i - 1; ii <= i+1; ii++ {
			for jj := j - 1; jj <= j+1; jj++ {
				if ii >= 0 && ii < lenI && jj >= 0 && jj < lenJ {
					total += img[ii][jj]
					count++
				}
			}
		}
		return total / count
	}
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ {
			result[i][j] = getAvg(len(img), len(img[i]), i, j)
		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,1],[1,0,1],[1,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[100,200,100],[200,50,200],[100,200,100]]\n
// @lcpr case=end

*/
