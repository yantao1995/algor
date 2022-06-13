package leetcode

/*
 * @lc app=leetcode.cn id=888 lang=golang
 *
 * [888] 公平的糖果交换
 */

// @lc code=start
// B给的 = x给的 + (sum(A)+sum(B))/2
func fairCandySwap(A []int, B []int) []int {
	totalA, totalB := 0, 0
	htA, htB := map[int]bool{}, map[int]bool{}
	for k := range A {
		htA[A[k]] = true
		totalA += A[k]
	}
	for k := range B {
		htB[B[k]] = true
		totalB += B[k]
	}
	dist := (totalA - totalB) / 2
	if dist < 0 {
		dist = 0 - dist
	}
	if totalA < totalB {
		for k := range A {
			if htB[A[k]+dist] {
				return []int{A[k], A[k] + dist}
			}
		}
	} else {
		for k := range B {
			if htA[B[k]+dist] {
				return []int{B[k] + dist, B[k]}
			}
		}
	}
	return nil
}

// @lc code=end
