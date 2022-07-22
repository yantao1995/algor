package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=757 lang=golang
 *
 * [757] 设置交集大小至少为2
 */

// @lc code=start
func intersectionSizeTwo(intervals [][]int) int {
	m := map[int]bool{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] ||
			(intervals[i][0] == intervals[j][0] &&
				intervals[i][1] < intervals[j][1])
	})
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//找交集
	getSet := func(i1, j1, i2, j2 int) (bool, int, int) {
		if i1 > i2 {
			i1, j1, i2, j2 = i2, j2, i1, j1
		}
		if i2 <= j1 {
			return true, max(i1, i2), min(j1, j2)
		}
		return false, 0, 0
	}

	flag, start, end, lastEnd := true, 0, 0, 0
	count := 2
nextI:
	for i := 0; i < len(intervals); i++ {
		count = 2
		//先看是否存在
		for k := range m {
			if k >= intervals[i][0] && k <= intervals[i][1] {
				count--
			}
			if count == 0 {
				continue nextI
			}
		}
		//向后找交集
		start, end, lastEnd = intervals[i][0], intervals[i][1], intervals[i][1]
		for j := i + 1; j < len(intervals); j++ {
			flag, start, end = getSet(start, end, intervals[j][0], intervals[j][1])
			if flag {
				lastEnd = end
			}
			if !flag || start == end {
				if !m[lastEnd] {
					m[lastEnd] = true
					count--
				}
				if count == 0 {
					continue nextI
				}
				break
			}
		}
		//补齐count
		start, end = intervals[i][0], lastEnd
		for ; start <= end && count > 0; end-- {
			if !m[end] {
				m[end] = true
				count--
			}
		}
	}
	return len(m)
}

// @lc code=end

/*
	将所有区间按升序排列，优先区间左值升序，左值相同就按右值升序
	（因为是升序排列，所以选择时优先选择最大值）
	m 用来存结果集
	每个区间需要找count=2的数量，先找m中是否存在，如果m中不存在，再将当前区间向后求交集
	一直到当前区间相交的重复区间内的整数，数量为1或者不存在为止(flag==false 或者 start==end)
	【为1那是因为前面的count过来的时候可能为1或者2，至少需要找1个数。如果不存在，那么上一次的交集肯定大于等于2】
	所以需要用lastEnd 记录上一次交集的最大值

	最后从最后一次交集的end开始向前补不满足count个数的数字加入到m中
*/
