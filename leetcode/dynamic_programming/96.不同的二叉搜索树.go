package leetcode

/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	nums := make([]int, n+1)
	nums[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}
	return nums[n]
}

// @lc code=end

/*
	思路：将n个节点拆分，比如：
	1个节点的时候有几种，2个节点的时候,3个......
	可以由此得到 2个节点 = 1个节点 * 1个节点  ， 3个节点 = 2 * 1 ...

	dp 数组含义: dp[i]是当节点有i个时，有多少种
	j的作用，使每一个节点，都轮流当root节点。
	这样 dp[i] 就应该加上 每轮的种树，每轮的数量应该拆为左右子树的种数乘积
	子节点同理，轮流当子root结点，
	依次向下递归
*/
