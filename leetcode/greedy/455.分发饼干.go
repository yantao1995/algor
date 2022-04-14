package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(s); j++ {
			if g[i] <= s[j] {
				count++
				s[j] = 0
				break
			}
		}
	}
	return count
}

// @lc code=end

/*
	贪心 ，先将g和s升序
	然后优先满足胃口最小的，吃刚好够的饼干，然后把饼干置0
	然后依次找满足胃口的。
*/
