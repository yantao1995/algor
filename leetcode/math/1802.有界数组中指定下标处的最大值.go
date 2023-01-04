package leetcode

/*
 * @lc app=leetcode.cn id=1802 lang=golang
 *
 * [1802] 有界数组中指定下标处的最大值
 */

// @lc code=start
func maxValue(n int, index int, maxSum int) int {
	if n == 1 {
		return maxSum
	}
	level := 1
	maxSum -= n //第一层铺上
	for width := 0; maxSum > 0; width++ {
		level++
		maxSum--
		if index > width {
			maxSum -= width
		} else {
			maxSum -= index
		}
		if n-1-index > width {
			maxSum -= width
		} else {
			maxSum -= n - 1 - index
		}
		if width >= n && maxSum > 0 {
			level += maxSum / n
			maxSum = 0
		}
	}
	if maxSum < 0 {
		level--
	}
	return level
}

// @lc code=end

/*
在超时的基础上，优化左右宽度的减去运算
如果左右宽度均超过了,那么直接用除法计算需要铺多少层即可。
此处代码简化为 width >= n && maxSum > 0 ，也可以左右宽度
分别判断是否超过，可以减少运算次数


超时
Time Limit Exceeded
65/370 cases passed (N/A)
2
1
865959216
就像叠罗汉一样
level记录当前层数
width 记录当前宽度，然后依次减去左右宽度
左右宽度的计算，当前索引与左右边界的距离，如果当前要减去的
宽度比当前距离小，就直接减去，如果超过了，那么只减去最大宽度
因为只要maxSum>0,循环内level就多加了一次，所以最外面需要判断
最后的一次累加是否有效，如果无效，那么应该回退一次
func maxValue(n int, index int, maxSum int) int {
	if n == 1 {
		return maxSum
	}
	level := 1
	maxSum -= n //第一层铺上
	for width := 0; maxSum > 0; width++ {
		level++
		maxSum--
		if index > width {
			maxSum -= width
		} else {
			maxSum -= index
		}
		if n-1-index > width {
			maxSum -= width
		} else {
			maxSum -= n - 1 - index
		}
	}
	if maxSum < 0 {
		level--
	}
	return level
}


*/
