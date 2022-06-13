package leetcode

/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */

// @lc code=start
func sortArrayByParity(A []int) []int {
	indexPrev, indexLast := 0, len(A)-1
	temp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			temp[indexPrev] = A[i]
			indexPrev++
		} else {
			temp[indexLast] = A[i]
			indexLast--
		}
	}
	return temp
}

// @lc code=end
