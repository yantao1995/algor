package leetcode

/*
 * @lc app=leetcode.cn id=779 lang=golang
 *
 * [779] 第K个语法符号
 */

// @lc code=start
func kthGrammar(n int, k int) int {
	log := make([]int, n+1)
	for tn, tk := n, k; tn > 0; tn, tk = tn-1, (tk+1)/2 {
		log[tn] = tk
	}
	log[1] = 0
	temp := [2]int{}
	for i := 2; i < len(log); i++ {
		if log[i-1] == 0 {
			temp[0], temp[1] = 0, 1
		} else {
			temp[0], temp[1] = 1, 0
		}
		if log[i]%2 == 1 {
			log[i] = temp[0]
		} else {
			log[i] = temp[1]
		}
	}
	return log[n]
}

// @lc code=end
