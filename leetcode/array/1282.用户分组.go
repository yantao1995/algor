package leetcode

/*
 * @lc app=leetcode.cn id=1282 lang=golang
 *
 * [1282] 用户分组
 */

// @lc code=start
func groupThePeople(groupSizes []int) [][]int {
	m := map[int]int{}
	result := [][]int{}
lab:
	for k := range groupSizes {
		for kk := range result {
			if len(result[kk]) < m[kk] && m[kk] == groupSizes[k] {
				result[kk] = append(result[kk], k)
				continue lab
			}
		}
		result = append(result, []int{k})
		m[len(result)-1] = groupSizes[k]
	}
	return result
}

// @lc code=end

/*
	result 记录每一个组
	m记录result结果集索引位置处的组所需要的人数
	遍历groupSizes,找到每个人所需要的组人数
	然后遍历结果集，查看当前结果集中的组是否人数和组内人员要求的大小是否满足
	满足则入组
	如果在result中未找到满足的组，那么就直接新建一个组
*/
