package leetcode

/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		if cost[i-1] < cost[i-2] {
			cost[i] += cost[i-1]
		} else {
			cost[i] += cost[i-2]
		}
	}
	if cost[len(cost)-2] < cost[len(cost)-1] {
		return cost[len(cost)-2]
	}
	return cost[len(cost)-1]
}

// @lc code=end

/*
	因为可以从 0 和 1 开始， 那么 dp初始化 0和1 就是  cost[0] 和cost[1]

	又因为 cost 数组只遍历一次就没用了，所以直接可以复用 成dp数组

	dp[i] 每次在 dp[i-1] 和 dp[i-2] 中取最小值，然后加上当前的cost[i] 就是当前 dp[i] 的最小花费

	因为一次可以跳1层或者2层，所以结果是 len(cost)-1  和 len(cost)-2 中任一一个，取最小值

*/
