package easy

/*
 * @lc app=leetcode.cn id=1013 lang=golang
 *
 * [1013] 将数组分成和相等的三个部分
 */

// @lc code=start
func canThreePartsEqualSum(A []int) bool {
	avg, total := 0, 0
	for k := range A {
		total += A[k]
	}
	if total%3 != 0 {
		return false
	}
	avg = total / 3
	count := 0
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		if sum == avg {
			count++
			sum = 0
		}
		if count == 2 && i < len(A)-1 { //计算2次 还有多余的就直接返回
			return true
		}
	}
	return false
}

// @lc code=end
