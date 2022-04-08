package leetcode

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	max := 0
	stack := make([]int, len(heights)) //存索引 底到顶增
	stack = stack[:0]
	tail := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] { //维护栈的单调性
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if i == len(heights)-1 || heights[i] > heights[i+1] { //只在下轮需要出栈时计算
			tail = len(stack) - 1
			for j := 0; j <= tail; j++ {
				if j == 0 {
					if max < heights[stack[j]]*(stack[tail]+1) {
						max = heights[stack[j]] * (stack[tail] + 1)
					}
				} else {
					if max < heights[stack[j]]*(stack[tail]-stack[j-1]) {
						max = heights[stack[j]] * (stack[tail] - stack[j-1])
					}
				}
			}
		}
	}
	return max
}

// @lc code=end

/*
	stack  单调递减栈：单调递减栈就是从栈底到栈顶数据是从小到大
	stack 存的是每个元素的索引，可以用来拿到矩形的长和高
	例如这样一组数： 2, 1, 5, 6, 2, 3  当前i走到了 3，即当前值为6
	当前stack从栈底到栈顶的索引值分别为: 1, 2, 3
	因为下轮i走到4时，值为2，比6小，不满足栈的单调性，需要依次出栈，在出栈前夕进行当前最大值计算
	遍历当前栈，当前栈索引指向的值就是矩形的宽
	栈底：第0个元素[1],是当前最小的值，也就是说当前走到i的所有值，均比它大，那么0-i 都可以是矩形的长
	栈中：第1个元素[2],是当前中间的值，(1-1,栈顶] 元素索引的范围就是 矩形的长
	栈顶：第2个元素[3],是当前最大的值，(2-1,栈顶] 元素索引的范围就是 矩形的长

	综上，除了第0个元素[栈底]需要特殊处理，其他的均是固定的 左开右闭 区间。
*/
