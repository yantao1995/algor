package leetcode

/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	ht := map[string]int{}
	for k := range list1 {
		ht[list1[k]] = k
	}
	minIndex := 1<<63 - 1
	find := false
	for k := range list2 {
		if v, ok := ht[list2[k]]; ok {
			if k+v < minIndex {
				minIndex = k + v
				find = true
			}
		}
	}
	if find {
		strs := []string{}
		for k := range list2 {
			if v, ok := ht[list2[k]]; ok {
				if k+v == minIndex {
					strs = append(strs, list2[k])
				}
			}
		}
		return strs
	}
	return nil
}

// @lc code=end
