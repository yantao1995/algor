package leetcode

/*
 * @lc app=leetcode.cn id=926 lang=golang
 *
 * [926] 将字符串翻转到单调递增
 */

// @lc code=start
func minFlipsMonoIncr(s string) int {
	count0, count1 := 0, 0
	for k := range s {
		if s[k] == '0' {
			count0++
		} else {
			count1++
		}
	}
	minf := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	min := minf(count0, count1)
	result := 0
	for k := range s {
		if s[k] == '0' {
			count0--
			min = minf(min, count0+result)
		} else {
			result++
			count1--
			min = minf(min, count1+result)
		}
	}
	return min
}

// @lc code=end

/*
	情况1 ： 全部转换成 0
	情况2 ： 全部转换成 1
	情况3 ： 在某个位置，将之前的全部转换成0，后面的全部转换成1

	count0记录当前索引到末尾的0的个数，count1记录当前索引到末尾的1的个数
	result 记录当前值如果为1转换成0的累计个数，因为如果当前索引的值1转换成0，那么前面所有的1都需要转换成0

	所以 情况1和2 可以合并成  min = minf(count0, count1)
	情况3 ，当前索引的值转化成1还是转化成0，取决于剩下的0和1的个数多少
	如果后面的1比0多，那么就将后面所有的0转化成1，但是有极端情况，
	比如这种  10001011111111 ，在当前首位索引的计算下，后面的1虽然比0多，但是当前位置仍然不是最优解的位置，
	所以需要每次比较后都计算 min 值
*/
