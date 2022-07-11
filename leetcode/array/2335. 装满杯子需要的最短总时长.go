package leetcode

import "sort"

/*
现有一台饮水机，可以制备冷水、温水和热水。每秒钟，可以装满 2 杯 不同 类型的水或者 1 杯任意类型的水。
给你一个下标从 0 开始、长度为 3 的整数数组 amount ，其中 amount[0]、amount[1] 和 amount[2] 分别表示需要装满冷水、温水和热水的杯子数量。返回装满所有杯子所需的 最少 秒数。

输入：amount = [1,4,2]
输出：4
解释：下面给出一种方案：
第 1 秒：装满一杯冷水和一杯温水。
第 2 秒：装满一杯温水和一杯热水。
第 3 秒：装满一杯温水和一杯热水。
第 4 秒：装满一杯温水。
可以证明最少需要 4 秒才能装满所有杯子。

*/
func fillCups(amount []int) int {
	if len(amount) < 2 {
		if len(amount) == 0 {
			return 0
		}
		return amount[0]
	}
	result := 0
	sort.Slice(amount, func(i, j int) bool {
		return amount[j] < amount[i]
	})
	for amount[0] > 0 {
		if amount[1] > 0 {
			amount[0]--
			amount[1]--
			result++
		} else {
			result += amount[0]
			amount[0] = 0
			break
		}
		sort.Slice(amount, func(i, j int) bool {
			return amount[j] < amount[i]
		})
	}
	return result
}

/*
	倒序排列，然后用第一个杯子和第二个杯子同时一起接，如果第二杯为0了，那么只需要全接第一杯了。
*/
