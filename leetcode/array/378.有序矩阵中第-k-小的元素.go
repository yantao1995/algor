package leetcode

/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第 K 小的元素
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	lastMin := 0
	for i := 1; i < k; i++ {
		matrix[lastMin] = matrix[lastMin][1:]
		if len(matrix[lastMin]) == 0 {
			for j := 0; j < len(matrix); j++ {
				if len(matrix[j]) > 0 {
					lastMin = j
					break
				}
			}
		}
		for j := 0; j < len(matrix); j++ {
			if len(matrix[j]) > 0 && matrix[j][0] < matrix[lastMin][0] {
				lastMin = j
			}
		}
	}
	return matrix[lastMin][0]
}

// @lc code=end

/*
	类似于归并排序
	每次都从第0列找当前最小的值，然后再删掉当前最小的元素，
	如果当前的行的元素已经为0了，那么就应该重新找出一个长度大于0的行来当上一轮的最小行
	这样下一轮的时候，第一列就又包含当前最小的元素了，
	然后重复上面的步骤,直到第k-1个元素(因为默认[0,0]坐标的元素为最小值，所以直接从第一个开始找)
*/
