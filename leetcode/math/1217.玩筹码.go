package leetcode

/*
 * @lc app=leetcode.cn id=1217 lang=golang
 *
 * [1217] 玩筹码
 */

// @lc code=start
func minCostToMoveChips(position []int) int {
	single, double := 0, 0
	for k := range position {
		if position[k]&1 == 0 {
			double++
		} else {
			single++
		}
	}
	if single < double {
		return single
	}
	return double
}

// @lc code=end

/*
	题目比较绕，第i个筹码的位置在position[i]
	暴力求解可以用map存储每个位置筹码的个数，然后依次枚举每次筹码移动的位置，但这道题
	不需要暴力
	由数据  [6, 4, 7, 8, 2, 10, 2, 7, 9, 7]

	可得到 关系
	个数: 	2		1		1	3	1	1	1
	-----------------------------------------------------
	索引: 1	2	3	4	5	6	7	8	9	10

	根据移动两格，cost为0，而移动一格，cost为1,可以推断出：
		奇数索引上的筹码可以移动到任意奇数位置上而不消耗cost，而偶数同理

	那么可以得到，所有奇数位置的筹码，和所有偶数位置的筹码都移动到某两个相邻的位置

	那么将最少个数的筹码移动到另一堆的次数，就是cost的花费
*/
