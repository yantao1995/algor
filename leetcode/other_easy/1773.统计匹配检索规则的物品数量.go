package leetcode

/*
 * @lc app=leetcode.cn id=1773 lang=golang
 *
 * [1773] 统计匹配检索规则的物品数量
 */

// @lc code=start
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	result := 0
	idx := 0
	if ruleKey == "color" {
		idx = 1
	} else if ruleKey == "name" {
		idx = 2
	}
	for _, item := range items {
		if item[idx] == ruleValue {
			result++
		}
	}
	return result
}

// @lc code=end

/*
	根据ruleKey定位索引位置，然后依次比较即可
*/
