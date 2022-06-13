package leetcode

/*
 * @lc app=leetcode.cn id=830 lang=golang
 *
 * [830] 较大分组的位置
 */

// @lc code=start
func largeGroupPositions(s string) [][]int {
	bt := []byte{}
	index := [][]int{}
	lastIndex := 0
	for i := 0; i < len(s); i++ {
		bt = append(bt, s[i])
		if i+1 < len(s) && s[i] != s[i+1] {
			bt = append(bt, ',')
			index = append(index, []int{lastIndex, i})
			lastIndex = i + 1
		}
	}
	// fmt.Println(string(bt))
	// fmt.Println(index)
	index = append(index, []int{lastIndex, len(s) - 1})
	result := [][]int{}
	for k := range index {
		if index[k][1]-index[k][0] >= 2 {
			result = append(result, []int{index[k][0], index[k][1]})
		}
	}
	return result
}

// @lc code=end
