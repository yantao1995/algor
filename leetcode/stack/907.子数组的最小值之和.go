package leetcode

/*
 * @lc app=leetcode.cn id=907 lang=golang
 *
 * [907] 子数组的最小值之和
 */

// @lc code=start
func sumSubarrayMins(arr []int) int {
	total := 0
	stack := []int{}
	mod := int(1e9 + 7)
	for i := 0; i <= len(arr); i++ {
		num := 0
		if i != len(arr) {
			num = arr[i]
		}
		for len(stack) > 0 && arr[stack[len(stack)-1]] > num {
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			l := -1
			if len(stack) > 0 {
				l = stack[len(stack)-1]
			}
			total += (r - l) * (i - r) * arr[r]
			total %= mod
		}
		stack = append(stack, i)
	}
	return total
}

// @lc code=end

/*
影响范围：在区间 [l,r] 内，i是其中最小的，那么在该区间内的所有包含i的子数组，最小值都是它
使用单调栈,栈内存储arr的索引, arr[i] 的影响范围索引，即 i 的左区间l，右区间r
nums[r]：在栈空时入栈 或者 栈不为空 栈顶元素>当前 :
	弹出栈顶元素 i ，对于 i 影响范围
	l为当前新的栈顶元素，r为当前 nums[r] 元素
如果栈顶元素 <= arr[i] : 则直接入栈

暴力模拟 超时
func sumSubarrayMins(arr []int) int {
	getMin := func(i, j int) int {
		result := arr[i]
		for i++; i <= j; i++ {
			if arr[i] < result {
				result = arr[i]
			}
		}
		return result
	}
	mod := int(1e9 + 7)
	total := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			total += getMin(i, j)
			total %= mod
		}
	}
	return total
}
*/
