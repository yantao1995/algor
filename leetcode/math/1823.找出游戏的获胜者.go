package leetcode

/*
 * @lc app=leetcode.cn id=1823 lang=golang
 *
 * [1823] 找出游戏的获胜者
 */

// @lc code=start
func findTheWinner(n int, k int) int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i
	}

	index := 0
	for len(list) > 1 {
		index = k % len(list)
		if index == 0 {
			list = list[:len(list)-1]
		} else {
			list = append(list[index:], list[:index-1]...)
		}
	}

	return list[0] + 1
}

// @lc code=end

/*
	直接模拟整个游戏过程
	轮到第k个，就直接删除第k个数据，
	然后把k+1到len(list)的数据移到 0到k-1 的数据的前面
	然后再次计算要剔除的数据
*/
