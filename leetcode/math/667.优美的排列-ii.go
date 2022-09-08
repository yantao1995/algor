package leetcode

/*
 * @lc app=leetcode.cn id=667 lang=golang
 *
 * [667] 优美的排列 II
 */

// @lc code=start
func constructArray(n int, k int) []int {
	result := make([]int, n)
	left, right, isLeft := 1, n, true
	for idx := 0; idx < len(result); idx++ {
		if isLeft {
			result[idx] = left
			left++
		} else {
			result[idx] = right
			right--
		}
		if k > 1 {
			isLeft = !isLeft
			k--
		}
	}
	return result
}

// @lc code=end

/*
	因为 k<n 所以总有满足要求的序列
	要构造 k 个不同的整数，那么可以先构造 k-1 个递减的整数
	然后 最后一个整数用 1来补
	比如n为10,即： 1 - 10
	可以先 [1,10,2,9...]
	这样可以保证整数序列是 [9,8,7...]
	k-1个序列构造出来之后，最后一个用1来补：
		如果当前是取的前面的数，那就从前面开始顺序取数字
		如果当前是取的后面的数，那就从后面开始顺序取数字
*/
