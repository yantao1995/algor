package leetcode

/*
 * @lc app=leetcode.cn id=1687 lang=golang
 *
 * [1687] 从仓库到码头运输箱子
 */

// @lc code=start
func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	dp := make([]int, len(boxes))
	dp[0] = 2
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	times := 0
	for i := 1; i < len(boxes); i++ {
		dp[i] = dp[i-1] + 2
		for j, tempWeight := i, 0; j >= 0 && i-j < maxBoxes && tempWeight+boxes[j][1] <= maxWeight; j-- {
			tempWeight += boxes[j][1]
			if j == i {
				times = 2
			} else if boxes[j][0] != boxes[j+1][0] {
				times++
			}
			if j == 0 {
				dp[i] = min(dp[i], times)
			} else {
				dp[i] = min(dp[i], dp[j-1]+times)
			}
		}
	}
	return dp[len(dp)-1]
}

// @lc code=end

/*


动态规划 超时 (35/39)： dp[i] 到第i个箱子的最小行程，第i个依次与前面的箱子进行组合打包
func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	dp := make([]int, len(boxes))
	dp[0] = 2
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	times := 0
	for i := 1; i < len(boxes); i++ {
		dp[i] = dp[i-1] + 2
		for j, tempWeight := i, 0; j >= 0 && i-j < maxBoxes && tempWeight+boxes[j][1] <= maxWeight; j-- {
			tempWeight += boxes[j][1]
			if j == i {
				times = 2
			} else if boxes[j][0] != boxes[j+1][0] {
				times++
			}
			if j == 0 {
				dp[i] = min(dp[i], times)
			} else {
				dp[i] = min(dp[i], dp[j-1]+times)
			}
		}
	}
	return dp[len(dp)-1]
}



dfs 超时 (5/39) : 枚举当前可以分成一组的和不能为一组的所有组合情况
func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	result := len(boxes) * 2
	var dfs func(index, times int)
	dfs = func(index, times int) {
		if index == len(boxes) {
			if times < result {
				result = times
			}
			return
		}
		times += 2
		for i, currentBox, currentWeight := index, 0, 0; i < len(boxes) &&
			currentBox < maxBoxes && currentWeight+boxes[i][1] <= maxWeight; i++ {
			currentBox++
			currentWeight += boxes[i][1]
			if currentBox > 1 && boxes[i][0] != boxes[i-1][0] {
				times++
			}
			dfs(i+1, times)
		}
	}
	dfs(0, 0)
	return result
}
*/
