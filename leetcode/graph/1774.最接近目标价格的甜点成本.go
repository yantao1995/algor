package leetcode

/*
 * @lc app=leetcode.cn id=1774 lang=golang
 *
 * [1774] 最接近目标价格的甜点成本
 */

// @lc code=start
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	absf := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	result := int(1e4)
	minAbs := int(1e4)
	var dfs func(index, total int)
	dfs = func(index, total int) {
		if val := absf(total, target); val < minAbs ||
			(val == minAbs && total < target) {
			minAbs = val
			result = total
		}
		if total >= target {
			return
		}
		if index < len(toppingCosts) {
			dfs(index+1, total)
			dfs(index+1, total+toppingCosts[index])
			dfs(index+1, total+toppingCosts[index]*2)
		}
	}

	for i := 0; i < len(baseCosts); i++ {
		dfs(0, baseCosts[i])
	}

	return result
}

// @lc code=end

/*
	根据题目数据规模，1e4的数据，那可以直接dfs爆搜
	在baseCost[i] 必选的情况下
		toppingCosts[j] 分3种情况:
			不选：dfs(index+1, total)
			选1：dfs(index+1, total+toppingCosts[index])
			选2：dfs(index+1, total+toppingCosts[index]*2)
*/
