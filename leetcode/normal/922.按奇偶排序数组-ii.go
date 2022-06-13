package leetcode

/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

// @lc code=start
func sortArrayByParityII(A []int) []int {
	odd, even := []int{}, []int{}
	for k := range A {
		if A[k]%2 == 1 {
			odd = append(odd, A[k])
		} else {
			even = append(even, A[k])
		}
	}
	oddIndex, evenIndex := 0, 0
	for k := range A {
		if k%2 == 1 {
			A[k] = odd[oddIndex]
			oddIndex++
		} else {
			A[k] = even[evenIndex]
			evenIndex++
		}
	}
	return A
}

// @lc code=end
