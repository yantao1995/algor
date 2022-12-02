package leetcode

/*
 * @lc app=leetcode.cn id=1769 lang=golang
 *
 * [1769] 移动所有球到每个盒子所需的最小操作数
 */

// @lc code=start
/* func minOperations(boxes string) []int {
	resultFront := make([]int, len(boxes))
	resultTail := make([]int, len(boxes))
	front := make([]int, len(boxes))
	tail := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		if i > 0 {
			front[i] = front[i-1]
		}
		if boxes[i] == '1' {
			front[i]++
		}
	}
	for i := len(boxes) - 1; i >= 0; i-- {
		if i < len(boxes)-1 {
			tail[i] = tail[i+1]
		}
		if boxes[i] == '1' {
			tail[i]++
		}
	}
	for i := 1; i < len(boxes); i++ {
		resultFront[i] = front[i-1] + resultFront[i-1]
	}
	for i := len(boxes) - 2; i >= 0; i-- {
		resultTail[i] = tail[i+1] + resultTail[i+1]
	}
	for i := 0; i < len(resultFront); i++ {
		resultFront[i] += resultTail[i]
	}
	return resultFront
} */

// @lc code=end

/*
	front 和 tail 分别记录从前往后 和 从后往前 的累积 1 的个数
	resultFront 和 resultTail  分别记录移动到当前位置的操作数
*/
