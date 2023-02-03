package leetcode

/*
 * @lc app=leetcode.cn id=1145 lang=golang
 *
 * [1145] 二叉树着色游戏
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	ll, rr := 0, 0
	maxs := func(a int, b ...int) int {
		for k := range b {
			if b[k] > a {
				a = b[k]
			}
		}
		return a
	}
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root != nil {
			l := dfs(root.Left)
			r := dfs(root.Right)
			if root.Val == x {
				ll, rr = l, r
			}
			return l + r + 1
		}
		return 0
	}
	dfs(root)
	return maxs(ll, rr, n-1-ll-rr)*2 > n
}

// @lc code=end

/*
	贪心，划分为3棵树 ll 为x 的左子树  rr为x的右子树  n-l-ll-rr 为除x子树外的剩余树
	只需要贪最大的一个子树即可被二号玩家全部染色
	1号玩家为n1,2号玩家为n2  ， n1 = n-n2
	那么只需要 n2 > n1 即  n2 > n-n2  , 2*n2 > n
*/
